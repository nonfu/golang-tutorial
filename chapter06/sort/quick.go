package main

import "fmt"

// 快速排序入口函数
func quickSort(nums []int, p int, r int) {
	// 递归终止条件
	if p >= r {
		return
	}
	// 获取分区位置
	q := partition(nums, p, r)
	// 递归分区（排序是在定位 pivot 的过程中实现的）
	quickSort(nums, p, q - 1)
	quickSort(nums, q + 1, r)
}

// 定位 pivot
func partition(nums []int, p int, r int) int {
	// 以当前数据序列最后一个元素作为初始 pivot
	pivot := nums[r]
	// 初始化 i、j 下标
	i := p
	// 后移 j 下标的遍历过程
	for j := p; j < r; j++ {
		// 将比 pivot 小的数丢到 [p...i-1] 中，剩下的 [i...j] 区间都是比 pivot 大的
		if nums[j] < pivot {
			// 互换 i、j 下标对应数据
			nums[i], nums[j] = nums[j], nums[i]
			// 将 i 下标后移一位
			i++
		}
	}

	// 最后将 pivot 与 i 下标对应数据值互换
	// 这样一来，pivot 就位于当前数据序列中间，i 也就是 pivot 值对应的下标
	nums[i], nums[r] = pivot, nums[i]
	// 返回 i 作为 pivot 分区位置
	return i
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	quickSort(nums, 0, len(nums) - 1)
	fmt.Println(nums)
}
