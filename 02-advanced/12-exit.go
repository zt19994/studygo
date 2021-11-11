package main

import (
	"fmt"
	"runtime"
	"time"
)

//return  ===> 返回当前函数
//exit ===> 退出当前进程
//Goexit ===> 提前退出当前goroutine

func main() {
	go func() {
		go func() {
			func() {
				fmt.Println("子goroutine内部的函数")
				//return // 这是返回当前函数
				//os.Exit(-1) // 退出进程
				runtime.Goexit() // 提前退出当前goroutine
			}()

			fmt.Println("子goroutine结束！")
			fmt.Println("go 222222")
		}()
		time.Sleep(2 * time.Second)
		fmt.Println("go 111111")
	}()

	fmt.Println("这是主goroutine")
	time.Sleep(3 * time.Second)
	fmt.Println("Over!")
}
