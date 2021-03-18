package main

import "fmt"

func main() {
    defer func() {
        fmt.Println("代码清理逻辑")
    }()

    var i = 1
    var j = 0
    if j == 0 {
        panic("除数不能为0！")
    }
    k := i / j
    fmt.Printf("%d / %d = %d\n", i, j, k)
}