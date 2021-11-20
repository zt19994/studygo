package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 注册路由 router

	// func是回调函数，用于路由的响应，这个回调函数原型是固定
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request : ===> 包含客户端发来的数据
		fmt.Println("用户请求详情:")
		fmt.Println("request:", request)

		//这里是具体处理业务逻辑xxx

		//writer :  ===> 通过writer将数据返回给客户端
		_, _ = io.WriteString(writer, "这是/user请求返回的数据!")
	})

	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/name请求返回的数据!")
	})

	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/id请求返回的数据!")
	})

	fmt.Println("Http Server start ...")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("http start failed, err:", err)
		return
	}
}
