package main

import "fmt"

const (    // iota 被重置为 0
	c0 = iota   // c0 = 0
	c1   // c1 = 1
	c2   // c2 = 2
)

const (
	u = iota * 2  // u = 0
	v  // v = 2
	w  // w = 4
)

const x = iota  // x = 0
const y = iota  // y = 0

func main() {
	fmt.Printf("x:%v\n", x)
	fmt.Printf("y:%v\n", y)
}