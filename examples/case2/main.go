//支持域敏感但不支持域敏感嵌套

package main

type No struct {
	Value interface{}
}

type Node struct {
	Value interface{}
	Noo   No
}

func main() {
	a := Node{Value: 1, Noo: No{}}
	_ = a.Value.(int)
	b := Node{Value: true, Noo: No{}}
	_ = b.Value.(bool)
	c := Node{Value: "hello", Noo: No{}}
	_ = c.Value.(string)
}
