package main

import "fmt"

func main()  {
	var intValueBit uint8
	intValueBit = 255
	intValueBit = ^intValueBit // 按位取反
	fmt.Println(intValueBit)   // 0
	intValueBit = 1
	intValueBit = intValueBit << 3 // 左移 3 位，相当于乘以 2^3 = 8
	fmt.Println(intValueBit)       // 8
}
