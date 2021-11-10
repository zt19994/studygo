package main

import "fmt"

// 实现go多态，必须实现定义接口

// IAttack 定义一个接口，注意类型interface，模拟不同等级的伤害
type IAttack interface {
	// Attack 接口函数可以有多个，但是只能有原型，不能有实现
	Attack()
}

// LowLevel 低等级
type LowLevel struct {
	name  string
	level int
}

func (a *LowLevel) Attack() {
	fmt.Println("我是：", a.name, ", 等级为：", a.level, ", 造成100点伤害")
}

// HighLevel 高等级
type HighLevel struct {
	name  string
	level int
}

func (a *HighLevel) Attack() {
	fmt.Println("我是：", a.name, ", 等级为：", a.level, ", 造成10000点伤害")
}

// DoAttack 定义一个多态的通用接口,传入不同的对象，调用同样的方法，实现不同的效果 ==》 多态
func DoAttack(a IAttack) {
	a.Attack()
}

// 多态 go语言不需要继承，只要实现相同的接口即可
func main() {
	var player IAttack // 定义一个包含Attack的接口变量

	lowLevel := LowLevel{
		name:  "peter",
		level: 1,
	}

	highLevel := HighLevel{
		name:  "David",
		level: 10,
	}

	lowLevel.Attack()
	highLevel.Attack()

	// 对player赋值lowLevel
	player = &lowLevel
	player.Attack()

	player = &highLevel
	player.Attack()

	fmt.Println("多态")
	DoAttack(&lowLevel)
	DoAttack(&highLevel)
}
