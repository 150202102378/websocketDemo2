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
	host    string
	addr    *string
	dialer  *websocket.Dialer
	record  chan *dao.Record
	connNum int
)

func init() {
	config := conf.GetConfig()
	connNum = config.ConnNum
	host = config.ServerIP + ":" + config.ReactPort1
	//host = config.ServerIP + ":8081"
	addr = flag.String("addr1", host, "https service address")
	record = make(chan *dao.Record)
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

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("Dial error", err)
		go connect(i)
		return
	}
	defer func() {
		go connect(i)
		conn.WriteMessage(websocket.CloseMessage, []byte{})
		conn.Close()
	}()
	//从服务器读取信息
	//conn.SetPingHandler(nil)
	//conn.SetReadDeadline(time.Now().Add(13 * time.Second))
	conn.SetPingHandler(func(string) error {
		conn.WriteMessage(websocket.PongMessage, []byte{})
		//conn.SetReadDeadline(time.Now().Add(13 * time.Second))
		return nil
	})

	for {
		re := dao.Record{}
		//调用ReadMessage(),一旦服务器主动断开连接则会接收到 err = websocket: close 1006 (abnormal closure): unexpected EOF

		_, message, err := conn.ReadMessage()
		if err != nil {
			/* re.Message = string(message)
			re.Flag = 0
			record <- &re
			fmt.Println("read:", err, frameType) */
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Println("error: ", err)
			}
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
	for i := 0; i < connNum; i++ {
		time.Sleep(5 * time.Millisecond)
		go connect(i)
	}
	wg.Wait()
}
