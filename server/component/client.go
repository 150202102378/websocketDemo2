package component

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 5 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (9 * pongWait) / 10

	// Poll file for changes with this period.
	filePeriod = 10 * time.Second
)

//Client 用户
type Client struct {
	ID       string
	Socket   *websocket.Conn
	Contract map[string]bool
	Channel  *Channel
}

//客户端关闭后会读到messageType = -1并退出
func (c *Client) Read() {
	//结束时关闭连接
	defer func() {
		c.Socket.Close()
		c.Channel.unregister <- c
	}()
	//设置读取的大小
	//c.Socket.SetReadLimit(512)

	//设置读取pong超时处理机制
	c.Socket.SetReadDeadline(time.Now().Add(pongWait))
	c.Socket.SetPongHandler(func(string) error {
		//fmt.Println("receive pong")
		c.Socket.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	//设置读超时机制
	//c.Socket.SetReadDeadline(time.Now().Add(pongWait))

	for {
		frameType, message, err := c.Socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Println("error: ", err)
			}
			break
		}
		if frameType == 1 {
			fmt.Println("service recv:", string(message))
		}
	}
}
