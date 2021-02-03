package main

func main() {
	p1 := new(int)    // 返回 int 类型指针，相当于 var p1 *int
	p2 := new(string) // 返回 string 类型指针
	p3 := new([3]int) // 返回数组类型指针，数组长度是 3

	type Student struct {
		id    int
		name  string
		grade string
	}
	p4 := new(Student) // 返回对象类型指针

	println("p1: ", p1)
	println("p2: ", p2)
	println("p3: ", p3)
	println("p4: ", p4)

	s1 := make([]int, 3)          // 返回初始化后的切片类型值，即 []int{0, 0, 0}
	m1 := make(map[string]int, 2) // 返回初始化的字典类型值，即散列化的 map 结构

	println(len(s1)) // 3
	for i, v := range s1 {
		println(i, v)
	}

	println(len(m1)) // 0
	m1["test"] = 100
	for k, v := range m1 {
		println(k, v)
	}
}
