package main

import "fmt"

func main() {
	var c = make(chan int, 4)

	go chann(c)
	for v := range c {
		fmt.Println(v)
		v, ok := <-c
		fmt.Println("v :", v, "  ok :", ok)

		/*
			if ok == true {
				close(c)
			}*/
	}
}

func chann(c chan int) {
	c <- 5
	c <- 4
	c <- 2
	c <- 1
}
