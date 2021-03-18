package main

import (
	"fmt"
)

// 定义链表节点
type Node struct {
	Value int
	Next  *Node
}

// 初始化队列
var size = 0
var queue = new(Node)

// 入队（从队头插入）
func Push(t *Node, v int) bool {
	if queue == nil {
		queue = &Node{v, nil}
		size++
		return true
	}

	t = &Node{v, nil}
	t.Next = queue
	queue = t
	size++

	return true
}

// 出队（从队尾删除）
func Pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}

	// 迭代队列，直到队尾
	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}

	v := (temp.Next).Value
	temp.Next = nil

	size--
	return v, true
}

// 遍历队列所有节点
func traverse(t *Node) {
	if size == 0 {
		fmt.Println("空队列!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	// 入队
	Push(queue, 10)
	fmt.Println("Size:", size)
	// 遍历
	traverse(queue)

	// 出队
	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	// 批量入队
	for i := 0; i < 5; i++ {
		Push(queue, i)
	}
	// 再次遍历
	traverse(queue)
	fmt.Println("Size:", size)

	// 出队
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	// 再次出队
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	// 再次遍历
	traverse(queue)
}