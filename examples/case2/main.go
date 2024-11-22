package main

type No struct {
	Value interface{}
}

type Node struct {
	Value interface{}
	Noo   No
}

func main() {
	a := Node{Value: 1, Noo: No{"hello"}}
	b := Node{Value: true, Noo: No{1}}
	c := Node{Value: "hello", Noo: No{false}}
	aa := a.Noo
	bb := b.Noo
	cc := c.Noo
	_ = aa.Value.(string)
	_ = bb.Value.(int)
	_ = cc.Value.(bool)
	_ = a.Value.(int)
	_ = b.Value.(bool)
	_ = c.Value.(string)
}
