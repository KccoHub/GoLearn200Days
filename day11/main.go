package main

import "fmt"

/*
* 当接口的动态类型和动态值都为 null 时，其接口类型值才为 nil
func main(){
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}*/

// 输出为 0 ，因为删除不存在的map时，不会报错，相当于没有进行任何操作，所以返回 0
func main() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}
