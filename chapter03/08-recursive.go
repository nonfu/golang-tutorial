package main

import (
	"fmt"
	"time"
)

type FibonacciFunc func(int) int

// 通过递归函数实现斐波那契数列
func fibonacci(n int) int {
	// 终止条件
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	// 递归公式
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciExecTime(f FibonacciFunc) FibonacciFunc {
	return func(n int) int {
		start := time.Now()      // 起始时间
		num := f(n)              // 执行斐波那契函数
		end := time.Since(start) // 函数执行完毕耗时
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return num // 返回计算结果
	}
}

const MAX = 50

var fibs [MAX]int

// 缓存中间结果的递归函数优化版
func fibonacci2(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	index := n - 1
	if fibs[index] != 0 {
		return fibs[index]
	}

	num := fibonacci2(n-1) + fibonacci2(n-2)
	fibs[index] = num
	return num
}

// 尾递归版本的斐波那契函数
func fibonacci3(n int) int {
	return fibonacciTail(n, 0, 1)
}

func fibonacciTail(n, first, second int) int {
	if n < 2 {
		return first
	}
	return fibonacciTail(n-1, second, first+second)
}

func main() {
	n1 := 5
	f1 := fibonacciExecTime(fibonacci)
	r1 := f1(n1)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, r1)

	n2 := 50
	r2 := f1(n2)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r2)

	f2 := fibonacciExecTime(fibonacci2)
	r3 := f2(n2)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r3)

	f3 := fibonacciExecTime(fibonacci3)
	r4 := f3(n2)
	fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r4)
}
