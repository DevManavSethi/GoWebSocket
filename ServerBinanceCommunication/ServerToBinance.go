package main

import (
	"log"

	ews "github.com/sacOO7/GoWebsocket"
)

func ConnectToBinance(ClientToServer chan string, BinanceToServer chan *ews.Socket) {

	for {

		query := <-ClientToServer

		url := "wss://stream.binance.com:9443/stream?streams=" + query
		log.Println("Connecting to socket: " + url)

		socket := ews.New(url)

		socket.OnConnected = func(socket ews.Socket) {
			log.Println("Connected to Binance!")
		}

		socket.Connect()

		BinanceToServer <- &socket

	}

}
