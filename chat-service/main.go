package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("Remote incoming connection from client: ", ws.RemoteAddr())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) handleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client orderbook feed:", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("orderbook data -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error: ", err)
			continue
		}
		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error: ", err)
			}
		}(ws)
	}
}

func main() {
	fmt.Println("chat-service")
	server := NewServer()

	wsServer := websocket.Server{
		Handshake: func(config *websocket.Config, req *http.Request) error {
			return nil
		},
		Handler: websocket.Handler(server.handleWS),
	}
	http.Handle("/ws", wsServer)

	wsOrderbook := websocket.Server{
		Handshake: func(config *websocket.Config, req *http.Request) error {
			return nil
		},
		Handler: websocket.Handler(server.handleWSOrderbook),
	}
	http.Handle("/orderbookfeed", wsOrderbook)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(-1)
	}
}
