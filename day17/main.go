package main

import "fmt"

/*func main(){
	var a int
	a, b := tttt(a)
	fmt.Println(a, b)
}


func tttt(a int)(int, int){
	return a, a+2
}
*/

/*知识点：defer、返回值。注意一下，increaseA() 的返回参数是匿名，increaseB() 是具名。
关于 defer 与返回值的知识点，后面我会写篇文章详细分析，到时候可以看下文章的讲解。
func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}*/

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

// 类型断言。这道题可以和第 15 天的第三题 和第 16 天的第三题结合起来看
func main() {
	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
