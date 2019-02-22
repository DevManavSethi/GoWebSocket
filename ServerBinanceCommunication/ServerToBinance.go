package ServerBinanceCommunication

import (
	"log"
	"golang.org/x/net/websocket"
)

type Data struct {
	Type         string  `json:"e"`
	Time         float64 `json:"E"`
	Symbol       string  `json:"s"`
	TradeID      int     `json:"t"`
	Price        string  `json:"p"`
	Quantity     string  `json:"q"`
	
	BuyerTradeId int   `json:"b"`
	SellerTradeID  int   `json:"a"`
	
	Timestamp    float64 `json:"T"`
	IsbuyerMarketMaker      bool    `json:"m"`
	Ignore bool `json:"M"`
}


type ResponseFromBinance struct{
	Stream string `json:"stream"`
	Data *Data `json:"data"`
}

func ListenFromBinance(ws *websocket.Conn) (*ResponseFromBinance, error){


	RecievedFromBinance_JSON := &ResponseFromBinance{
		Data: &Data{},
		}
	

	if err := websocket.JSON.Receive(ws, &RecievedFromBinance_JSON); err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("Received from Binance!")
	return RecievedFromBinance_JSON, nil

}