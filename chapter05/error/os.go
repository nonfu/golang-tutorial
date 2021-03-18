package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // 获取指定路径文件信息，对应类型是 FileInfo
    // 如果文件不存在，则返回 PathError 类型错误
    fi, err := os.Stat("add.go")
    if err != nil {
        switch err.(type) {
        case *os.PathError:
            fmt.Println(err)
        case *os.LinkError:
            fmt.Println(err)
        case *os.SyscallError:
            fmt.Println(err)
        case *exec.Error:
            fmt.Println(err)
        }
    } else {
        fmt.Print(fi)
    }
}
