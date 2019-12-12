package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Slot struct {
	Bet			int 	`json:"bet"`
	WinAmount	int 	`json:"winAmount"`
	Result 		[]int	`json:"result"`
}

// PlaySlot API
func PlaySlot(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "OPTIONS" {
		return
	}
	var bet Slot
	if err := json.NewDecoder(request.Body).Decode(&bet); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	x := bet.Bet
	bet.Result, bet.WinAmount = BetOneLine(x)

	if err := json.NewEncoder(writer).Encode(bet); err != nil {
		log.Fatal(err)
	}
}


