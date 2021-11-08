package main

import "fmt"

func main() {
	// 1 定义
	name := "peter"

	// 需要换行，原生输出字符串时，使用反引号``
	usage := `1 2 3 4 5
			a b c d e
			`
	// c语言使用单引号 + \来解决
	fmt.Println("name:", name)
	fmt.Println("usage:", usage)

	// 2 长度，访问
	// C++:   name.length
	// GO: string没有.length方法，可以使用自由函数len()进行处理
	// len: 很常用

	l1 := len(name)
	fmt.Println("l1:", l1)

	// 不需要加()
	for i := 0; i < len(name); i++ {
		fmt.Printf("i:%d, v:%c\n", i, name[i])
	}

	// 3 拼接
	i, j := "hello", "world"
	fmt.Println("i + j = ", i+j)

	// 使用const修饰为常量，不能修改
	const address = "chengdu"
	// address = "shanghai"
	fmt.Println("address:", address)

	// 可以直接对比字符串，使用ASCII的值进行比较
	if i < j {
		fmt.Println(i)
	}

}
