package main

import (
	"fmt"
	"sync"
)

//channel练习
// 1. 启动一个goroutine，生成100个随机数发送到ch1
// 2. 启动一个goroutine，从ch1通道中读取值，然后计算其平方，放到ch2中
// 3. 在main中，从ch2取值打印出来

/*
知识点：
1、开启多个goroutine，执行同一个函数的时候，需要开启 sync.Once ，保证某个操作只执行一次
		var once sync.Once
		once.Do(func() { close(ch2) })

2、开启多个goroutine的时候，需要开启线程等待。
	var wg sync.WaitGroup
	wg.Done()
	wg.Add(3)
	wg.Wait()

3、只读、只写chan
*/

// var once sync.Once
// 往 ch 里面塞 10 个值
var once sync.Once

func f1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

// 从ch里面拿所有值出来，然后进行平方处理，放到 cha 中
func f2(ch, cha chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		cha <- v * v
	}
	once.Do(func() { close(cha) })
}

func main() {
	var wg sync.WaitGroup
	var ch = make(chan int, 10)
	var cha = make(chan int, 10)
	wg.Add(3)
	go f1(ch, &wg)
	go f2(ch, cha, &wg)
	go f2(ch, cha, &wg)
	for ret := range cha {
		fmt.Println(ret)
	}
}
