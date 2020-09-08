package main

import "fmt"

/*type person struct {
	name string
}

func main(){

	//1 & 2 & 3 都是 chan 的正确使用方式，4 不是；写 channel 必须带上值，所以 4 错误
	//var ch0 chan int		// 1
	//ch1 := make(chan int)	// 2
	//<- ch0					// 3
	//ch1 <-					// 4

	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])//输出为 0
}*/

func hello(num ...int) {
	num[0] = 18
}

func main() {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])//输出 18，可变参数
}
