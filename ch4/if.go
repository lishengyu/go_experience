package main

import "fmt"

// 接口定义，只有声明，没有实现
type Ifer interface {
	SayHi()
}

type Per interface {
	Ifer
	Sing(lyrics string)
}

// St结构体
type St struct {
	name string
	score float64
}

// St结构体实现SayHi()方法
func (s *St) SayHi() {
	fmt.Printf("St[%s %f] say hi!!!\n", s.name, s.score)
}

func (s *St) Sing(lyrics string) {
	fmt.Printf("St sing[%s]\n", lyrics)
}

// Te结构体
type Te struct {
	name string
	group string
}

// Te结构体实现SayHi()方法
func (t *Te) SayHi() {
	fmt.Printf("Te[%s %s] say hi!!!\n", t.name, t.group)
}

// 自定义类型
type Mystr string

// 自定义类型实现SayHi()方法
func (str Mystr) SayHi() {
	fmt.Printf("Mystr[%s] say hi!!!\n", str)
}

// 普通函数
func WhoSayHi(i Ifer) {
	i.SayHi()
}

func main() {
	s := &St{"aaa", 88.88}
	t := &Te{"hhh", "Go"}
	var str Mystr = "测试"

	fmt.Println("==================================")
	s.SayHi()
	t.SayHi()
	str.SayHi()

	fmt.Println("==================================")
	// 多态，调用同一接口，不同表现
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(str)

	fmt.Println("==================================")
	x := make([]Ifer, 3)
	x[0], x[1], x[2] = s, t, str
	for _, value := range x {
		value.SayHi()
	}

	fmt.Println("==================================")
	var i2 Per
	i2 = s
	i2.SayHi()
	i2.Sing("lalalalalalaalal")
}
