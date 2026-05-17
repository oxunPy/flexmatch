package main

import "sync"

type Actor chan interface{}

type ActorHub struct {
	mu    sync.RWMutex
	store map[string]Actor
}

func NewActorHub() *ActorHub {
	return &ActorHub{
		store: make(map[string]Actor),
	}
}

func (ah *ActorHub) Register(id string, ref Actor) {
	ah.mu.Lock()
	defer ah.mu.Unlock()
	ah.store[id] = ref
}

func (ah *ActorHub) UnRegister(id string) {
	ah.mu.Lock()
	defer ah.mu.Unlock()
	if ref, exists := ah.store[id]; exists {
		close(ref)
		delete(ah.store, id)
	}
}

func (ah *ActorHub) Send(id string, message interface{}) bool {
	ah.mu.Lock()
	defer ah.mu.Unlock()
	ref, exists := ah.store[id]
	if exists {
		ref <- message
		return true
	}
	return false
}
