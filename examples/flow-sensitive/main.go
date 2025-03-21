package main

type Person interface {
	Name() string
}

type Stu struct {
	name string
}

func (s Stu) Name() string {
	return s.name
}

type Tea struct {
	name string
}

func (tea Tea) Name() string {
	return tea.name
}

func main() {
	var i Person
	i = &Stu{"stu"}
	i = &Tea{"tea"}
	_ = i.(Tea)
}

// i_0 = Stu{"stu"}
// i_1 = Tea{"tea"}
// _ = i_1.(Tea)
