package main

import "fmt"

type ChatActor struct {
	mailbox Actor
	hub     *ActorHub
}

func NewChatActor(hub *ActorHub) *ChatActor {
	c := &ChatActor{
		mailbox: make(Actor, 100),
		hub:     hub,
	}
	hub.Register("chat_service", c.mailbox)
	go c.loop()
	return c
}

func (c *ChatActor) loop() {
	for msg := range c.mailbox {
		switch m := msg.(type) {
		case ChatMessage:
			fmt.Printf("[Global Chat] %s: %s\n", m.Sender, m.Content)

			c.hub.mu.RLock()
			for id, ref := range c.hub.store {
				if id != "chat_service" && id != "enemy_boss" {
					ref <- ChatPayload{Sender: m.Sender, Message: m.Content}
				}
			}
			c.hub.mu.RUnlock()
		}
	}
}
