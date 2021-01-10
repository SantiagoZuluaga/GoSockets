package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {

	var err error
	var messageType int
	var message []byte

	for {
		messageType, message, err = conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(message))

		err = conn.WriteMessage(messageType, []byte("Hi from Backend"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("CLIENT SUCCESSFULLY CONNECTED")
	reader(ws)
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}

	http.HandleFunc("/ws", wsEndpoint)
	http.ListenAndServe(port, nil)
}
