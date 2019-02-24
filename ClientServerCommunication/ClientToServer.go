package main

import (
	"log"

	"golang.org/x/net/websocket"
)

func ListenFromClient(ws *websocket.Conn, ClientToServer chan string) {

	for {
		log.Println("Listening to client...")

		//read message from client
		var RecievedFromClient string

		websocket.Message.Receive(ws, &RecievedFromClient)
		log.Println(RecievedFromClient)
		query := RecievedFromClient
		log.Println(query)

		ClientToServer <- query

	}

}
