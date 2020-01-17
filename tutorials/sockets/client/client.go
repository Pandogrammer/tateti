package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

func main() {
	client()
}

func client() {
	// init
	// schema – can be ws:// or wss://
	// host, port – WebSocket server
	schema := "ws"
	host := "localhost"
	port := "8080"
	u := url.URL{
		Scheme: schema,
		Host:   host + ":" + port,
		Path:   "/",
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		// handle error
	}
	// send message
	message := []byte("hola soy el cliente")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		// handle error
	}
	// receive message
	_, message, err = c.ReadMessage()
	fmt.Println(string(message))
	if err != nil {
		// handle error
	}
}