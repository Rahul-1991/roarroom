package handlers

import (
	"log"
	"net/http"
	client "roarroom/store"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allowing connections from any origin for demonstration purposes
		return true
	},
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	clientId := r.URL.Query().Get("clientId")
	log.Println("Client connnected: ", clientId)

	// Add connection to the connections map
	client.AddClientConnection(clientId, conn)
	defer client.RemoveClientConnection(clientId)

	for {

		resultMap := client.ReadFromClient(conn)
		if resultMap == nil {
			break
		}

		recipientId := resultMap["recepient"]

		recipientConn := client.GetClientConnection(recipientId)

		if recipientConn != nil {
			err := client.WriteToClient(recipientConn, resultMap["message"])
			if err {
				break
			}
		} else {
			log.Println("Could not find connection to client")
		}

		// Send a response or echo the received message back to the client

	}
}
