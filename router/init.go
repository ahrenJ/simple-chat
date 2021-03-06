package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-chat/ws"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("pages/*")
	e.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello simple-chat!")
	})
	e.GET("/wsPage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "websocket.html", nil)
	})
	//websocket
	var wsPool = ws.NewWsPool()
	go wsPool.Start()
	e.GET("/ws", func(c *gin.Context) {
		//建立websocket连接
		ws.HandleWs(wsPool, c.Writer, c.Request)
	})
	e.GET("/wsCount", func(c *gin.Context) {
		wsPool.CountClients()
	})
	e.Run(":8080")
}
