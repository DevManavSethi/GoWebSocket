package main

import (
	"log"

	"golang.org/x/net/websocket"
)

func SendToClient(ws *websocket.Conn, ServerToClient chan []byte) {
	for {

		MessageToClient_byte := <-ServerToClient
		log.Println("Sending message to client")
		MessageToClient := string(MessageToClient_byte)
		websocket.Message.Send(ws, MessageToClient)

	}

}
