package main

import (
    "fmt"
    . "go-tutorial/chapter04/animal"
    "reflect"
)

/*
type IAnimal interface {
    GetName() string
    Call() string
    FavorFood() string
}*/

func main() {
    var animal = NewAnimal("中华田园犬")
    var pet = NewPet("泰迪")
    var any interface{} = NewDog(&animal, pet)
    if dog, ok := any.(Dog); ok {
        fmt.Println(dog.GetName())
        fmt.Println(dog.Call())
        fmt.Println(dog.FavorFood())
        fmt.Println(reflect.TypeOf(dog))
    }
    fmt.Println(reflect.TypeOf(any))
}
