package main

type Node struct {
	Value interface{}
}

func main() {
	var a, b, c Node
	a.Value = 1
	b.Value = true
	c.Value = "hello"
	_ = a.Value.(int)
	_ = b.Value.(bool)
	_ = c.Value.(string)

	var d, e, f interface{}
	d = 1
	e = true
	f = "hello"
	_ = d.(int)
	_ = e.(bool)
	_ = f.(string)
}
