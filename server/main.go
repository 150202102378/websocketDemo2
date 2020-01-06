package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"websocketDemo2/conf"
	"websocketDemo2/server/component"
	"websocketDemo2/server/controller"
)

var (
	listenPort1 string
	listenPort2 string
)

func init() {
	listenPort1 = conf.GetConfig().ListenPort1
	listenPort2 = conf.GetConfig().ListenPort2
}

func main() {
	fmt.Println("Starting application...")
	runtime.GOMAXPROCS(2)
	cm := component.GetChannelManager()
	cm.Start()
	//可以改用gin

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	//推送和注册分开两个端口号监听，防止注册量太大时阻塞推送功能
	//port: 8081
	mux := http.NewServeMux()
	mux.HandleFunc("/register", controller.RegisterHandle)
	go http.ListenAndServe(":"+listenPort1, mux)

	//port: 8082
	mux = http.NewServeMux()
	mux.HandleFunc("/push", controller.PushHandle)
	http.ListenAndServe(":"+listenPort2, mux)
	//带tls
	//http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
}
