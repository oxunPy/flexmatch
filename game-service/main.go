package main

import (
	"fmt"
	"net/http"
)

func main() {
	hub := NewActorHub()
	NewChatActor(hub)
	NewEnemyActor("enemy_boss", hub)

	http.HandleFunc("/ws", handleWebSocket(hub))

	fmt.Println("Game server is launched on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server launching err:", err)
	}
}
