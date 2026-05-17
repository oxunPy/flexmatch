package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type PlayerActor struct {
	id      string
	conn    *websocket.Conn
	mailbox Actor
	hub     *ActorHub
	x, y    int
}

func NewPlayerActor(id string, conn *websocket.Conn, hub *ActorHub) *PlayerActor {
	return &PlayerActor{
		id:      id,
		conn:    conn,
		mailbox: make(Actor, 100),
		hub:     hub,
	}
}

func (p *PlayerActor) Start() {
	p.hub.Register(p.id, p.mailbox)
	go p.loop()
}

func (p *PlayerActor) loop() {
	defer func() {
		_ = p.conn.Close()
		p.hub.UnRegister(p.id)
		fmt.Printf("[PlayerActor %s] exit\n", p.id)
	}()

	for msg := range p.mailbox {
		switch m := msg.(type) {
		case MovePayload:
			p.x = m.X
			p.y = m.Y
			fmt.Printf("[Player %s] move to the pos: (X:%d, Y:%d)\n", p.id, p.x, p.y)
		case SystemMessage:
			bytes, err := json.Marshal(map[string]string{
				"type":    "system",
				"message": m.Text,
			})
			if err != nil {
				_ = p.conn.WriteJSON(string(bytes))
			}
		case ChatPayload:
			bytes, err := json.Marshal(map[string]string{
				"type":    "chat",
				"sender":  m.Sender,
				"message": m.Message,
			})
			if err != nil {
				_ = p.conn.WriteJSON(string(bytes))
			}
		}
	}
}
