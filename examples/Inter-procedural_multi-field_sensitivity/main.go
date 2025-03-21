//支持域敏感但不支持域敏感嵌套

package main

type C struct {
	Data interface{}
}

type B struct {
	c    C
	Data interface{}
}

type A struct {
	b    B
	Data interface{}
}

func main() {
	var a A
	a.b.c.Data = 1
	a.b.Data = 2
	a.Data = 3
	_ = a.b.c.Data.(int)
	_ = a.b.Data.(int)
	_ = a.Data.(int)
}
