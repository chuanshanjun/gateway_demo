package main

import (
	"fmt"
	"net"
)

func main() {
	// 1 连接udp服务器
	// 会拿到一个socket连接?
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	defer conn.Close()
	if err != nil {
		fmt.Printf("connet server failed err: %v\n", err)
		return
	}

	// 以下为何要把write及read放到一块呢？
	// 因为此socket是不可靠的连接，所以发送的结果不一定是我读取的结果
	// 为了保证发送和读取是一致的，可以用下面比较简单的方法(发一条，收一条，但发送完了也不一定能接收的到)
	//
	for i := 0; i < 100; i++ {
		// 2 发送数据
		// 直接在socket连接上write?
		if _, err := conn.Write([]byte("hello server!")); err != nil {
			fmt.Printf("send data failed, err: %v\n", err)
			return
		}

		// 3 接收数据
		// 直接在socket上read
		result := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(result)
		if err != nil {
			fmt.Printf("received data failed, err: %v\n", err)
			return
		}
		fmt.Printf("receive from addr: %v data: %v\n", remoteAddr, string(result[:n]))
	}
}
