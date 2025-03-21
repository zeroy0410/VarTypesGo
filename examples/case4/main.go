package main

// curriedAdd 函数接受一个整数 x，并返回一个新的函数。
// 该新函数接受另一个整数 y，然后返回 x 和 y 的和。
func curriedAdd(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
func main() {
	// 创建一个新的函数 addTwo，它是 curriedAdd(2) 的结果。
	addTwo := curriedAdd(2)
	// 调用 addTwo(3)，相当于 curriedAdd(2)(3)，结果为 5。
	_ = addTwo(3) // 输出: 5
	// 也可以直接用柯里化的方式调用
	_ = curriedAdd(5)(10) // 输出: 15
}
