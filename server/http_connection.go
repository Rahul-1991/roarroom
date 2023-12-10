package server

import (
	"log"
	"net/http"
)

func CreateHttpConnection() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
