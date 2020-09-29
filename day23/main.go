package main

import "fmt"

//切片的底层是数组，当使用s1[1:]时获得切片s2，此时两个切片用的是一个底层数组，这会导致改变s2会影响s1
func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	s2[1] = 4
	fmt.Println(s1) //1 2 4 4 5 6
	fmt.Println(s2) //2 4
	//新加的数组，如果加上后的总长度不超过原始切片的长度，则不新new一个数组，如果长度超过了原始切片长度，则s2存储新的数组上
	s2 = append(s2, 5, 6, 7, 8)
	fmt.Println(s1) //1 2 4 5 6 7
	fmt.Println(s2) //2 4 5 6 7 8
}
