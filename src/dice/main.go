package main

import (
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.HandleFunc("/testrand", testrand)
		http.HandleFunc("/PlayEtheRoll", PlayEtheRoll)
		http.HandleFunc("/PlayRollADice", PlayRollADice)
		http.HandleFunc("/PlayTwoDice", PlayTwoDice)
		http.HandleFunc("/PlayCoinFilp", PlayCoinFilp)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}