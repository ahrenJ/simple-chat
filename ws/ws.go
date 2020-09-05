package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 512,
	ReadBufferSize:  512,
}

var wsConnPool = make([]*websocket.Conn, 0)

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	wsConnPool = append(wsConnPool, conn)
	log.Printf("websocket connection created successfully!")

	var data = "welcome to simple-chat"
	conn.WriteMessage(websocket.TextMessage, []byte(data))
}

func SendDataTask() {
	for {
		time.Sleep(time.Second * 3)
		log.Println("SendDataTask begin...")
		for _, conn := range wsConnPool {
			var data = "hi"
			err := conn.WriteMessage(websocket.TextMessage, []byte(data))
			if err != nil {
				log.Printf("SendDataTask-send data fail: %v", err)
			}
		}
	}
}
