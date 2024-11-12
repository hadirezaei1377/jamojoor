package protog

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error while connecting to WebSocket:", err)
	}
	defer conn.Close()

	for {

		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		if err != nil {
			log.Println("Error while writing message:", err)
			break
		}

		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading message:", err)
			break
		}
		fmt.Printf("Received from server: %s\n", msg)

		time.Sleep(2 * time.Second)
	}
}
