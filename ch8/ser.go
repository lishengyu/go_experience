package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func dealConn(conn net.Conn) {
	// 函数结束时，关闭套接字
	defer conn.Close()

	// client的网络地址
	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, "连接成功")

	// 缓冲区 用于接收client发送的数据
	buf := make([]byte, 1024)

	for {
		// 阻塞等待client发送的数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 切片截取读到的数据
		result := buf[:n]
		fmt.Printf("接收数据[%s] ==> [%d]:%s\n", ipAddr, n, string(result))
		// 如果发送的内容是exit，退出连接
		if "exit" == string(result) {
			fmt.Println(ipAddr, "退出连接")
			return
		}
		// 将接收到的数据转换为大写，发送给client
		conn.Write([]byte(strings.ToUpper(string(result))))
	}
}

func main() {
	// 创建 监听socket
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	// 延迟关闭
	defer listenner.Close()

	for {
		// 阻塞等待客户端连接
		conn, err := listenner.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go dealConn(conn)
	}
}
