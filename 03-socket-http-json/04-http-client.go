package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// http包
	client := http.Client{}

	response, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("CLIENT.Get err:", err)
		return
	}

	ct := response.Header.Get("Content-Type")
	date := response.Header.Get("Date")
	server := response.Header.Get("Server")

	fmt.Println("Content-Type:", ct)
	fmt.Println("Date:", date)
	fmt.Println("Server:", server)

	url := response.Request.URL
	code := response.StatusCode
	status := response.Status

	fmt.Println("url:", url)
	fmt.Println("code:", code)
	fmt.Println("status:", status)

	body := response.Body
	fmt.Println("body :", body)
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}

	fmt.Println("body string:", string(readBodyStr))
}
