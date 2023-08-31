package main

import (
	"fmt"
)

// chan<-  只写
func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i  // 如果对方不读，会阻塞
	}
}

// <-chan 只读
func printer(in <-chan int) {
	for num := range in {
		fmt.Println("Printer:", num)
	}
}

func main() {
	c := make(chan int)	// 读写

	go counter(c)  // 生产者
	//go printer(c)	// 为什么这里没有输出？？
	printer(c)	// 消费者

	fmt.Println("done")
}
