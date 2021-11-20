package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int
	Name string
	Age  int
	// gender小写
	gender string
}

func main() {
	//在网络中传输的时候，把Student结构体，编码成json字符串，传输  ===》 结构体 ==》 字符串  ==》 编码
	//接收字符串，需要将字符串转换成结构体，然后操作 ==》 字符串 ==》 结构体  ==》解密

	peter := Student{
		Id:     1,
		Name:   "peter",
		Age:    20,
		gender: "man",
	}

	// 由于gender在结构体中是小写字母开头的，json编码不会参与编码
	marshal, err := json.Marshal(&peter)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	fmt.Println("json.Marshal :", string(marshal))

	var peter2 Student
	if err := json.Unmarshal([]byte(marshal), &peter2); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}

	fmt.Println("id:", peter2.Id)
	fmt.Println("name:", peter2.Name)
	fmt.Println("age:", peter2.Age)
	fmt.Println("gender:", peter2.gender)
}
