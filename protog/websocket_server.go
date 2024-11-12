package protog

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("New client connected")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error while reading message:", err)
			break
		}
		fmt.Printf("Received message: %s\n", msg)

		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			fmt.Println("Error while writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	fmt.Println("WebSocket server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
