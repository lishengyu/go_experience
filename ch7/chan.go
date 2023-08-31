package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子协程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("子协程[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Printf("主协程num[%d]:%d\n", i, num)
	}

	fmt.Println("main协程结束")
}
