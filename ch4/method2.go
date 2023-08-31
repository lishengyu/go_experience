package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

func (p *Person) PrintInfo() {
	fmt.Printf("Person: %s, %c, %d\n", p.name, p.sex, p.age)
}

type Student struct {
	Person	// 匿名字段
	id int
	addr string
}

func (s *Student) PrintInfo() {
	fmt.Printf("Student: %s, %c, %d\n", s.name, s.sex, s.age)
}

func main() {
	// 继承
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	s := Student{Person{"yoyo", 'f', 20}, 2, "sz"}
	s.PrintInfo()

	fmt.Println("=================================")
	// 重写
	s.PrintInfo()
	s.Person.PrintInfo()

	fmt.Println("=================================")
	// 方法值
	pFunc1 := p.PrintInfo
	pFunc1()

	fmt.Println("=================================")
	// 方法表达式
	// 显式传递参数
	pFunc2 := (*Person).PrintInfo
	pFunc2(&p)
}
