package main

import (
	"encoding/json"
	"log"

	ews "github.com/sacOO7/GoWebsocket"
)

func ListenFromBinance(BinanceToServer chan *ews.Socket, ServerComputation chan *ResponseFromBinance) {

	for {

		socket := <-BinanceToServer

		RecievedFromBinance_JSON := &ResponseFromBinance{
			Data: &Data{},
		}

		socket.OnTextMessage = func(RecievedFromBinance string, socket ews.Socket) {
			log.Printf("Received from Binance!")

			json.Unmarshal([]byte(RecievedFromBinance), RecievedFromBinance_JSON)
			ServerComputation <- RecievedFromBinance_JSON
		}

		//recieve data from *socket

	}
}
