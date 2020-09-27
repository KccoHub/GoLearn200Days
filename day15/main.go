package main

import "fmt"

/* 1. nil切片和空切片
func main(){
	s1 为 nil，s2 为 not nil
	var s1 []int
	//var s2 = []int{}

	if s1 == nil{
		fmt.Println("yes nil")
	}else{
		fmt.Println("no nil")
	}
}*/

/*输出为A
func main()  {
	i := 65
	fmt.Println(string(i))
}
*/

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

//知识点：接口。一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。
func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}
