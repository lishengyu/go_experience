package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子协程结束")

		for i:= 0; i < 3; i++ {
			c <- i
			fmt.Printf("子协程正在运行[%d]:len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
		close(c)
	} ()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	/*
	for i:= 0; i < 3; i ++ {
		num := <-c
		fmt.Printf("main协程正在运行num[%d] = %d\n", i, num)
	}

	time.Sleep(2 *time.Second)
	*/
	fmt.Println("main协程结束")
}
