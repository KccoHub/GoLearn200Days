package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 将切片存储的数据存储到 map[int]*int 类型里面，并取出
// 错误写法，原由：因为 &val 的地址是不会变的，所以 m1[key] = &val 中，
// 在 range 的过程中，将最后一个 &val 地址上的值赋值给所有的 m1[key] 上了。
func main() {
	sortMap(make(map[string]int))
	slice := []int{0, 1, 2, 3}
	m1 := make(map[int]*int)

	for key, val := range slice {
		m1[key] = &val
	}
	for k, v := range m1 {
		fmt.Println(k, "-1->", *v)
	}

	fmt.Println("--------------------------")
	m2 := correct(slice)
	for k, v := range m2 {
		fmt.Println(k, "-2->", *v)
	}
}

// 正确写法：  (注意：遍历range出来的数据与插入的顺序无关)
func correct(slice []int) map[int]*int {
	m := make(map[int]*int)
	for key, value := range slice {
		val := value
		m[key] = &val
	}
	return m
}

// 按照指定顺序遍历 map
// 将 map 中的key 进行排序，再将key存入切片中，将切片进行排序，再输出
func sortMap(map1 map[string]int) {
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d",i)
		val := rand.Intn(100)
		map1[key] = val
	}

	keys := make([]string, 0, 100)
	for key := range map1 {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for index,value := range keys{
		fmt.Println(value,map1[value])
		fmt.Println(index)
	}
}
