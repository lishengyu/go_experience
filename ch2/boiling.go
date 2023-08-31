package main

import "fmt"

//可被外部引用
const boilingF = 212.0

func main() {
	//只可在函数内部使用
	var f = boilingF
	var c = (f -32) * 5 / 9
	fmt.Printf("boiling point = %g'F or %g'C\n", f, c)
}
