package main

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"golang.org/x/net/websocket"
)

type Trade struct {
	Symbol    string `json:"symbol"`
	TradeID   string `json:"tradeId"`
	Price     string `json:"price"`
	Size      string `json:"side"`
	Timestamp int64  `json:"timestamp"`
}

type Message struct {
	Topic  string  `json:"topic"`
	Symbol string  `json:"symbol"`
	Data   []Trade `json:"data"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func() {

	},
}

func (t *Trade) ToJson() ([]byte, error) {
	return json.Marshal(t)
}

func main() {

}
