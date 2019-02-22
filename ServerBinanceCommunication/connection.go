package ServerBinanceCommunication

import (
	"log"
	"golang.org/x/net/websocket"
)

func ConnectToBinance(query string) (*websocket.Conn, error){
	origin := "http://stream.binance.com:9443/"
	url := "wss://stream.binance.com:9443/stream/streams=" + query

	log.Println("Connecting to Binance... Query: "+ query)
	ws, err := websocket.Dial(url, "", origin)
	
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Connected to Binance!")
	return ws, nil

}