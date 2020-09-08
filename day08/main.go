package main

import "fmt"

func hello() []string {
	return nil
}
//是将 hello() 赋值给变量 h，而不是函数的返回值，所以输出 not nil。
func main() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

/*func GetValue() int {
	return 1
}
// 结果：不能通过编译，考点：类型选择，类型选择的语法形如：i.(type)，其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型选择
func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}*/