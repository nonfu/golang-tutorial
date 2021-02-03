package main

import (
	"fmt"
)

// 将函数作为返回值类型
func deferAdd(a, b int) func() int {
	return func() int {
		return a + b
	}
}

func main() {
	add := func(a, b int) int {
		return a + b
	}

	// 将匿名函数作为参数
	func(call func(int, int) int) {
		fmt.Println(call(1, 2))
	}(add)

	// 此时返回的是匿名函数
	addFunc := deferAdd(1, 2)
	// 这里才会真正执行加法操作
	fmt.Println(addFunc())
}
