package main

import "fmt"

func init() {
	fmt.Println("this is first init")
}

//0.这个init会在包被引用的时候(import)进行自动调用
//1.init函数没有参数，没有返回值，原型固定如下
//2.一个包中包含多个init时，调用顺序是不确定的(同一个包的多个文件中都可以有init)
//3.init函数时不允许用户显示调用的
//4.有的时候引用一个包，可能只想使用这个包里面的init函数（mysql的init对驱动进行初始化）
//但是不想使用这个包里面的其他函数，为了防止编译器报错，可以使用_形式来处理
//import _ "xxx/xx/sub"
func main() {
	fmt.Println("this is main func")
}
