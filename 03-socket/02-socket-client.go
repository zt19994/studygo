package main

import (
	"fmt"
	"net"
)

func main() {
	// 1 指定server Dial 建立连接
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("net.Dial error:", err)
	}

	fmt.Println("Client Server 建立连接成功")
	sendData := []byte("hello world")

	// 2 write 向服务器发送数据
	cnt, err := conn.Write(sendData)
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return
	}

	fmt.Println("Client 发送数据到 Server cnt:", cnt, ", 数据：", sendData)

	// 3 接收服务器返回的数据
	// 创建buffer，用于接收服务器返回的数据
	buffer := make([]byte, 1024)
	cnt, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("conn.Read error:", err)
		return
	}

	fmt.Println("Client 接收到 Server, cnt:", cnt, ", 数据：", string(buffer[0:cnt]))

	// 4 close 关闭连接
	err = conn.Close()
	if err != nil {
		fmt.Println("conn.Close error:", err)
		return
	}

}
