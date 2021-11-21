package main

import (
	"fmt"
	"sync"
	"time"
)

var idNames = make(map[int]string)
var lock sync.RWMutex

//map不允许同时读写，如果有不同go程同时操作map，需要对map上锁
// concurrent map read and map write
func main() {
	go func() {
		for {
			fmt.Println("111111")
			lock.Lock()
			fmt.Println("222222")
			idNames[0] = "peter"
			fmt.Println("333333")
			lock.Unlock()
		}
	}()

	go func() {
		for {
			fmt.Println("aaaaaa")
			lock.Lock()
			fmt.Println("bbbbbb")
			name := idNames[0]
			fmt.Println("name :", name)
			lock.Unlock()
		}
	}()

	for {
		fmt.Println("OVER")
		time.Sleep(1 * time.Second)
	}
}
