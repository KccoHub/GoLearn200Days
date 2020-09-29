package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c) // 注释的话就会被阻塞
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}
