package main

import "fmt"
// 基于类型 int 创建了新类型 MyInt1
type MyInt1 int
// 创建了 int 的类型别名 MyInt2，注意类型别名的定义时 =
type MyInt2 = int

func main() {
	var i int =0
	// 将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过
	var i1 MyInt1 = i //可以将上面的定义 int 改为MyInt1 ，这里就可以通过编译了
	// MyInt2 只是 int 的别名，本质上还是 int，可以赋值。
	var i2 MyInt2 = i
	fmt.Println(i1,i2)
}