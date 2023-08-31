package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)

	i := 0
	go func() {
		for {
			<-ticker.C
			i++
			fmt.Println("i = ", i)

			if i == 5 {
				ticker.Stop()
			}
		}
	}()

	for {

	}
}
