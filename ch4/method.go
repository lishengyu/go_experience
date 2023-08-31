package main

import "fmt"

type MyInt int		// 自定义类型

// 面向对象
func (a MyInt) Add(b MyInt) MyInt {
	return a + b
}

// 面向过程
func Add(a, b MyInt) MyInt {
	return a + b
}

type Person struct {
	name string
	sex byte
	age int
}

// 给Person添加方法
func (p Person) PrintInfo() {
	fmt.Println(p.name, p.sex, p.age)
}

func main() {
	var a MyInt = 1
	var b MyInt = 1

	fmt.Println("===============================")
	fmt.Println("a.Add(b) = ", a.Add(b))
	fmt.Println("Add(a, b) = ", Add(a, b))

	fmt.Println("===============================")
	p := Person{"aaa", 'm', 18}
	p.PrintInfo()
}
