package main

import "fmt"

func main() {
	var a int = 3
	if a == 3 {
		fmt.Println("a==3")
	}

	if b :=3; b == 3 {
		fmt.Println("b==3")
	}
}
