package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(hub *ActorHub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("WS upgrader err: ", err)
			return
		}

		playerID := fmt.Sprintf("player_%d", time.Now().UnixNano()%10000)
		fmt.Printf("Connected new player: %s\n", playerID)

		playerActor := NewPlayerActor(playerID, conn, hub)
		playerActor.Start()

		playerActor.mailbox <- SystemMessage{Text: "Welcome to the game player: " + playerID}

		for {
			_, rawData, err := conn.ReadMessage()
			if err != nil {
				break
			}

			var gameMessage GameMessage
			if err := json.Unmarshal(rawData, &gameMessage); err != nil {
				continue
			}

			switch gameMessage.Type {
			case "move":
				var moveData MovePayload
				if err := json.Unmarshal(gameMessage.Payload, &moveData); err == nil {
					hub.Send(playerID, moveData)
				}
			case "chat":
				var chatData ChatPayload
				if err := json.Unmarshal(gameMessage.Payload, &chatData); err == nil {
					hub.Send("chat_service", ChatMessage{
						Sender:  playerID,
						Content: chatData.Message,
					})
				}
			}
		}

		hub.UnRegister(playerID)
	}
}
