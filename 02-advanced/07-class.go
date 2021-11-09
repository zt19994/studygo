package main

import "fmt"

type Person struct {
	// 成员属性：
	name   string
	age    int
	gender string
	score  float64
}

// Eat 在类外面绑定方法 (p Person)Eat()
func (p Person) Eat() {
	fmt.Println("Person is eating")
	// 类的方法，可以使用自己的成员
	fmt.Println(p.name + " is eating")
}

// Sleep 使用指针绑定方法
func (p *Person) Sleep() {
	fmt.Println("Person is sleeping")
	fmt.Println(p.name + " is sleeping")
}

// Sleep2 使用指针可以修改Person的值
func (p *Person) Sleep2() {
	fmt.Println("Person is sleeping")
	p.name = "mike"
	fmt.Println(p.name + " is sleeping")
}

// Sleep3 不是指针，不可以修改Person的值
func (p Person) Sleep3() {
	fmt.Println("Person is sleeping")
	p.name = "mike"
	fmt.Println(p.name + " is sleeping")
}

func main() {
	peter := Person{
		name:   "peter",
		age:    20,
		gender: "man",
		score:  10,
	}

	fmt.Println("修改前peter:", peter)
	peter.Eat()
	peter.Sleep()
	peter.Sleep3()
	fmt.Println("不是指针修改后peter:", peter)
	peter.Sleep2()
	fmt.Println("是指针修改后peter:", peter)
}
