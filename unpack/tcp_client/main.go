package main

import (
	"fmt"
	"gateway_demo/unpack/unpack"
	"net"
)

func main() {
	// 连接服务器拿到套接字
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	// 通过socket执行encode方法,conn中默认实现了write及read方法
	unpack.Encode(conn, "hello world 0 !!!")
}
