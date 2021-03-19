package main

import "fmt"

// BF 算法实现函数
func bfSearch(s, p string) int {
	begin := 0
	i, j := 0, 0
	n, m := len(s), len(p)  // 主串、子串长度
	for i = 0; i < n; begin++ {
		// 通过 BF 算法暴力匹配子串和主串
		for j = 0;j < m; j++ {
			if i < n && s[i] == p[j] {
				// 如果子串和主串对应字符相等，逐一往后匹配
				i++
			} else {
				// 否则退出当前循环，从主串下一个字符继续开始匹配
				break
			}
		}
		if j == m {
			// 子串遍历完，表面已经找到，返回对应下标
			return i - j
		}
		// 从下一个位置继续开始匹配
		i = begin
		i++
	}
	return -1
}

// 基于 BF 算法实现字符串查找函数
func strStrV1(haystack, needle string) int {
	// 子串长度=0
	if len(needle) == 0 {
		return 0
	}
	//主串长度=0，或者主串长度小于子串长度
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}
	// 调用 BF 算法查找子串
	return bfSearch(haystack, needle)
}

func main() {
	s := "Hello, 学院君!"
	p := "学院君"
	pos := strStrV1(s, p)
	fmt.Printf("Find \"%s\" at %d in \"%s\"\n", p, pos, s)
}