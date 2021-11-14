package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 1 创建监听
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)

	// 2 accept 接受连接
	// net.Listen("tcp", "8848) 简写，冒号前面默认是本机:127.0.0.1
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net.Listen error:", err)
	}

	fmt.Println("监听中...")

	// server可以监听多个连接，可以处理多个数据请求
	// 主程序负责监听，子goroutine负责数据处理

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
			return
		}

		fmt.Println("连接建立成功!")
		go handleData(conn)
	}

}

// 处理数据
func handleData(conn net.Conn) {
	// for 循环，保证每一个连接可以多次接收处理客户端请求
	for {
		// 3 read 读取客户端发来的数据
		// 创建一个容器，用于接收读取到的数据
		buffer := make([]byte, 1024)

		// Read(b []byte) (n int, err error)
		// cnt:真正读取client发来的数据的长度
		cnt, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("conn.Read error:", err)
			return
		}

		fmt.Println("Client 传输到 Server,长度：", cnt, ", 数据：", string(buffer[0:cnt]))

		// 服务端对客户端请求进行响应，将数据转成大写
		// func ToUpper(s string) string
		upperData := strings.ToUpper(string(buffer[0:cnt]))

		// 4 write 向客户端发送数据
		cnt, err = conn.Write([]byte(upperData))
		if err != nil {
			fmt.Println("conn.Write error:", err)
			return
		}
		fmt.Println("Client 获取到 Server,长度：", cnt, ", 数据：", upperData)
	}

	// 5 close 关闭连接
	_ = conn.Close()
}
