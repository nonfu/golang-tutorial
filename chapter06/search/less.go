package main

import "fmt"

// 二分查找变形版本4：查找最后一个小于等于给定值的元素
func binarySearchV4(nums []int, num, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if num >= nums[mid] {
		// 判断条件变更，找到最后一个小于等于给定值的元素
		// 最右侧元素值，或者当前 mid 位置后一个元素值大于给定值
		if mid == len(nums) - 1 || nums[mid + 1] > num {
			return mid
		} else {
			return binarySearchV4(nums, num, mid + 1, high)
		}
	} else {
		return binarySearchV4(nums, num, low, mid - 1)
	}
}

func main() {
	nums:= []int{1, 2, 3, 3, 4, 5, 6}
	num := 3
	index := binarySearchV4(nums, num, 0, len(nums) - 1)
	if index != -1 {
		fmt.Printf("Find last num less or equal than %d at index %d\n", num, index)
	} else {
		fmt.Printf("Num %d not exists in nums\n", num)
	}
}