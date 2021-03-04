package main

import (
    "fmt"
    . "go-tutorial/chapter04/animal"
    "reflect"
)

func main() {
    animal := NewAnimal("中华田园犬")
    pet := NewPet("泰迪")
    dog := NewDog(&animal, pet)

    // 返回的是 reflect.Type 类型值
    dogType := reflect.TypeOf(dog)
    fmt.Println("dog type:", dogType)

    // 返回的是 dog 指针对应的 reflect.Value 类型值
    dogValue := reflect.ValueOf(&dog).Elem()
    // 获取 dogValue 的所有属性
    fmt.Println("================ Props ================")
    for i := 0; i < dogValue.NumField(); i++ {
        // 获取属性名
        fmt.Println("name:", dogValue.Type().Field(i).Name)
        // 获取属性类型
        fmt.Println("type:", dogValue.Type().Field(i).Type)
        // 获取属性值
        fmt.Println("value:", dogValue.Field(i))
    }
    // 获取 dogValue 的所有方法
    fmt.Println("================ Methods ================")
    for j := 0; j < dogValue.NumMethod(); j++ {
        // 获取方法名
        fmt.Println("name:", dogValue.Type().Method(j).Name)
        // 获取方法类型
        fmt.Println("type:", dogValue.Type().Method(j).Type)
        // 调用该方法
        fmt.Println("exec result:", dogValue.Method(j).Call([]reflect.Value{}))
    }
}
