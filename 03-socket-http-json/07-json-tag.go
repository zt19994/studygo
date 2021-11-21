package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"-"`                 //在使用json编码时，这个不参与编码
	Subject string `json:"subject_name"`      //在json编码时，定义别名
	Age     int    `json:"age,string"`        //在json编码时，将age转换为string类型，格式：名字,类型。中间不能有空格
	Address string `json:"address,omitempty"` //在json编码时，如果这个字段为空，那么忽略掉，不参与编码
	// gender小写
	gender string
}

type Master struct {
	Name    string
	Subject string
	Age     int
	Address string
	gender  string
}

func main() {

	t1 := Teacher{
		Name:    "Peter",
		Subject: "Math",
		Age:     22,
		gender:  "man",
	}

	fmt.Println("t1:", t1)
	marshal, _ := json.Marshal(&t1)

	fmt.Println("marshal:", string(marshal))

	//解码
	t2 := Teacher{}
	_ = json.Unmarshal(marshal, &t2)
	fmt.Println("t2:", t2)
	fmt.Println("t2:", t2.Subject)

	m1 := Master{}
	_ = json.Unmarshal(marshal, &m1)
	fmt.Println("m1:", m1)
}
