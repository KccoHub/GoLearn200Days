package main

import "fmt"

/*func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}
// FDM
func main() {
	f()
	fmt.Println("M")
}
*/

type Person struct {
	age int
}

/**
1处.person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；

2处.defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后面的函数执行的时候取出仍是 28；

3处.闭包引用，person 的值已经被改变，指向结构体 Person{29}，所以输出 29*/
func main() {
	person := &Person{28}

	// 1.28
	defer fmt.Println(person.age)

	// 2.28
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.29
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}
