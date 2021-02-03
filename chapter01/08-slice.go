package main

import "fmt"

func main() {
	// 基本使用
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	q2 := months[3:6]    // 第二季度
	summer := months[5:8]  // 夏季

	fmt.Println(q2)
	fmt.Println(summer)

	// 切片复制
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{5, 4, 3}

	// 复制 slice1 到 slice 2
	//copy(slice2, slice1) // 只会复制 slice1 的前 3 个元素到 slice2 中
	//fmt.Println(slice2)
	// 复制 slice2 到 slice 1
	//copy(slice1, slice2) // 只会复制 slice2 的 3 个元素到 slice1 的前 3 个位置
	//fmt.Println(slice1)

	// 数据共享
	slice1 := make([]int, 4, 5)
	slice2 := slice1[1:3]
	slice1 = append(slice1, 0)
	slice1[1] = 2
	slice2[1] = 6

	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
}
