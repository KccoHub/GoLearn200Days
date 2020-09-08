package main

import "fmt"

func main(){
	/*1. make关键字的用法
	s := make([]int,10)
	s = append(s,1,2,3)
	fmt.Println(s)

	//结果是：【0 0 0 0 0 0 0 0 0 0 1 2 3】
	*/

	/*2. 续
	s := make([]int,0)
	s = append(s,1,2,3)
	fmt.Println(s)

	//结果是：【1 2 3】
	*/

	//new和Make的区别
	newAndMake()
}
// 返回值的命名规则，可以不用命名，如果命名必须全部命名，如：(i int,err error)；且如果只有一个返回值时，有命名的话也需要加上括号
func funcMui(x,y int)(int,error){
	return x+y,nil
}

// make ---> slice hash channel  	*make 关键字的作用是创建切片、哈希表和 Channel 等内置的数据结构。
// new  ---> pointer				*new  关键字的作用是为类型申请一片内存空间，并返回指向这片内存的指针。
func newAndMake(){
	i := new(int)
	fmt.Println(i)

	var v  int
	i = &v
	fmt.Println(i)
	// 以上两个不同的初始化方法是等价的
	// 只能接收一个类型作为参数然后返回一个指向该类型的指针：

	slice := make([]int64,5,3)
	hash := make(map[string]int)
	ch := make(chan string,5)
	fmt.Println(slice,hash,ch)
	// slice 是一个包含 data、cap 和 len 的私有结构体 internal/reflectlite.sliceHeader；
	// hash 是一个指向 runtime.hmap 结构体的指针；
	// ch 是一个指向 runtime.hchan 结构体的指针；

	//in := new(int)
	//u := new(user)
}
