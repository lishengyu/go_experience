package main

import "fmt"

func main() {
	fmt.Println("============================================")
	var m1 map[int]string		// 声明一个map，没有初始化;
	fmt.Println(m1 == nil)

	m2 := map[int]string{1: "abc", 2: "def"}
	m3 := make(map[int]string)
	fmt.Println(m2, m3)

	m4 := make(map[int]string, 10)
	fmt.Println("map:", m4)

	fmt.Println("============================================")
	m2[1] = "xxx"
	m2[3] = "zzz"			// 超过范围，会自动追加
	fmt.Println("m2:", m2)

	fmt.Println("============================================")
	for k, v := range m2 {
		fmt.Println("key:", k, "value:", v)
	}
	fmt.Println(m2)

	for _, v := range m2 {
		fmt.Println("value:", v)
	}

	fmt.Println("============================================")
	// check
	v1, ok := m2[1]
	fmt.Println("value:", v1, "ok:", ok)

	v2, ok2 := m2[5]
	fmt.Println("value:", v2, "ok2:", ok2)

	fmt.Println("============================================")
	// delete
	for k, v := range m2 {
		fmt.Printf("%d --> %s\n", k, v)
	}

	delete(m2, 2)
	fmt.Println("del 2 from m2")
	for k, v := range m2 {
		fmt.Printf("%d --> %s\n", k, v)
	}

}
