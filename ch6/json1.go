package main

import (
	"fmt"
	"encoding/json"
)

type IT struct {
	// Company 不会导出到JSON中
	Company string `json:"-"`

	// Subjects的值会进行二次JSON编码
	Subjects []string `json:"subjects"`

	// 转换为字符串，再输出
	IsOk bool `json:",string"`

	// 如果Price为空，则不输出到JSON串中
	Price float64 `json:"price, omitempty"`
}

func main() {
	t1 := IT{Company:"itcast", Subjects:[]string{"Go", "C++", "Python", "Rust"}, IsOk:true}

	b, err := json.Marshal(t1)

	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(string(b))
}
