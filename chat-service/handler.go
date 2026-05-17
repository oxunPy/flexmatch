package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWS(s *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		senderID := r.URL.Query().Get("sender")
		if senderID == "" {
			http.Error(w, "Sender ID required", http.StatusBadRequest)
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Upgrader err: ", err)
			return
		}

		client := &Client{
			id:   senderID,
			send: make(chan SendMsg, 100),
		}

		s.Register(senderID, client)
		fmt.Printf("[Server] %s connected \n", senderID)

		go client.readPump(conn, s)
		go client.writePump(conn, s)
	}
}
