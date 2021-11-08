package main

import "fmt"

func main() {
	// 定义定义：var
	// 常量定义：const

	// 1 先定义变量，再赋值 var 变量名 数据类型
	var name string
	name = "peter"

	var age int
	age = 20

	fmt.Println(name)
	// fmt.Printf 格式化打印
	fmt.Printf("name is :%s, %d\n", name, age)

	// 2 定义时直接赋值
	var gander = "man"
	fmt.Println("gander:", gander)

	// 3 定义时直接赋值，使用自动推导 (最常用)
	address := "chengdu"
	fmt.Println("address:", address)

	// 4 平行赋值
	i, j := 10, 20 // 同时定义两个变量
	fmt.Printf("变换前：i:%d , j:%d", i, j)
	fmt.Println()
	i, j = j, i
	fmt.Printf("变换后：i:%d , j:%d", i, j)
}
