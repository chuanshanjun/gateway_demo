package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 1 监听获取socket连接
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	defer listen.Close()
	if err != nil {
		fmt.Printf("listen failed: %v\n", err)
		return
	}

	// 2 建立套接字
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("build connect failed err: %v\n", err)
			// 注意它建立一个socket失败后就continue，因为它是server 不要用break,return
			continue
		}
		// 3 创建处理协程
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var data [1024]byte
		_, err := conn.Read(data[:])
		if err != nil {
			fmt.Printf("read from client failed err:%v\n", err)
			break
		}
		msg := string(data[:])
		splitMsg := strings.Split(msg, "\r\n")
		requestRAW := strings.Split(splitMsg[0], " ")
		fmt.Printf("Request Method is: %s\n", requestRAW[0])
		fmt.Printf("Resuest Param is: %s\n", requestRAW[1])
		fmt.Printf("HTTP Version is: %s\n", requestRAW[2])
	}
}
