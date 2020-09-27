package main

func main() {
	//编译报错
	/*str := "hello"
	str[0] = 'x'
	fmt.Println(str)*/

	//incr() 函数里的 p 是 *int 类型的指针，指向的是 main() 函数的变量 p 的地址。第 2 行代码是将该地址的值执行一个自增操作，incr() 返回自增后的结果。
	/*输出 p 为 2
	p := 1
	incr(&p)
	fmt.Println(p)*/

	add(1, 2, 3)
	add(5, 6, 9, 5, 1)
	add([]int{1, 5, 6}...)
	//调用失败，可变函数的知识
	//add([]int{1,5,6})
}

func incr(p *int) int {
	*p++
	return *p
}

func add(args ...int) int {

	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
