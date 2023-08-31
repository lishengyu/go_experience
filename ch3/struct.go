package main

import "fmt"

type St struct {
	id int
	name string
	age int
	addr string
}

func main() {
	fmt.Println("======================================")
	// 顺序初始化，必须每个成员都初始化
	var s1 St = St{1, "mike", 18, "sz"}
	s2 := St{2, "aaa", 20, "wh"}
	// 指定初始化某个成员，没有初始化的成员为零值
	s3 := St{id: 3, name: "hz"}

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	fmt.Println("======================================")
	var s5  *St = &St{5, "weibo", 10, "bj"}
	s6 := &St{6, "tencent", 20, "sz"}
	fmt.Println("s5:", *s5)
	fmt.Println("s6:", *s6)

	fmt.Println("======================================")
	fmt.Printf("id=%d, name=%s, age=%d, addr=%s\n", s1.id, s1.name, s1.age, s1.addr)
	s1.age = 99
	s1.addr = "shenzhen"
	fmt.Printf("id=%d, name=%s, age=%d, addr=%s\n", s1.id, s1.name, s1.age, s1.addr)

	fmt.Println("======================================")
	var  s7 *St = &s1
	s7.id = 77
	(*s7).name = "nnn"
	fmt.Println(s7, *s7, s1)

	fmt.Println("======================================")
	fmt.Println("s1 == s7", s1 == *s7)
	fmt.Println("s1 == s2", s1 == s2)
}
