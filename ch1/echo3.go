// Echo1 打印自身命令行参数

package main

import "fmt"

func Test() {
	fmt.Println("This is a test func")
}

func Max(n1 int, n2 int)(min int, max int) {
	if n1 > n2 {
		max = n1
		min = n2
	} else {
		max = n2
		min = n1
	}
	return 
}

func main() {
	str2 := `hello
	123 \n \r 哈哈
	`
	fmt.Println("str2 = ", str2)

	var  v1 complex64
	v1 = 3.2 + 12i

	fmt.Println(v1, real(v1), imag(v1))

	fmt.Printf("v1 = %t\n", true)

	var v int
	fmt.Println("请输入一个整形：")
	fmt.Scanf("%d", &v)
	fmt.Println("v = ", v)

	Test()
	min, max := Max(12, 22)
	fmt.Printf("min = %d, max = %d\n", min, max)
}
