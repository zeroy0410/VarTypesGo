package typeAnalyzer

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/callgraph/vta"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"strings"
)

func getTypes(t types.Type) (types.Type, string) {
	switch tt := t.(type) {
	case *types.Pointer:
		_, kind := getTypes(tt.Elem())
		return tt, "pointer to " + kind
	case *types.Named:
		return tt.Underlying(), tt.Underlying().String()
	case *types.Interface:
		return tt, "interface"
	case *types.Struct:
		return tt, "struct"
	default:
		return tt, fmt.Sprintf("%T", tt)
	}
}

func extractInterfaceFromPointer(t types.Type) (*types.Interface, bool) {
	ptrType, ok := t.(*types.Pointer)
	if !ok {
		return nil, false
	}

	elemType := ptrType.Elem()
	for {
		switch tt := elemType.(type) {
		case *types.Named:
			elemType = tt.Underlying()
		case *types.Interface:
			return tt, true
		default:
			return nil, false
		}
	}
}

func PrintAssertionsInfo(resultTypes map[*ssa.TypeAssert][]types.Type) {
	for node, possibleTypes := range resultTypes {
		parentFunc := node.X.Parent()
		if pkg := parentFunc.Package(); pkg != nil {
			fmt.Println("Package: ", pkg)
		}
		fmt.Println("Function: ", parentFunc.Name())
		fmt.Println("Node: ", node.X)
		fmt.Print("Assertion: ", node.AssertedType)

		assertedType, assertedTypeStr := getTypes(node.AssertedType)
		fmt.Print("    ", assertedTypeStr)
		fmt.Println()
		fmt.Println("Possible Types: ")

		for _, typ := range possibleTypes {
			fmt.Print("    ", typ)
			actualType, typeStr := getTypes(typ)
			fmt.Print("    ", typeStr)

			if strings.Contains(assertedTypeStr, "pointer to struct") && strings.Contains(typeStr, "pointer to interface") {
				interfaceType, ok := extractInterfaceFromPointer(actualType)
				if !ok {
					fmt.Print("    Unable to extract interface type")
					continue
				}

				interfaceType.Complete()
				implements := types.Implements(assertedType, interfaceType)
				fmt.Print("    Implements: ", implements)
			}
			fmt.Println()
		}
		fmt.Println("-----------------------")
	}
}

// Runner represents a analysis runner
type Runner struct {
	ModuleName string
	PkgPath    []string
	Debug      bool
	Dir        string
}

func NewRunner(PkgPath ...string) *Runner {
	return &Runner{PkgPath: PkgPath, ModuleName: "",
		Debug: false}
}

func (r *Runner) Run() error {
	mode := packages.NeedName |
		packages.NeedFiles |
		packages.NeedCompiledGoFiles |
		packages.NeedSyntax |
		packages.NeedTypesInfo |
		packages.NeedImports |
		packages.NeedTypesSizes |
		packages.NeedTypes |
		packages.NeedDeps
	cfg := &packages.Config{Mode: mode, Dir: r.Dir}
	initial, err := packages.Load(cfg, r.PkgPath...)

	if err != nil {
		return err
	}

	prog, _ := ssautil.AllPackages(initial, 0)

	prog.Build()

	mainFuncs := make([]*ssa.Function, 0)
	for _, pkg := range initial {
		mainPkg := prog.Package(pkg.Types)
		if mainPkg != nil && mainPkg.Pkg.Name() == "main" && mainPkg.Func("main") != nil {
			mainFuncs = append(mainFuncs, mainPkg.Func("main"))
		}
	}
	if len(mainFuncs) == 0 {
		return new(NoMainPkgError)
	}

	vta.CallGraph(ssautil.AllFunctions(prog), nil)
	resultTypes := vta.GetTypeAsserts(ssautil.AllFunctions(prog), nil)
	PrintAssertionsInfo(resultTypes)

	return nil
}
