package main

import "fmt"

var i = 2

func hello(i int) {
	fmt.Println(i)
}
func main() {
	/*i := 5
	// 输出的是 5 ，hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5.
	defer hello(i)
	i = i + 10*/

	// 同理
	i = 5
	defer fmt.Println(i)
	i++
	fmt.Println(i)

}
