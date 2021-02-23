package main

import "fmt"

type Integer int

type Math interface {
	Add(i Integer) Integer
	Multiply(i Integer) Integer
}

func (a Integer) Equal(b Integer) bool {
	return a == b
}

func (a Integer) Add(b Integer) Integer {
	return a + b
}

func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

func main() {
	var x Integer
	var y Integer
	x, y = 10, 15
	fmt.Println(x.Equal(y))
}
