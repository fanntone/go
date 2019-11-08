package main

import (
	"log"
	"net/http"
	"encoding/json"
)

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

type CoinFilps struct {
	BetNumber		int		`json:"number"`
	BetAmount		float64	`json:"amount"`
	RandomResult	int 	`josn:"result"`
	WinAmount		float64	`json:"winamount"`
}

type RollADices struct {
	BetNumber		[6]int	`json:"number"`
	BetAmount		float64	`json:"amount"`
	RandomResult	int 	`josn:"result"`
	WinAmount		float64	`json:"winamount"`
}

type TwoDices struct {
	BetNumber		[12]int		`json:"number"`
	BetAmount		float64		`json:"amount"`
	RandomResult	int 		`josn:"result"`
	WinAmount		float64		`json:"winamount"`
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

// PlayEtheRoll API
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

// PlayCoinFilp API
func PlayCoinFilp(writer http.ResponseWriter,  request *http.Request) {
	var bet CoinFilps
	if err := json.NewDecoder(request.Body).Decode(&bet); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	x := bet.BetNumber
	y := bet.BetAmount
	bet.RandomResult,bet.WinAmount = CoinFilp(x,y)

	if err := json.NewEncoder(writer).Encode(bet); err != nil {
		log.Fatal(err)
	}
}

// PlayRollADice API
func PlayRollADice(writer http.ResponseWriter,  request *http.Request) {
	var bet RollADices
	if err := json.NewDecoder(request.Body).Decode(&bet); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	var x [6]int
	for i:=0; i<6; i++ {
		x[i] = bet.BetNumber[i]
	}
	y := bet.BetAmount
	bet.RandomResult,bet.WinAmount = RollADice(x,y)

	if err := json.NewEncoder(writer).Encode(bet); err != nil {
		log.Fatal(err)
	}
}

// PlayTwoDice API
func PlayTwoDice(writer http.ResponseWriter,  request *http.Request) {
	var bet TwoDices
	if err := json.NewDecoder(request.Body).Decode(&bet); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	var x [12]int
	for i:=0; i<12; i++ {
		x[i] = bet.BetNumber[i]
	}
	y := bet.BetAmount
	bet.RandomResult,bet.WinAmount = TwoDice(x,y)

	if err := json.NewEncoder(writer).Encode(bet); err != nil {
		log.Fatal(err)
	}
}