package main

import (
	"log"
	"net/http"
	handlers "roarroom/handlers"
	server "roarroom/server"
)

func main() {
	http.HandleFunc("/ws", handlers.WebSocketHandler)
	log.Println("Server started on http://localhost:8080")
	server.CreateHttpConnection()
}
