package main

import (
    "fmt"
    "reflect"
)

type Container struct {
    s reflect.Value
}

// 通过传入存储元素类型和容量来初始化容器
func NewContainer(t reflect.Type, size int) *Container {
    if size <= 0  {
        size = 64
    }
    // 基于切片类型实现这个容器，这里通过反射动态初始化这个底层切片
    return &Container{
        s: reflect.MakeSlice(reflect.SliceOf(t), 0, size),
    }
}

// 添加元素到容器，通过空接口声明传递的元素类型，表明支持任何类型
func (c *Container) Put(val interface{})  error {
    // 通过反射对实际传递进来的元素类型进行运行时检查，
    // 如果与容器初始化时设置的元素类型不同，则返回错误信息
    if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
        return fmt.Errorf("put error: cannot put a %T into a slice of %s",
            val, c.s.Type().Elem())
    }
    // 如果类型检查通过则将其添加到容器中
    c.s = reflect.Append(c.s, reflect.ValueOf(val))
    return nil
}

// 从容器中读取元素，将返回结果赋值给 val，同样通过空接口指定元素类型
func (c *Container) Get(val interface{}) error {
    // 还是通过反射对元素类型进行检查，如果不通过则返回错误信息
    if reflect.ValueOf(val).Kind() != reflect.Ptr ||
        reflect.ValueOf(val).Elem().Type() != c.s.Type().Elem() {
        return fmt.Errorf("get error: needs *%s but got %T", c.s.Type().Elem(), val)
    }
    // 将容器第一个索引位置值赋值给 val
    reflect.ValueOf(val).Elem().Set( c.s.Index(0) )
    // 然后删除容器第一个索引位置值
    c.s = c.s.Slice(1, c.s.Len())
    return nil
}

func main() {
    nums := []int{1, 2, 3, 4, 5}

    // 初始化容器，元素类型和 nums 中的元素类型相同
    c := NewContainer(reflect.TypeOf(nums[0]), 16)

    // 添加元素到容器
    for _, n := range nums {
        if err := c.Put(n); err != nil {
            panic(err)
        }
    }

    // 从容器读取元素，将返回结果初始化为 0
    num := 0
    if err := c.Get(num); err != nil {
        panic(err)
    }

    // 打印返回结果值
    fmt.Printf("%v (%T)\n", num, num)
}