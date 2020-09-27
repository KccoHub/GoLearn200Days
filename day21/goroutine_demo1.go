package main

import (
	"fmt"
	"time"
)

/* 使用time.Sleep()来等待协程结束，缺点是我们实际应用时，不清楚会等多久时间
func say(s string){
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main()  {
	go say("Hello, Mr.Zheng")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("over!")
}*/

//使用channel接收状态，缺点是执行多少次 done<-true 就要执行多少次 <-done

func main() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("hello, Mr.Zheng")
		}
		done <- true
	}()
	<-done
	fmt.Println("over!")
}

/*func say(s string, waitGroup *sync.WaitGroup){
	defer waitGroup.Done()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	say("Hello", &wg)
	say("world", &wg)
	fmt.Println("over!")
}*/
