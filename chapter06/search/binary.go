package main

import (
	"fmt"
	"sort"
)

// 二分查找实现代码
func binarySearch(nums []int, num int, low int, high int) int {
	// 递归终止条件
	if low > high {
		return -1
	}

	// 通过中间元素进行二分查找
	mid := (low + high) / 2
	// 递归查找
	if num > nums[mid] {
		// 如果待查找数据大于中间元素，则在右区间查找
		return binarySearch(nums, num, mid + 1, high)
	} else if num < nums[mid] {
		// 如果待查找数据小于中间元素，则在左区间查找
		return binarySearch(nums, num, low, mid - 1)
	} else {
		// 找到了，返回索引值
		return mid
	}
}

func main() {
	nums := []int{4, 6, 5, 3, 1, 8, 2, 7}
	sort.Ints(nums)  // 先对待排序数据序列进行排序
	fmt.Printf("Sorted nums: %v\n", nums)
	num := 5
	index := binarySearch(nums, num, 0, len(nums)-1)
	if index != -1 {
		fmt.Printf("Find num %d at index %d\n", num, index)
	} else {
		fmt.Printf("Num %d not exists in nums\n", num)
	}
}
