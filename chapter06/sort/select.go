package main

import "fmt"

func selectionSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 已排序区间初始化为空，未排序区间初始化待排序切片
	for i := 0; i < len(nums); i++ {
		// 未排序区间最小值初始化为第一个元素
		min := i
		// 从未排序区间第二个元素开始遍历，直到找到最小值
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// 将最小值与未排序区间第一个元素互换位置（等价于放到已排序区间最后一个位置）
		if min != i {
			nums[i],nums[min] = nums[min], nums[i]
		}
	}
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	selectionSort(nums)
	fmt.Println(nums)
}
