package main

import "fmt"

func main() {
	// 数组的初始化
	a1 := [5]uint8{1, 2}
	fmt.Println(a1)

	a2 := [3]string{"hello", "world"}
	fmt.Println(a2)

	// 遍历数组
	arr := [5]int{1,2,3,4,5}
	/*for i := 0; i < len(arr); i++ {
		fmt.Println("Element", i, "of arr is", arr[i])
	}*/
	for i, v := range arr {
		fmt.Println("Element", i, "of arr is", v)
	}

	// 通过二维数组打印九九乘法表
	var multi [9][9]string
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 {  // 摒除重复的记录
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1 * n2)
		}
	}

	// 打印九九乘法表
	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2)  // 位宽为8，左对齐
		}
		fmt.Println()
	}
}
