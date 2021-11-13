package main

import (
	"fmt"
	"time"
)

func main() {
	// 单向读通道
	//var numChanReadOnly <- chan int
	// 单向写通道
	//var numChanWriteOnly chan <- int

	// 1 在主函数中创建一个双向通道
	numChan := make(chan int, 5)

	// 2 将numChan传递给producer，负责生产
	go producer(numChan) //双向通道可以赋值给同类型的单向通道, 单向不能转双向

	// 3 将numChan传递给consumer，负责消费
	go consumer(numChan)

	time.Sleep(2 * time.Second)
	fmt.Println("OVER!")
}

// 生产者，提供一个只写的通道
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		fmt.Println("向通道写入数据 write ：", i)
	}
	close(out)
}

// 消费者，提供一个只读的通道
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("从通道读取数据：", v)
	}
	fmt.Println("consumer end")
}
