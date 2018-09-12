package main

import (
	//"sync"
	//"strings"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

const (
	MaxQueue = 100 // Size of message queue
)

type ChangeStats struct {
	Symbol             string `json:"s"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	PrevClosePrice     string `json:"x"`
	BidPrice           string `json:"b"`
	AskPrice           string `json:"a"`
	HighPrice          string `json:"h"`
	Volume             string `json:"v"`
}

func main() {

	address := fmt.Sprintf("wss://stream.binance.com:9443/ws/!ticker@arr")

	var wsDialer websocket.Dialer
	wsConn, _, err := wsDialer.Dial(address, nil)
	if err != nil {
		panic(err)
	}
	defer wsConn.Close()
	log.Println("Dialed:", address)

	for {
		_, message, err := wsConn.ReadMessage()

		if err != nil {
			log.Println("[ERROR] ReadMessage:", err)
		}

		var msg []ChangeStats
		fmt.Println("Received message !")

		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("[ERROR] Parsing:", err)
			continue
		}
		fmt.Println(msg)
	}
}
