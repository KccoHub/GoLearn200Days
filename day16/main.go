package main

import "fmt"

/*1.  a、b、c 的长度和容量分别是 0 3、2 3、1 2。
知识点：数组或切片的截取操作。截取操作有带 2 个或者 3 个参数，
形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组长度为 l。
在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。
操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。

func main(){
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]

	fmt.Println(s,a,b,c)
}*/

/*2.  在 A 处只声明了map m ,并没有分配内存空间，不能直接赋值，需要使用 make()，都提倡使用 make() 或者字面量的方式直接初始化 map。
B 处，v,k := m[“b”] 当 key 为 b 的元素不存在的时候，v 会返回值类型对应的零值，k 返回 false。
func main() {
	var m map[string]int        //A
	m["a"] = 1
	if v := m["b"]; v != nil {  //B
		fmt.Println(v)
	}
}*/

/* 建议改成这个样子
func main() {
	m := make(map[string]int)
	m["a"] = 1
	if v,ok := m["b"]; ok {
		fmt.Println(v)
	}
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

//知识点：接口的静态类型。a、b 具有相同的动态类型和动态值，分别是结构体 work 和 {3}；
//a 的静态类型是 A，b 的静态类型是 B，接口 A 不包括方法 ShowB()，接口 B 也不包括方法 ShowA()，编译报错。看下编译错误：
//a.ShowB undefined (type A has no field or method ShowB)
//b.ShowA undefined (type B has no field or method ShowA)
/*func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowB())
	fmt.Println(b.ShowA())
}*/
