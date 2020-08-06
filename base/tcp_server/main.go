package main

import (
	"fmt"
	"gateway_demo/base/tcp_server/static"
	"io"
	"net"
	"net/http"
	"strings"
)

// 使用map存储URI与Handler映射关系
var handlerMap map[string]HandlerFunc

// 定义函数类型HandlerFunc，其参数为net.Conn，无返回值
type HandlerFunc func(w io.Writer)

// 为函数类型HandlerFunc定义其方法DoRequest
func (f HandlerFunc) DoRequest(w io.Writer) {
	// 2.7返回信息给客户
	f(w)
}

// 响应头
var respHeader = `HTTP/1.1 200 OK
        Date: Sat, 29 Jul 2017 06:18:23 GMT
        Content-Type: text/html
        Connection: Keep-alive
        Server: Golang

`

// 自定义函数HelloWorldHandler
func HelloWorldHandler(w io.Writer) {
	w.Write([]byte(respHeader))
	// 输出html模版
	static.ReadHtml(w)
}

// 自定义函数good
func Good(w io.Writer) {
	w.Write([]byte(respHeader))
	w.Write([]byte("Good Job"))
}

// 错误检查函数
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Init() {
	handlerMap = make(map[string]HandlerFunc)
	hf := HandlerFunc(HelloWorldHandler)
	handlerMap["/helloWorldHandler"] = hf

	good := HandlerFunc(Good)
	handlerMap["/good"] = good
}

func main() {
	// 1.1 绑定函数与URI
	Init()
	// 1.2 监听获取Socket
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	defer listen.Close()
	if err != nil {
		fmt.Printf("listen failed: %v\n", err)
		return
	}

	// 2.1 接收请求，产生连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("build connect failed err: %v\n", err)
			// 注意它建立一个socket失败后就continue，因为它是server 不要用break,return
			continue
		}
		// 2.2 使用协程处理新连接的业务逻辑
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	// 2.3 读取文本
	var data [1024]byte
	_, err := conn.Read(data[:])
	if err != nil {
		fmt.Printf("read from client failed err:%v\n", err)
		return
	}

	msg := string(data[:])
	splitMsg := strings.Split(msg, "\r\n")
	requestRAW := strings.Split(splitMsg[0], " ")
	DisRequest(requestRAW, conn)
}

func DisRequest(r []string, c net.Conn) {
	requestType := r[0]
	// 2.4 匹配HTTP请求类型
	switch requestType {
	case http.MethodGet:
		fmt.Printf("Request Method is: %s\n", requestType)
		getProcess(r, c)
	case http.MethodPost:
		// TODO
	case http.MethodPut:
		// TODO
	case http.MethodDelete:
		// TODO
	}
}

func getProcess(r []string, c net.Conn) {
	uri := r[1]
	if len(uri) <= 1 {
		c.Write([]byte("Has No URI"))
		return
	}
	// 2.5 匹配URI对应的Handler方法
	handler := handlerMap[uri]

	if handler == nil {
		c.Write([]byte("404 NOT FOUND"))
		return
	}

	fmt.Printf("Request URI is :%s\n", uri)
	// 2.6 执行Handler方法
	handler.DoRequest(c)
}
