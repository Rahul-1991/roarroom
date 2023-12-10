package store

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Store active WebSocket connections
var connections = make(map[string]*websocket.Conn)

func AddClientConnection(clientId string, conn *websocket.Conn) {
	connections[clientId] = conn
}

func RemoveClientConnection(clientId string) {
	delete(connections, clientId)
}

func GetClientConnection(clientId string) *websocket.Conn {
	if conn, ok := connections[clientId]; ok {
		return conn
	} else {
		return nil
	}
}

func ReadFromClient(conn *websocket.Conn) map[string]string {
	// Read message from the WebSocket connection
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error reading message:", err)
		return nil
	}
	// Handle received message (printing it to console in this example)
	fmt.Printf("Received message: %s\n", msg)

	// Define a map to store the unmarshaled JSON data
	var resultMap map[string]string

	// Unmarshal the JSON string into the map
	err = json.Unmarshal([]byte(msg), &resultMap)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return resultMap
}

func WriteToClient(conn *websocket.Conn, message string) bool {
	response := []byte(message)
	if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
		log.Println("Error writing message:", err)
		return true
	} else {
		return false
	}

}
