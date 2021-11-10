package main

import "fmt"

type Human struct {
	// 成员属性：
	name   string
	age    int
	gender string
}

// Eat 绑定方法
func (this *Human) Eat() {
	fmt.Println("this is :", this.name)
}

// Student1 定义一个学生类，去嵌套一个Hum
type Student1 struct {
	hum    Human // 包含Human类型的变量,此时类的嵌套
	score  float64
	school string
}

// Teacher 定义一个teacher类，实现一个Hum
type Teacher struct {
	Human          // 直接写Hum类型，没有字段名字
	subject string //学科
}

func main() {
	s1 := Student1{
		hum: Human{
			name:   "Peter",
			age:    18,
			gender: "man",
		},
		school: "一中",
	}

	fmt.Println("s1.name:", s1.hum.name)
	fmt.Println("s1.school:", s1.school)

	t1 := Teacher{}
	t1.subject = "数学" //下面几个字段都是继承自human
	t1.name = "赵老师"
	t1.gender = "man"

	fmt.Println("t1:", t1)
	t1.Eat()

	// 继承时，虽然没有定义字段名字，但是会自动创建一个默认的同名字段
	// 这是为了在子类中依然可以操作父类，因为：子类父类可能出现同名的字段
	fmt.Println("t1.Human.name:", t1.Human.name)
}
