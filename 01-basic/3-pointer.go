package main

import "fmt"

func main() {
	// go语言也有指针
	// 结构体成员调用时： c语言： ptr->name go语言：ptr.name
	// go语言在使用指针时，会使用内部的垃圾回收机制（gc:garbage collector）开发人员不需要手动释放内存
	// c语言不允许返回栈上指针，go语言可以栈上返回，程序会在编译的时候确定类这个变量的分配位置
	// 在编译的时候，如果发现有必要的话，就将变量分配到堆上

	// 1 使用&定义
	name := "peter"
	ptr := &name
	fmt.Println("name:", *ptr)
	fmt.Println("name ptr:", ptr)

	// 2 使用new关键字定义 * 赋值
	name2ptr := new(string)
	*name2ptr = "mike"
	fmt.Println("name2:", *name2ptr)
	fmt.Println("name2 ptr:", name2ptr)

	// 可以返回栈上的指针，编译器在编译程序时会自动判断这段代码，将city变量分配在堆上，内存逃逸
	res := testPtr()
	fmt.Println("res city:", *res)
	fmt.Println("res city ptr:", res)
}

// 定义一个函数，返回一个string类型的指针，go语言返回写在参数列表后面
func testPtr() *string {
	city := "chengdu"
	ptr := &city
	return ptr
}
