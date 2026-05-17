package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	id   string
	send chan SendMsg
}

func (c *Client) writePump(conn *websocket.Conn, s *Server) {
	defer func() {
		conn.Close()
		s.UnRegister(c.id)
	}()

	for msg := range c.send {
		err := conn.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}

func (c *Client) readPump(conn *websocket.Conn, s *Server) {
	defer func() {
		conn.Close()
		s.UnRegister(c.id)
		fmt.Printf("[Server] %s connection lost\n", c.id)
	}()

	for {
		_, rawData, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var msg ReceiveMsg
		if err := json.Unmarshal(rawData, &msg); err != nil {
			continue
		}

		success := s.SendTo(msg.Receiver, SendMsg{
			Sender:  c.id,
			Content: msg.Content,
		})
		if !success {
			c.send <- SendMsg{
				Sender:  c.id,
				Content: fmt.Sprintf("User %s offline.", msg.Receiver),
			}
		}
	}
}
