package main

import (
	"fmt"
	. "go-tutorial/chapter04/animal"
)

func main() {
	animal := NewAnimal("中华田园犬")
	pet := NewPet("宠物狗")
	dog := NewDog(&animal, pet)

	fmt.Println(dog.GetName())
	fmt.Println(dog.Call())
	fmt.Println(dog.FavorFood())
}