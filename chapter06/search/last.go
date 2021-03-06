package main

import "fmt"

// 二分查找变形版本2：查找最后一个值等于给定值的元素
func binarySearchV2(nums []int, num, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if num < nums[mid] {
		return binarySearchV2(nums, num, low, mid - 1)
	} else if num > nums[mid] {
		return binarySearchV2(nums, num, mid + 1, high)
	} else {
		// 除了需要判断最后一个与 num 相等的元素索引外，其他和普通二分查找逻辑一致
		if mid == len(nums) - 1 || nums[mid + 1] != num {
			return mid
		} else {
			return binarySearchV2(nums, num, mid + 1, high)
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 3, 4, 5, 6}
	num := 3
	index := binarySearchV2(nums, num, 0, len(nums) - 1)
	if index != -1 {
		fmt.Printf("Find num %d last at index %d\n", num, index)
	} else {
		fmt.Printf("Num %d not exists in nums\n", num)
	}
}
