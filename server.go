package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Server struct {
	Router           *gin.Engine
	SocketConnection websocket.Upgrader
}

// Here we create our new server
func NewServer() *Server {
	router := gin.Default()
	// here we opened cors for all
	router.Use(cors.Default())
	return &Server{
		Router: router,
	}
}
