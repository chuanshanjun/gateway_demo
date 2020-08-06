package main

import (
	"log"
	"net/http"
	"time"
)

var Addr = ":9090"

func main() {
	// 创建路由器
	mux := http.NewServeMux()
	// 创建handler
	mux.HandleFunc("/hello", hello)
	// 创建Server
	server := &http.Server{
		Addr:         Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	// 监听端口并提供服务
	log.Println("Starting httpserver at " + Addr)
	log.Fatal(server.ListenAndServe())
}

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	w.Write([]byte("This is Server"))
}
