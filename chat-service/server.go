package main

import (
	"sync"
)

type Server struct {
	mu    sync.RWMutex
	conns map[string]*Client
}

func NewServer() *Server {
	return &Server{
		conns: make(map[string]*Client),
	}
}

func (s *Server) Register(id string, client *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.conns[id] = client
}

func (s *Server) UnRegister(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if client, exists := s.conns[id]; exists {
		close(client.send)
		delete(s.conns, id)
	}
}

func (s *Server) SendTo(receiver string, msg SendMsg) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	client, exists := s.conns[receiver]
	if !exists {
		return false
	}

	client.send <- msg
	return true
}
