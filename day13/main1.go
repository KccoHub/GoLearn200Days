package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

// 打印出：
//showA
//showB
// 这道题可以结合第 12 天的第三题一起看，Teacher 没有自己 ShowA()，所以调用内部类型 People 的同名方法，
// 需要注意的是第 5 行代码调用的是 People 自己的 ShowB 方法
func main() {
	t := Teacher{}
	t.ShowA()
}
