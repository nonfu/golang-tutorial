package main

import "fmt"

// 二分查找变形版本3：查找第一个大于等于给定值的元素
func binarySearchV3(nums []int, num, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if num <= nums[mid] {
		// 判断条件变更，找到第一个大于等于给定值的元素
		// 最左侧元素值，或者当前 mid 位置前一个元素值小于给定值
		if mid == 0 || nums[mid-1] < num {
			return mid
		} else {
			return binarySearchV3(nums, num, low, mid-1)
		}
	} else {
		return binarySearchV3(nums, num, mid+1, high)
	}
}

func main()  {
	nums := []int{1, 2, 3, 3, 4, 5, 6}
	num := 3
	index := binarySearchV3(nums, num, 0, len(nums) - 1)
	if index != -1 {
		fmt.Printf("Find first num greater or equal than %d at index %d\n", num, index)
	} else {
		fmt.Printf("Num %d not exists in nums\n", num)
	}
}
