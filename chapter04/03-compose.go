package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (a Animal) Call() string {
	return "动物的叫声..."
}

func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}

func (a Animal) GetName() string  {
	return a.Name
}

type Pet struct {
	Name string
}

func (p Pet) GetName() string  {
	return p.Name
}

type Dog struct {
	animal *Animal
	pet Pet
}

func (d Dog) FavorFood() string {
	return "骨头"
}

func (d Dog) Call() string {
	return "汪汪汪"
}

func main() {
	animal := Animal{"中华田园犬"}
	pet := Pet{"宠物狗"}
	dog := Dog{&animal, pet}

	fmt.Println(dog.animal.GetName())
	fmt.Print(dog.animal.Call())
	fmt.Println(dog.Call())
	fmt.Print(dog.animal.FavorFood())
	fmt.Println(dog.FavorFood())
}