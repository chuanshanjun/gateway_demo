package main

import (
	"fmt"
	"net"
)

// 因为是udp服务器，那么他就不涉及到监听，然后创建socket连接了
// 他是直接去读这个数据的
func main() {
	// 1 监听服务器
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	defer listen.Close()
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}

	// 2 循环读取消息内容
	for {
		// 读取1024个字节的长度
		var data [1024]byte
		// 循环读消息的时候，不是用socket去读取的，而是直接使用listen读取的
		// ReadFromUDP从c(listen)读取一个UDP数据包，将有效负载拷贝到b(data[:])，返回拷贝 字节数(n) 和 数据包来源地址 。
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read failed from add:%v, err: %v\n", addr, err)
			break
		}

		// 使用协程去回复这个数据
		go func() {
			// 回复的时候加一些逻辑的处理流程
			// todo sth
			// 3 回复数据
			// 写入的地址，就直接使用上面read返回的 udpAddr
			fmt.Printf("addr: %v, data: %v, count: %v\n", addr, string(data[:n]), n)
			_, err = listen.WriteToUDP([]byte("received success!"), addr)
			if err != nil {
				fmt.Printf("write failed, err: %v\n", err)
			}
		}()
	}
}
