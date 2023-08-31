package main

import "fmt"
import "time"

func main() {
	fmt.Println("=================================")
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Println("i:", i)
		}
	}

	fmt.Println("=================================")
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()

	c <- 666
	<-o
}
