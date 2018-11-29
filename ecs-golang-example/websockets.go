package main

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			retStr := strings.ToUpper(string(msg))
			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), retStr)

			// Write message back to browser
			if err = conn.WriteMessage(msgType, []byte(retStr)); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websocket/index.html")
	})

	http.ListenAndServe(":8080", nil)
}