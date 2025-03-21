//完整支持域敏感

package main

type E struct {
	Data interface{}
}

type D struct {
	e    E
	Data interface{}
}

type C struct {
	d    D
	Data interface{}
}

type B struct {
	c    C
	Data interface{}
}

type A struct {
	b    *B
	Data interface{}
}

func main() {
	a := A{}
	pb := B{}
	pb.c.d.e.Data = 114514
	pb.c.Data = true
	a.b = &pb
	pb.c.d.e.Data = false
	_ = a.b.c.d.e.Data.(int)
	_ = a.b.c.Data.(bool)
}
