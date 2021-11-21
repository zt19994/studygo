package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type User struct {
	// 名字
	name string
	// id
	id string
	// 消息
	msg chan string
}

// 创建一个全局map，用于保存所有用户
var allUsers = make(map[string]User)

// 定义一个message的全局通道，用于接收任何人发送过来的消息
var message = make(chan string, 10)

func main() {
	// 创建服务器
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}

	// 启动全局唯一的goroutine，负责监听message管道，写给所有用户
	go broadcast()

	fmt.Println("服务器启动成功!")

	for {
		fmt.Println("主goroutine监听中")

		// 监听
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listener.Accept error:", err)
			return
		}

		fmt.Println("建立连接成功！")

		// 启动处理业务的goroutine
		go handler(conn)
	}
}

// 向所有用户广播消息，启动一个全局唯一的goroutine
func broadcast() {
	fmt.Println("广播goroutine启动成功")
	defer fmt.Println("broadcast 程序退出")

	for {
		// 1 从message中读取数据
		info := <-message

		// 2 将数据写入每一个用户的msg管道中
		for _, user := range allUsers {
			// 如果msg是非缓冲的，那么会在这里阻塞
			user.msg <- info
		}
	}
}

func handler(conn net.Conn) {
	for {
		fmt.Println("启动业务...")
		// 客户端与服务器建立连接，将ip和port当做user的id
		clientAddr := conn.RemoteAddr().String()
		fmt.Println("clientAddr:", clientAddr)
		// 创建user
		newUser := User{
			id:   clientAddr,
			name: clientAddr,
			msg:  make(chan string, 10),
		}
		// 添加user到map结构
		allUsers[newUser.id] = newUser

		// 定义一个退出信号，用于监听client退出
		var isQuit = make(chan bool)

		// 定义一个用于重置计数器的通道，用于告知watch函数，当前用户正在输入
		var restTimer = make(chan bool)

		// 启动goroutine，负责监听退出信号
		go watch(&newUser, conn, isQuit, restTimer)

		// 启动goroutine程，负责将msg信息返回给客户端
		go writeBackToClient(&newUser, conn)

		// 向message写入数据，当前用户上线的消息，用于通知所有人（广播）
		loginInfo := fmt.Sprintf("[%s]:[%s] ----> 上线了 login！", newUser.id, newUser.name)
		message <- loginInfo

		buf := make([]byte, 1024)

		// 读取客户端发送过来的请求数据
		read, err := conn.Read(buf)
		if read == 0 {
			fmt.Println("客户端主动关闭，准备退出！")
			isQuit <- true
		}

		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("服务器接收客户端发送过来的数据为：", string(buf[:read-1]), ",read:", read)

		// 业务逻辑处理
		userInput := string(buf[:read-1])
		if len(userInput) == 3 && userInput == "who" {
			fmt.Println("用户即将查询所以用户信息")

			// 这个切片包含所有的用户信息
			var userInfos []string

			for _, user := range allUsers {
				userInfo := fmt.Sprintf("userId:%s ,username:%s", user.id, user.name)
				userInfos = append(userInfos, userInfo)
			}

			// 最终写入管道中，一定是一个字符串
			r := strings.Join(userInfos, "\n")

			newUser.msg <- r

		} else if len(userInput) == 7 && userInput[:7] == "\\rename" {
			newUser.name = strings.Split(userInput, "|")[1]
			// 更新map中的user
			allUsers[newUser.id] = newUser
			// 通知客户端，更新成功
			newUser.msg <- "rename successfully"
		} else {
			// 若不是命令who，那么只需要写到广播通道中即可
			message <- userInput
		}

		restTimer <- true
	}
}

// 每个用户都有一个自己的watch goroutine，负责监听退出信号
func watch(user *User, conn net.Conn, isQuit <-chan bool, restTimer <-chan bool) {
	fmt.Println("启动监听退出信号的goroutine")
	defer fmt.Println("watch goroutine退出！")

	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s exit already! \n", user.name)
			fmt.Println("删除当前用户：", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo

			_ = conn.Close()
			return
		case <-time.After(5 * time.Second):
			logoutInfo := fmt.Sprintf("%s timeout exit already! \n", user.name)
			fmt.Println("删除当前用户：", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo
			_ = conn.Close()
			return
		case <-restTimer:
			fmt.Printf("连接%s 重置计数器！\n", user.name)
		}
	}
}

// 用户应该有一个用来监听自己msg管道的goroutine，负责将数据返回给客户端
func writeBackToClient(user *User, conn net.Conn) {
	fmt.Printf("user : %s 的goroutine正在监听自己的msg管道：\n", user.name)
	for data := range user.msg {
		fmt.Printf("user : %s 回写给客户端的数据为：%s \n", user.name, data)

		_, _ = conn.Write([]byte(data))
	}
}
