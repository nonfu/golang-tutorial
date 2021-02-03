package main

import "fmt"

func main() {
	var testMap map[string]int
	testMap = map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	k := "two"
	v, ok := testMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}

	fmt.Println(testMap)
}
