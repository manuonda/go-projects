package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/net/websocket"
)

type Trade struct {
	Symbol    string `json:"symbol"`
	TradeID   string `json:"tradeId"`
	Price     string `json:"price"`
	Size      string `json:"side"`
	Timestamp int64  `json:"timestamp"`
	Side      string `json:"side"`
}

type Message struct {
	Topic  string  `json:"topic"`
	Symbol string  `json:"symbol"`
	Data   []Trade `json:"data"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (t *Trade) ToJson() ([]byte, error) {
	return json.Marshal(t)
}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var symbols = []string{"BTC_USDT", "ETH_USDT", "BNB_USDT", "ADA_USDT", "XRP_USDT", "DOGE_USDT", "DOT_USDT", "UNI_USDT", "SOL_USDT", "LINK_USDT"}
var lastPrice = 50000.0

func generateMockTrade(timestamppb int64) Trade {
	priceChange := rnd.Float64()*20 - 10 //cambio de precio +10 o -10
	lastPrice += priceChange
	return Trade{
		Symbol:    symbols[rnd.Intn(len(symbols))],
		TradeID:   strconv.Itoa(rnd.Intn(1000000)),
		Price:     strconv.FormatFloat(lastPrice, 'f', 2, 64),
		Size:      strconv.FormatFloat(0.01+rnd.Float64()*10, 'f', 2, 64),
		Side:      []string{"BUY", "SHELL"}[rnd.Intn(2)],
		Timestamp: timestamp,
	}
}

func main() {

}
