package main

func main(){
	//编译错误
	/*a := 5
	b := 6.1
	fmt.Println(a + b)*/

	/*//截取新切片，a[i,j,k] 从i开始，j结束，容量为k-i   -->  输出 4
 	a := [5]int{1,2,3,4,5}
	t := a[3:4:4]
	fmt.Println(t[0])*/

	//Go 中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的组成部分，所以 a 和 b 是不同的类型，是不能比较的，所以编译错误
	/*a := [2]int{5, 6}
	b := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}*/
}