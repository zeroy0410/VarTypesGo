package main

func process(i interface{}) interface{} {
	return i
}

func wrap(i interface{}) interface{} {
	return process(i) // 调用点A
}

func assertInt(i interface{}) int {
	return i.(int) // 类型断言点X
}

func assertString(i interface{}) string {
	return i.(string) // 类型断言点Y
}

func main() {
	// 两个不同的数据源
	intVal := 42
	strVal := "hello"

	// 通过不同路径传递
	path1 := wrap(intVal) // call@25 -> call@8
	path2 := wrap(strVal) //  call@26 -> call@8

	// 类型断言
	assertInt(path1)    // 应接收int
	assertString(path2) // 应接收string
}
