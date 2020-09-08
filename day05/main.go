package main

import "fmt"

func main() {
	//函数体可以用 ==/！=进行比较，但是不能比较大小
	//比较结构体时，如果每一项都相等，两个结构体才相等，否则不相等
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "q"}
	// sn1和sn3是不相等的，不能作为比较
	/*sn3:= struct {
		name string
		age  int
	}{age:11,name:"qq"}*/
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	//可比较的有Bool、数值型、字符、指针、数组
	//但 切片、map、函数是不可比较的
	/*sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}*/
}
