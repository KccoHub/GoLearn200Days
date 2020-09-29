package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go func() {
		c <- 4 + 5
		c <- 4 - 3
	}()
	i := <-c
	fmt.Println(i)
}
