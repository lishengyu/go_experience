package main

import "fmt"
import "os"

func main() {
	args := os.Args		//获取命令行参数

	if args == nil || len(args) < 2 {
		fmt.Println("err: xx ip port")
		return 
	}

	ip := args[1]
	port := args[2]

	fmt.Printf("ip = %s, port = %s\n", ip, port)
}
