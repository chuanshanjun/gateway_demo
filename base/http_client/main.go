package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	// 创建连接池
	transport := &http.Transport{
		// TCP层的连接参数
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接
		IdleConnTimeout:       90 * time.Second, // 空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, // TLS握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  // 100-continue状态码超时时间
	}
	// 创建客户端
	client := &http.Client{
		Timeout:   time.Second * 30, // 请求超时时间-这个是控制总的,上面的超时时间是每个连接的
		Transport: transport,
	}
	// 请求数据
	resp, err := client.Get("http://127.0.0.1:9090/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 读取内容
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
