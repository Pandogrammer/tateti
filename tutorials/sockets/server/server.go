// https://stackoverflow.com/questions/1099672/when-is-it-appropriate-to-use-udp-instead-of-tcp
// https://yalantis.com/blog/how-to-build-websockets-in-go/
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", socket)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func socket(w http.ResponseWriter, r *http.Request) {
	// init
	u := websocket.Upgrader{}
	c, err := u.Upgrade(w, r, nil)
	if err != nil {
		// handle error
	}
	// receive message
	_, message, err := c.ReadMessage()
	fmt.Println(string(message))
	if err != nil {
		// handle error
	}
	// send message
	message = []byte("hola soy el server")
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		// handle error
	}
}
