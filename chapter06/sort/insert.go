package main

import "fmt"

func insertionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}

	return nums
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	nums = insertionSort(nums)
	fmt.Println(nums)
}