package main

import (
	"log"
)

type user struct {
	name string
	age  int
}

func filterAge(users []user) interface{} {
	var slice []user
	for _, u := range users {
		if u.age >= 18 && u.age <= 35 {
			slice = append(slice, u)
		}
	}
	return slice
}

func mapAgeToSlice(users []user) interface{} {
	var slice []int
	for _, u := range users {
		slice = append(slice, u.age)
	}
	return slice
}

func sumAge(users []user, pipes ...func([]user) interface{}) int {
	var ages []int
	var sum int
	for _, f := range pipes {
		result := f(users)
		switch result.(type) {
		case []user:
			users = result.([]user)
		case []int:
			ages = result.([]int)
		}
	}
	if len(ages) == 0 {
		log.Fatalln("没有在管道中加入 mapAgeToSlice 方法")
	}
	for _, age := range ages {
		sum += age
	}
	return sum
}

func main() {
	var users = []user{
		{
			name: "张三",
			age: 18,
		},
		{
			name: "李四",
			age: 22,
		},
		{
			name: "王五",
			age: 20,
		},
		{
			name: "赵六",
			age: -10,
		},
		{
			name: "孙七",
			age: 60,
		},
		{
			name: "周八",
			age: 10,
		},
	}

	sum := sumAge(users, filterAge, mapAgeToSlice)
	log.Printf("用户年龄累加结果: %d\n", sum)
}
