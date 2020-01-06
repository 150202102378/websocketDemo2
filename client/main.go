package main

import (
	"flag"
	"fmt"
	"net/url"
	"runtime"
	"sync"
	"time"
	"websocketDemo2/conf"
	"websocketDemo2/dao"

	"github.com/gorilla/websocket"
)

var (
	host   string
	addr   *string
	dialer *websocket.Dialer
	record chan *dao.Record
)

func init() {
	config := conf.GetConfig()
	//host = config.ServerIP + ":" + config.ReactPort1
	host = config.ServerIP + ":8081"
	addr = flag.String("addr1", host, "https service address")
	record = make(chan *dao.Record)
	dialer = &websocket.Dialer{HandshakeTimeout: 10000 * time.Hour}
}

func RecordMessage() {
	for {
		select {
		case rec := <-record:
			dao.RecordMessage(rec)
		}
	}
}

//链接
func connect(i int) {
	u := url.URL{
		Scheme:   "ws",
		Path:     "/register",
		RawQuery: "channel=0",
	}
	u.Host = *addr

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		go connect(i)
		return
	}
	defer func() {
		conn.WriteMessage(websocket.CloseMessage, []byte{})
		conn.Close()
	}()
	//从服务器读取信息

	for {
		re := dao.Record{}
		//调用ReadMessage(),一旦服务器主动断开连接则会接收到 err = websocket: close 1006 (abnormal closure): unexpected EOF
		_, message, err := conn.ReadMessage()
		if err != nil {
			/* re.Message = string(message)
			re.Flag = 0
			record <- &re
			fmt.Println("read:", err, frameType) */
			return
		}
		re.Message = string(message)
		re.Flag = 1
		record <- &re
		//fmt.Printf("received: %s\n", string(message))
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(1)
	go RecordMessage()
	for i := 0; i < 30000; i++ {
		time.Sleep(10 * time.Millisecond)
		go connect(i)
	}
	wg.Wait()
}
