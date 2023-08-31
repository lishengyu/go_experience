package main

import (
	"fmt"
)

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	var a int = 10
	var p *int = nil

	fmt.Printf("============================================\n")
	p = &a
	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("p = %p, &p = %p, *p = %d\n", p, &p, *p)

	fmt.Printf("============================================\n")
	*p = 111
	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("p = %p, &p = %p, *p = %d\n", p, &p, *p)

	fmt.Printf("============================================\n")
	var p1 *int = nil
	p1 = new(int)
	fmt.Println("*p1 = ", *p1)

	p2 := new(int)
	*p2 = 111
	fmt.Println("*p2 = ", *p2)

	fmt.Printf("============================================\n")
	val1 := 10
	val2 := 20
	swap(&val1, &val2)
	fmt.Printf("val1= %d, val2 = %d\n", val1, val2)
}
