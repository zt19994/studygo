package main

import "fmt"

// 在go语言中，有专门的关键字 interface来代表接口
// interface不仅仅用于处理多态，它可以接受任意的数据类型，有点类似void
func main() {

	var i, j, k interface{}
	name := []string{"peter", "mike"}
	i = name
	fmt.Println("i:", i)

	age := 20
	j = age
	fmt.Println("j:", j)

	gender := "man"
	k = gender
	fmt.Println("k:", k)

	// 只知道k是interface，但是不能够明确知道它代表的数据的类型
	kvalue, ok := k.(int)
	if !ok {
		fmt.Println("k不是int")
	} else {
		fmt.Println("k是int，值为：", kvalue)
	}

	// 最常用的场景，把interface当成一个函数的参数，使用switch来判断用户输入的不同类型
	// 根据不同的类型，做不同的逻辑处理
	// 创建一个具有三个接口类型的切片
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "Hello world"
	//array[2] = true
	array[2] = 3.14

	for _, value := range array {
		switch v := value.(type) {
		case int:
			fmt.Printf("当前类型为int：内容为：%d\n", v)
		case string:
			fmt.Printf("当前类型为string：内容为：%s\n", v)
		case bool:
			fmt.Printf("当前类型为bool：内容为：%v\n", v)
		default:
			fmt.Println("不是合理的数据类型")
		}
	}

}
