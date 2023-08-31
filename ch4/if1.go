package main

import "fmt"

type Element interface{}

type Person struct {
	name string
	age int
}

type Mystr string

func main() {
	list := make([]Element, 4)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"yoyo", 18}
	var s Mystr = "well"
	list[3] = s

	fmt.Println("============================================")
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is :%d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is string and value is :%s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is Person and value is :[%s, %d]\n", index, value.name, value.age)
		} else {
			fmt.Printf("list[%d] is an unkown type\n", index)
		}
	}

	fmt.Println("============================================")
	for index, element := range list {
		switch value := element.(type) {
			case int:
				fmt.Printf("list[%d] type:int, value:%d\n", index, value)
			case string:
				fmt.Printf("list[%d] type:string, value:%s\n", index, value)
			case Person:
				fmt.Printf("list[%d], type:Person, value:[%s, %d]\n", index, value.name, value.age)
			default:
				fmt.Printf("list[%d] is unkown type\n", index)
		}
	}
}
