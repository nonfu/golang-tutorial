package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("Hello, World!")
	hash := md5.Sum(data)
	fmt.Printf("原始值: %s\n", data)
	fmt.Printf("MD5值: %x\n", hash)
}
