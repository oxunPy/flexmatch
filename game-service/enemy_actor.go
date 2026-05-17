package main

import (
	"fmt"
	"time"
)

type EnemyActor struct {
	name    string
	mailbox Actor
	hub     *ActorHub
	hp      int
}

func NewEnemyActor(name string, hub *ActorHub) *EnemyActor {
	e := &EnemyActor{
		name:    name,
		mailbox: make(Actor, 100),
		hub:     hub,
		hp:      1000,
	}
	hub.Register(name, e.mailbox)
	go e.loop()
	go e.aiRoutine()
	return e
}

func (e *EnemyActor) loop() {
	for msg := range e.mailbox {
		switch m := msg.(type) {
		case SystemMessage:
			fmt.Printf("[%s LOG] %s\n", e.name, m.Text)
		}
	}
}

func (e *EnemyActor) aiRoutine() {
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		e.hub.Send("chat_service", ChatMessage{
			Sender:  e.name,
			Content: "I kill you all.",
		})
	}
}
