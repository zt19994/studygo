package main

import (
	"fmt"
	"time"
)

func main() {
	// go语言使用通道 channel  来保持资源同步，避免资源竞争问题
	// 使用通道不需要我们进行加解锁
	// A往通道里面写数据 B从通道里面读数据，go自动帮我们做好了数据同步

	// 使用make创建管道  channel 无缓冲
	numChan := make(chan int)

	go func() {
		for i := 0; i < 50; i++ {
			data := <-numChan
			fmt.Println("子goroutine1 读取数据data:", data)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			// 向管道中写入数据
			numChan <- i
			fmt.Println("子goroutine 写入数据：", i)
		}
	}()

	// 创建2个goroutine
	for i := 20; i < 50; i++ {
		numChan <- i
		fmt.Println("主goroutine，写入数据：", i)
	}

	time.Sleep(5 * time.Second)
}
