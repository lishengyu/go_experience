package main

import (
	"fmt"
)

func modify(array [3]int) {
	array[0] = 111
	fmt.Println("In modify(), array values:", array)
}

func modify1(array *[3]int) {
	(*array)[0] = 222
	fmt.Println("in modify1(), array values:", *array)
}

func main() {
	var b [10]int
	for i := 0; i < 10; i++ {
		b[i] = i + 1
		fmt.Printf("b[%d] = %d\n", i, b[i])
	}

	fmt.Println("==============================================")
	for i, v := range b {
		fmt.Println("b[", i, "] =", v)
	}

	fmt.Println("==============================================")
	fmt.Println("len =", len(b), ", cap =", cap(b))

	fmt.Println("==============================================")
	a := [3]int{1,2,3}
	d := [3]int{1,2,3}
	c := [3]int{1,2}
	fmt.Println("if[a == d] =", a == d, ",if[d == c] =", d == c)

	var e [3]int
	e = d
	fmt.Println(e)

	fmt.Println("==============================================")
	modify(e)
	fmt.Println(e)

	fmt.Println("==============================================")
	modify1(&e)
	fmt.Println(e)
}
