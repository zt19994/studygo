package main

import "fmt"

type MyInt int

// Student go语言结构体使用type + struct来处理
type Student struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {
	t1 := struct {
		a int
	}{
		a: 100,
	}

	fmt.Println(t1)

	var i, j MyInt
	i, j = 10, 20
	fmt.Println("i+j:", i+j)

	// 创建变量并赋值
	peter := Student{
		name:   "peter",
		age:    25,
		gender: "man",
		score:  99,
	}

	// 使用结构体
	fmt.Println("peter:", peter.name, peter.age, peter.gender, peter.score)

	//结构体没有-> 操作
	s1 := &peter
	fmt.Println("lily 使用指针s1.name打印:", s1.name, s1.age, s1.gender, s1.score)
	fmt.Println("lily 使用指针(*s1).name打印:", (*s1).name, s1.age, s1.gender, s1.score)

	// 在定义期间对结构体赋值时，如果每个字段都赋值了，那么字段的名字可以省略不写
	// 如果只对局部变量赋值，那么必须明确指定变量名字
	Mike := Student{
		name:   "Mike",
		age:    22,
		gender: "man",
	}
	Mike.score = 100

	fmt.Println("Mike:", Mike)

}
