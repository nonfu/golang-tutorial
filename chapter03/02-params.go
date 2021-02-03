package main

import "fmt"

func add(a, b *int) int {
	*a *= 2
	*b *= 3
	return *a + *b
}

func main() {
	x, y := 1, 2
	z := add(&x, &y)
	fmt.Printf("add(%d, %d) = %d\n", x, y, z)
}
