// 多态调用 dispatch 能力

package main

import (
	"fmt"
	"os/exec"
)

type Animal interface {
	Run(s string)
	Eat(s string)
}

type Cat struct {
	Name string
}

func (c Cat) Run(s string) {
	fmt.Println(c.Name, " run")
}

func (c Cat) Eat(s string) {
	fmt.Println(c.Name, " eat")
}

type Dog struct {
	Name string
}

func (d Dog) Run(s string) {
	_ = exec.Command(s, "")
	fmt.Println(d.Name, " run")
}

func (d Dog) Eat(s string) {
	fmt.Println(d.Name, " eat")
}

func Solve(source string) {
	a := Cat{Name: "Cat"}
	a.Run(source)
}

func main() {
	Solve("haha")
}
