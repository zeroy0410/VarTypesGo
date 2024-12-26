// The example context insensitive analysis can not be solved.
// 上下文敏感

package main

import "fmt"

func processBool(i interface{}) bool {
	return i.(bool) // type assertion
}
func processint(i interface{}) int {
	return i.(int) // type assertion
}
func processString(i interface{}) string {
	return i.(string) // type assertion
}
func process(i interface{}) interface{} {
	return i
}

func main() {
	fmt.Println("Hello, world!")
	var boolInterface interface{} = true
	var intInterface interface{} = 1
	var stringInterface interface{} = "hello"
	boolInterface = process(boolInterface)
	intInterface = process(intInterface)
	stringInterface = process(stringInterface)
	processBool(boolInterface)
	processint(intInterface)
	processString(stringInterface)

	type No struct {
		Value interface{}
	}

	type Node struct {
		Value interface{}
		Noo   No
	}

	a := Node{Value: 1, Noo: No{}}
	_ = a.Value.(int)
	b := Node{Value: true, Noo: No{}}
	_ = b.Value.(bool)
	c := Node{Value: "hello", Noo: No{}}
	_ = c.Value.(string)
}
