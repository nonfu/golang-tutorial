package main

import (
	"fmt"
)

// 定义链表节点
type Node struct {
	Value int
	Next  *Node
}

// 初始化栈结构（空栈）
var size = 0
var stack = new(Node)

// 进栈
func Push(v int) bool {
	// 空栈的话直接将值放入头节点即可
	if stack == nil {
		stack = &Node{v, nil}
		size = 1
		return true
	}

	// 否则将插入节点作为栈的头节点
	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

// 出栈
func Pop(t *Node) (int, bool) {
	// 空栈
	if size == 0 {
		return 0, false
	}

	// 只有一个节点
	if size == 1 {
		size = 0
		stack = nil
		return t.Value, true
	}

	// 将栈的头节点指针指向下一个节点，并返回之前的头节点数据
	stack = stack.Next
	size--
	return t.Value, true
}

// 遍历（不删除任何节点，只读取值）
func traverse(t *Node) {
	if size == 0 {
		fmt.Println("空栈!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	stack = nil
	// 读取空栈
	v, b := Pop(stack)
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop() 失败!")
	}

	// 进栈
	Push(100)
	// 遍历栈
	traverse(stack)
	// 再次进栈
	Push(200)
	// 再次遍历栈
	traverse(stack)

	// 批量进栈
	for i := 0; i < 10; i++ {
		Push(i)
	}

	// 批量出栈
	for i := 0; i < 15; i++ {
		v, b := Pop(stack)
		if b {
			fmt.Print(v, " ")
		} else {
			// 如果已经是空栈，则退出循环
			break
		}
	}

	fmt.Println()
	// 再次遍历栈
	traverse(stack)
}