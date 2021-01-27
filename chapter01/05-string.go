package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "hello, world"
	str1 := str[:5]  // 获取索引5（不含）之前的子串
	str2 := str[7:]  // 获取索引7（含）之后的子串
	str3 := str[0:5]  // 获取从索引0（含）到索引5（不含）之间的子串
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	fmt.Println("str3:", str3)

	str4 := str[:]
	fmt.Println("str4:", str4)

	msg := "Hello, 世界"
	n := len(msg)  //  获取到的是 msg 的字节长度
	for i := 0; i < n; i++ {
		ch := msg[i]    // 依据下标取字符串中的字符，ch 类型为 byte
		fmt.Println(i, ch, reflect.TypeOf(ch))
	}

	for i, ch := range msg {
		fmt.Println(i, ch, reflect.TypeOf(ch))  // 通过 range 遍历，ch 类型是 rune
	}
}
