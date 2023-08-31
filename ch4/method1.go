package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//指针作为接收者，引用语义
func (p *Person) SetInfoPointer() {
	(*p).name = "yoyo"
	p.sex = 'f'
	p.age = 22
}

//值作为接收者，值语义
func (p Person) SetInfoValue() {
	p.name = "xxx"
	p.sex = 'm'
	p.age = 33
}

func main() {
	//p 为指针类型
	fmt.Println("====================================")
	var p *Person = &Person{"mike", 'm', 18}
	fmt.Println(p)
	p.SetInfoPointer() //func (p) SetInfoPointer()
	fmt.Println(p)
	p.SetInfoValue()    //func (*p) SetInfoValue()
	fmt.Println(p)
	(*p).SetInfoValue() //func (*p) SetInfoValue()
	fmt.Println(p)

	fmt.Println("====================================")
	var p1 Person = Person{"mike", 'm', 18}
	fmt.Println(p1)
	(&p1).SetInfoPointer() //func (p) SetInfoPointer()
	fmt.Println(p1)
	p1.SetInfoPointer() //func (p) SetInfoPointer()
	fmt.Println(p1)
	p1.SetInfoValue()    //func (*p) SetInfoValue()
	fmt.Println(p1)
	(&p1).SetInfoValue() //func (*p) SetInfoValue()
	fmt.Println(p1)
}
