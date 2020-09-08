package main

import "fmt"

/*func main(){
	str0 := "a" + "b"
	str1 := 'c' + 'd'
	str2 := fmt.Sprintf("a%d",1)
	// value: ab + Type: string
	fmt.Printf("value: %v + Type: %T\n",str0,str0)
	// value: 199 + Type: int32
	fmt.Printf("value: %v + Type: %T\n",str1,str1)
	// value: a1 + Type: string
	fmt.Printf("value: %v + Type: %T\n",str2,str2)
}*/

//0 2 zz zz 5
/*const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)*/

// 0 0 zz zz 4
const x = iota
const (
	y = iota
	z = "zz"
	k
	_
	p = iota
)
// iota 在 const 关键字出现时将被重置为 0 ( const 内部的第一行之前)
// iota 是 golang 语言的常量计数器,只能在常量的表达式中使用。
func main()  {
	// nil 只能赋值给指针、chan、interface、切片、map、func类型
	/*var x0 = nil
	var x1 interface{} = nil
	var x2 string = nil
	var x3 error = nil*/
	fmt.Println(x,y,z,k,p)
}


