package main

import (
	"fmt"
	"encoding/json"
)

type IT struct {
	Company string `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk bool `json:"isok"`
	Price float64 `json:"price"`
}

func main() {
	b := []byte(`{
		"company":"itcast",
		"subjects":[
			"Go",
			"C++",
			"Python",
			"Rust"
		],
		"isok":true,
		"price":666.666
	}`)

	fmt.Println("===================================")
	// 解析到结构体
	var t IT
	err := json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(t)

	fmt.Println("===================================")
	// 解析到interface
	var i interface{}
	err1 := json.Unmarshal(b, &i)
	if err1 != nil {
		fmt.Println("json err1:", err1)
	}
	fmt.Println(i)

	fmt.Println("===================================")
	m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case bool:
			fmt.Println(k, "is bool", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type unkown")
		}
	}
}
