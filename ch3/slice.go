package main

import "fmt"

func main() {
	var s1 []int
	s2 := []int{}
	var s3 []int = make([]int, 0)
	s4 :=make([]int, 0, 0)
	s5 := []int{1,2,3}

	fmt.Println("=========================================")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Println("s5:", s5, "len:", len(s5), "cap:", cap(s5))

	fmt.Println("=========================================")
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array[:6:8]:", array[:6:8], "len:", len(array[:6:8]), "cap:", cap(array[:6:8]))
	fmt.Println("array[:6]:", array[:6], "len:", len(array[:6]), "cap:", cap(array[:6]))
	fmt.Println("array[5:]:", array[5:], "len:", len(array[5:]), "cap:", cap(array[5:]))
	fmt.Println("array[:]:", array[:], "len:", len(array[:]), "cap:", cap(array[:]))
	fmt.Println("array[8]:", array[8])

	fmt.Println("=========================================")
	str1 := array[2:5]
	str1[2] = 112
	fmt.Println("str1:", str1, "len:", len(str1), "cap:", cap(str1), "array[2:5]:", array[2:5])

	fmt.Println("=========================================")
	var str2 []int
	str2 = append(str2, 1)
	str2 = append(str2, 2, 3)
	str2 = append(str2, 11, 12, 13)
	fmt.Println(str2)

	fmt.Println("=========================================")
	c2 := make([]int, 5)
	c2 = append(c2, 6)
	fmt.Println(c2)

	fmt.Println("=========================================")
	c3 := []int{1, 2, 3}
	c3 = append(c3, 4, 5)
	fmt.Println(c3)

	fmt.Println("=========================================")
	// append
	aa := make([]int, 0, 1)
	fmt.Println("aa:", aa, "len:", len(aa), "cap:", cap(aa))
	bb := cap(aa)
	for i := 0; i < 50; i ++ {
		aa = append(aa, i)
		if n := cap(aa); n > bb {
			fmt.Printf("cap:%d -> %d\n", bb, n)
			fmt.Println("aa:", aa, "len:", len(aa), "cap:", cap(aa))
			bb = n
		}
	}

	fmt.Println("=========================================")
	// copy
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s11 := data[8:]
	s12 := data[:5]
	copy(s12, s11)
	fmt.Println("s11:", s11, "len:", len(s11))
	fmt.Println("s12:", s12, "len:", len(s12))
	fmt.Println("data:", data, "len:", len(data))
}
