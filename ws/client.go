package ws

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	id   int
	conn *websocket.Conn
}

//监听客户端发送的消息
func (c *Client) listen() {
	defer c.conn.Close()
	for {
		messageType, bytes, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("读取客户端数据错误: %v\n", err)
			break
		}
		log.Printf("接收客户端数据(type=%d): %s", messageType, string(bytes))
	}
}

//发送消息到客户端
func (c *Client) write() {

}
