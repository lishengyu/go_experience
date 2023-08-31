package main

import (
	"fmt"
	"encoding/json"
)

type IT struct {
	Company string
	Subjects []string
	IsOk	bool
	Price	float64
}

func main() {
	t1 := IT{"itcast", []string{"Go", "C++", "Python", "Rust"}, true, 666.666}

	b, err := json.MarshalIndent(t1, "", "  ")

	if err != nil {
		fmt.Println("json err:", err)
	}

	//fmt.Println(b)
	fmt.Println(string(b))
}
