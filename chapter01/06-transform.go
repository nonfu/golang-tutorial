package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整型 <=> 整型
	//v1 := -1
	//v2 := uint8(v1)
	//fmt.Println(v2)  // 255

	// 整型 <=> 浮点型
	//v3 := 99
	//v4 := float32(v3)
	//fmt.Println(v4)  // 99

	// 字符串
	v1 := "100"
	v2, _ := strconv.Atoi(v1) // 将字符串转化为整型，v2 = 100
	fmt.Println(v2)

	v3 := 100
	v4 := strconv.Itoa(v3)   // 将整型转化为字符串, v4 = "100"
	fmt.Println(v4)

	v5 := "true"
	v6, _ := strconv.ParseBool(v5) // 将字符串转化为布尔型
	v5 = strconv.FormatBool(v6)    // 将布尔值转化为字符串
	fmt.Println(v5)
	fmt.Println(v6)

	v7 := "100"
	v8, _ := strconv.ParseInt(v7, 10, 64) // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
	v7 = strconv.FormatInt(v8, 10)        // 将整型转化为字符串，第二个参数表示进制
	fmt.Println(v7)
	fmt.Println(v8)

	v9, _ := strconv.ParseUint(v7, 10, 64) // 将字符串转化为无符号整型，参数含义同 ParseInt
	v7 = strconv.FormatUint(v9, 10)        // 将字符串转化为无符号整型，参数含义同 FormatInt
	fmt.Println(v7)
	fmt.Println(v9)

	v10 := "99.99"
	v11, _ := strconv.ParseFloat(v10, 64) // 将字符串转化为浮点型，第二个参数表示精度
	v10 = strconv.FormatFloat(v11, 'E', -1, 64)
	fmt.Println(v10)
	fmt.Println(v11)

	q := strconv.Quote("Hello, 世界")    // 为字符串加引号
	q = strconv.QuoteToASCII("Hello, 世界")  // 将字符串转化为 ASCII 编码
	fmt.Println(q)
}
