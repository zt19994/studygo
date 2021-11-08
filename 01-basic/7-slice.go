package main

import "fmt"

func main() {
	//切片：slice，它的底层也是数组，可以动态改变长度
	//定义一个切片，包含多个地名
	//names := [10]string{"北京", "上海", "广州", "深圳"}
	names := []string{"北京", "上海", "广州", "深圳"}
	for i, v := range names {
		fmt.Println("i:", i, "v:", v)
	}

	//1 追加数据
	fmt.Println("追加元素前,name的长度:", len(names), "，容量:", cap(names))
	names = append(names, "chengdu")
	//fmt.Println("names追加元素后赋值给自己:", names)
	fmt.Println("追加元素后,name的长度:", len(names), "，容量:", cap(names))
	names = append(names, "西藏")
	fmt.Println("追加元素后,name的长度:", len(names), "，容量:", cap(names))

	//2.对于一个切片，不仅有'长度'的概念len()，还有一个'容量'的概念cap()
	var nums []int
	for i := 0; i < 50; i++ {
		nums = append(nums, i)
		fmt.Println("len:", len(nums), ", cap:", cap(nums))
	}
}
