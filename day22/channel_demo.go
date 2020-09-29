package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	//切片类型是一个int型，映射的类型是键为int型，值为*int，即值是一个地址。
	myMap := make(map[int]*int)

	for index, value := range slice {
		temp := value //没有这一步，直接返回&value，返回值非预期
		// 因为for range创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，就会导致错误
		myMap[index] = &temp
	}
	fmt.Println("===================")
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for index, value := range myMap {
		//返回地址的引用 *
		fmt.Printf("map[%v] : %v\n", index, *value)
	}
}
