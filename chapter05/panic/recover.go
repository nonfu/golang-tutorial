package main

import (
    "fmt"
)

func divide() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Printf("捕获到运行时 panic: %v\n", err)
        }
    }()

    var i = 1
    var j = 1
    k := i / j
    fmt.Printf("%d / %d = %d\n", i, j, k)
}

func main() {
    divide()
    fmt.Println("divide 方法调用完毕，回到 main 函数")
}