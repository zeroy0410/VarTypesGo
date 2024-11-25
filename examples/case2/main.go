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

	aa := a.Noo
	aa.Value = "haha"
	_ = aa.Value.(string)
	bb := b.Noo
	bb.Value = 2
	_ = bb.Value.(int)
	cc := c.Noo
	cc.Value = false
	_ = cc.Value.(bool)
}
