package main

import "fmt"

type Person struct {
	name	string
	sex	byte
	age	int
}

type Student struct {
	Person		// 匿名字段
	id	int
	addr	string
}

func main() {
	s1 := Student{Person{"aaa", 'm', 18}, 1, "sz"}
	fmt.Printf("s1 = %+v\n", s1)

	s2 := Student{Person: Person{"aaa", 'm', 18}, addr: "sz"}
	fmt.Printf("s2 = %+v\n", s2)

	s1.name = "ccc"
	s1.addr = "hz"
	fmt.Println(s1)
}
