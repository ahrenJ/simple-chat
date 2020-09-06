package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WsPool struct {
	wsCreateChan chan *Client
	wsCloseChan  chan *Client
	clients      map[int]*Client
}

var upgrader = websocket.Upgrader{
	WriteBufferSize: 512,
	ReadBufferSize:  512,
}

func NewWsPool() *WsPool {
	return &WsPool{
		wsCreateChan: make(chan *Client),
		wsCloseChan:  make(chan *Client),
		clients:      make(map[int]*Client),
	}
}

func HandleWs(wsPool *WsPool, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket连接创建失败: %v", err)
		panic(err)
	}

	var client = Client{conn: conn}
	wsPool.wsCreateChan <- &client
	go client.listen()
	log.Printf("websocket connection created successfully!")

	var msg = "welcome to simple-chat"
	conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (wsPool *WsPool) Start() {
	for {
		select {
		case client := <-wsPool.wsCreateChan:
			client.id = len(wsPool.clients)
			wsPool.clients[client.id] = client
		case client := <-wsPool.wsCloseChan:
			delete(wsPool.clients, client.id)
		}
	}
}

func (wsPool *WsPool) CountClients() {
	log.Printf("当前websocket连接数: %d", len(wsPool.clients))
}
