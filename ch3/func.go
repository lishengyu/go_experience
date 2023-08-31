package main

import "fmt"

type FuncType func(int, int) int  // 声明一个函数类型

func Calc(a, b int , f FuncType) (result int) {
	result = f(a, b)
	return
}

func Add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	result := Calc(1, 1, Add)
	fmt.Println(result)

	var f FuncType = sub
	fmt.Println("result = ", f(10, 2))
}
