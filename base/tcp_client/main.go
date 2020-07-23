package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1 连接服务器
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("build connect failed %v\n", err)
		return
	}

	// 2 读取命令行
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 3 一直读取直到读到\n
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed err: ", err)
			break
		}
		// 4 读到Q时，退出
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		// 5 写数据到服务器
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write msg to server err: %v\n", err)
			break
		}
	}
}
