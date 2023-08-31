package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// client 连接到 server
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	// 缓冲区
	buf := make([]byte, 1024)
	for {
		fmt.Printf("请输入发送的内容：")
		fmt.Scan(&buf)
		fmt.Printf("发送:%s\n", string(buf))

		// 发送数据
		conn.Write(buf)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		result := buf[:n]
		fmt.Printf("接收[%d]:%s\n", n, string(result))
	}
}
