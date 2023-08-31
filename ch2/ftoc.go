package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g'F = %g'F\n", freezingF, fToC(freezingF))
	fmt.Printf("%g'F = %g'F\n", boilingF, fToC(boilingF))
}
// func 函数名(参数列表) 返回值列表 {函数体}
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
