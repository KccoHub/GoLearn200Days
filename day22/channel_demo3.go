package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, value := range s {
		sum += value
	}
	fmt.Println("sum =", sum)
	c <- sum
}

func main() {
	s := []int{1, 3, 5, 7, 9, -10, 1}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	i, j := <-c, <-c

	fmt.Println("i =", i, "j =", j)
}
