package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	t1 := make(map[string]interface{})
	t1["company"] = "itcast"
	t1["subjects"] = []string{"Go", "C++", "python", "Test"}
	t1["isok"] = true
	t1["price"] = 666.666

	//b, err := json.Marshal(t1)
	b, err := json.MarshalIndent(t1, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(string(b))
}
