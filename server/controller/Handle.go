package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"websocketDemo2/server/component"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

//注册控制器
func RegisterHandle(res http.ResponseWriter, req *http.Request) {
	//获取请求参数
	vars := req.URL.Query()
	var channelNum int
	var err error
	if _, ok := vars["channel"]; ok {
		channelNum, err = strconv.Atoi(vars["channel"][0])
	} else {
		channelNum = 0
	}
	contracts := vars["contract"]

	//解析一个连接
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(res, req, nil)
	if err != nil {
		io.WriteString(res, "failed!")
		return
	}
	uid, _ := uuid.NewV4()
	sha1 := uid.String()
	channel := component.GetChannelManager().GetChannel(channelNum)
	//初始化client
	client := &component.Client{
		ID:       sha1,
		Socket:   conn,
		Contract: make(map[string]bool),
		Channel:  channel,
	}
	for _, contract := range contracts {
		client.Contract[contract] = true
	}
	//注册订阅
	channel.Register(client)
	//设置读取pong超时处理机制
	go client.Read()
}

//推送信息控制器
func PushHandle(res http.ResponseWriter, req *http.Request) {
	fileName := "demo.txt"
	msg, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(msg))
	component.GetChannelManager().GetChannel(0).Push(msg)
	component.GetChannelManager().GetChannel(1).Push(msg)
}
