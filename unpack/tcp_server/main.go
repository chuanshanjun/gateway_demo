package main

import (
	"fmt"
	"gateway_demo/unpack/unpack"
	"net"
)

func main() {
	// simple tcp server
	// 1.listen ip + port
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}

	// 2.accept client request
	// 3.create goroutine for each request
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed, err: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from connect fail, err: %v\n", err)
			// 如果读取错误则跳出for循环,然后关闭conn
			break
		}

		str := string(bt)
		fmt.Println("receive from client, data: %v\n", str)
	}
}
