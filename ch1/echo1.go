// Echo1 打印自身命令行参数

package main

import "fmt"
import "os"

func main() {
	//var定义了两个string类型的变量s和sep
	var s, sep string
  for i :=1; i < len(os.Args); i++ {
		// 连接字符串sep和os.Args[i]
		s += sep + os.Args[i]
		sep = " "
	}
  fmt.Println(s)
}
