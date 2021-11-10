package main

import (
	"fmt"
	"time"
)

//这个将用于子goroutine使用
func display(num int) {
	count := 1
	for {
		fmt.Println("这是子goroutine程:", num, "当前count值:", count)
		count++
		//time.Sleep(1 * time.Second)
	}
}

func main() {
	// 启动子goroutine
	//go display()

	//启动子go程
	for i := 0; i < 3; i++ {
		go display(i)
	}

	// 匿名
	/*go func() {
		count := 1
		for {
			fmt.Println("这是子goroutine:", count)
			count++
			time.Sleep(1 * time.Second)
		}
	}()*/

	// 主goroutine
	count := 1
	for {
		fmt.Println("这是主goroutine:", count)
		count++
		time.Sleep(1 * time.Second)
	}

}
