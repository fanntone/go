package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type Resp struct {
    Code    string `json:"code"`
    Msg     string `json:"msg"`
}

type  Auth struct {
    Username string `json:"username"`
    Pwd      string `json:"password"`
}

type Random struct {
	Range	int `json:"range"`
	Result  int `json:"reslut"`
}

type EtheRoll struct {
	BetNumber		int		`json:"number"`
	BetAmount		float64	`json:"amount"`
	RandomResult	int 	`josn:"result"`
	WinAmount		float64	`json:"winamount"`
}

//post接口接收json數據
func testrand(writer http.ResponseWriter,  request *http.Request) {
	var rand Random
	if err := json.NewDecoder(request.Body).Decode(&rand); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}

	rand.Range = rand.Range
	rand.Result = GetRandom(rand.Range)
	if err := json.NewEncoder(writer).Encode(rand); err != nil {
		log.Fatal(err)
	}
}

//
func PlayEtheRoll(writer http.ResponseWriter,  request *http.Request) {
	var bet EtheRoll
	if err := json.NewDecoder(request.Body).Decode(&bet); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	x := bet.BetNumber
	y := bet.BetAmount
	bet.RandomResult,bet.WinAmount = Etheroll(x,y)

	if err := json.NewEncoder(writer).Encode(bet); err != nil {
		log.Fatal(err)
	}
}