package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("==================写文件=================")
	// 创建文件
	fout, err := os.Create("./123.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 延迟到函数结束前，关闭文件
	//defer fout.Close()

	for i := 0; i < 5; i++ {
		// 给定不同的字符串
		outstr := fmt.Sprintf("%s:%d\n", "Hello go", i)
		// 写入string信息
		fout.WriteString(outstr)
		// 写入byte信息
		fout.Write([]byte("abcd\n"))
	}
	fout.Close()

	fmt.Println("==================读文件=================")
	fin, err := os.Open("./123.txt")
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 1024)     // 开辟1024个字节的slice作为缓冲
	for{
		n, _ := fin.Read(buf)  // 读取文件
		if n == 0 {             // 文件结束
			break
		}

		fmt.Println("n:", n ,string(buf)) // 打印读取的内容
	}
	fin.Close()
}
