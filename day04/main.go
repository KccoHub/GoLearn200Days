package main

// 声明变量的简短模式有限制：
// 	1. 必须使用显示初始化
//	2. 不提供数据类型，编译器会自动推导
//	3. 只能在函数体内声明
var (
	/*size     := 1024
	max_size = size * 2*/
)

func main(){
	//这里的list是指针类型，不能在append上声明
	/*list := new([]int)
	list = append(list,1)*/

	//这里append，第二个参数不能直接用切片，可以后加...操作符，或写1，2，3
	/*list0 := []int{12,3,4}
	list1 := []int{3,9,5}
	list := append(list0,list1...)
	fmt.Println(list)*/

	//fmt.Println(max_size)

}
