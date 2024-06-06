package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connection = make(map[string]*websocket.Conn)

type Server struct {
	router *gin.Engine
	port   string
}

func NewServer(port string) *Server {
	router := gin.Default()
	return &Server{
		router: router,
		port:   port,
	}
}

func (s *Server) Run() error {
	return s.router.Run(fmt.Sprint(":", s.port))
}

func (s *Server) Websocket() {
	s.router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close()
		uuid := uuid.New().String()
		connection[uuid] = conn
		fmt.Printf("Connection: %d\n", len(connection))
		for {
			_, msg, err := conn.ReadMessage()
			fmt.Printf("Message: %s\n", msg)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			conn.WriteMessage(websocket.TextMessage, msg)
		}
	})
}
