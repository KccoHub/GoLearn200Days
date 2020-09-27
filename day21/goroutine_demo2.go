package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*在第 1 步，调度器开始运行 goroutine A，而 goroutine B 在运行队列里等待调度。
在第 2 步，调度器交换了 goroutine A 和 goroutine B。由于 goroutine A 并没有完成工作，因此被放回到运行队列。
在第 3 步，goroutine B 完成了它的工作并被系统销毁。这也让 goroutine A 继续之前的工作。
func main(){
	// 分配一个逻辑处理器给golang调度器使用
	runtime.GOMAXPROCS(1)

	// wg用来等待程序完成，计数2，表示等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutine...")

	// 声明一个函数，并创建一个goroutine
	go func(){
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("go!")
		}
		fmt.Println("goroutine1 over")
	}()

	go func(){
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("run!")
		}
		fmt.Println("goroutine2 over")
	}()

	fmt.Println("end goroutine.")
	wg.Wait()
	fmt.Println("Terminating Program")
}*/

/* 此处给一个逻辑处理器是为并发的最好解释，如果想要并行，分配两个逻辑处理器给调度器使用即可
输出说明：goroutine B 先执行，然后切换到 goroutine A，再切换到 goroutine B 运行至任务结束，
最后又切换到 goroutine A，运行至任务结束。注意，每次运行这个程序，调度器切换的时间点都会稍有不同*/
func main() {
	// wg用来等待程序完成
	var wg sync.WaitGroup

	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A", &wg)
	go printPrime("B", &wg)

	// 等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime 显示5000以内的素数值
func printPrime(prefix string, wg *sync.WaitGroup) {
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
