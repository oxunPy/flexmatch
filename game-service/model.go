package main

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type GameMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type ChatPayload struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

type MovePayload struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ConnMessage struct {
	PlayerID string
	Conn     *websocket.Conn
}

type DisConnMessage struct {
	PlayerID string
}

type ChatMessage struct {
	Sender  string `json:"sender"`
	Content string `json: "content"`
}

type SystemMessage struct {
	Text string
}
