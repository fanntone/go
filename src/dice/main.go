package main

import (
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/testrand", testrand)
	http.HandleFunc("/PlayEtheRoll", PlayEtheRoll)
	http.HandleFunc("/PlayRollADice", PlayRollADice)
	http.HandleFunc("/PlayTwoDice", PlayTwoDice)
	http.HandleFunc("/PlayCoinFilp", PlayCoinFilp)
 
	err := http.ListenAndServe(":8080", nil)
	if (err != nil) {
		log.Fatal(err)
	}
}