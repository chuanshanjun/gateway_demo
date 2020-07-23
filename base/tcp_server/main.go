package main

import (
	"fmt"
	"net"
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
		var data [256]byte
		n, err := conn.Read(data[:])
		if err != nil {
			fmt.Printf("read from client failed err:%v\n", err)
			break
		}
		fmt.Printf("received msg from: %v\n", string(data[:n]))
	}
}
