package component

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

//Channel 订阅频道
type Channel struct {
	clients      map[*Client]bool
	broadcastMsg chan []byte
	register     chan *Client
	unregister   chan *Client
	lock         *sync.RWMutex
}

//Push 推送信息
func (channel *Channel) Push(msg []byte) {
	channel.broadcastMsg <- msg
}

//Register 注册
func (channel *Channel) Register(client *Client) {
	channel.register <- client
}

//Unregister 取消注册
func (channel *Channel) Unregister(client *Client) {
	channel.unregister <- client
}

//Start 运行频道
func (channel *Channel) Start() {
	//设置ping定时器
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
	}()
	for {
		select {
		case conn := <-channel.register: //订阅
			//锁操作
			channel.lock.Lock()
			channel.clients[conn] = true
			//conn.Socket.SetPongHandler(func(string) error { return nil })
			fmt.Println(len(channel.clients))
			//解锁
			channel.lock.Unlock()

		case conn := <-channel.unregister: //取消订阅
			if ok := channel.clients[conn]; ok {
				//锁操作
				channel.lock.Lock()
				delete(channel.clients, conn)
				//释放client空间
				conn.Socket.Close()
				conn = nil
				fmt.Println(len(channel.clients))
				//解锁
				channel.lock.Unlock()
			}

		case message := <-channel.broadcastMsg: //广播
			for conn := range channel.clients {
				if err := conn.Socket.WriteMessage(websocket.TextMessage, message); err != nil {
					fmt.Println(err)
					//需要先发送关闭信息再关闭conn
					conn.Socket.WriteMessage(websocket.CloseMessage, []byte{})
					//同一线程下对通道操作会阻塞，需要开goroutine来取消注册
					go func() { channel.unregister <- conn }()
				}
			}

		case <-pingTicker.C:
			fmt.Println("send ping!")
			//发送 ping
			for conn := range channel.clients {
				//conn.Socket.SetWriteDeadline(time.Now().Add(writeWait))
				if err := conn.Socket.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					//需要先发送关闭信息再关闭conn
					conn.Socket.WriteMessage(websocket.CloseMessage, []byte{})
					//同一线程下对通道操作会阻塞，需要开goroutine来取消注册
					go func() { channel.unregister <- conn }()
				}
			}
		}

	}
}
