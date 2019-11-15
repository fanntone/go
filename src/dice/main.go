package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {

	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/testrand", testrand)
	mux.HandleFunc("/PlayEtheRoll", PlayEtheRoll)
	mux.HandleFunc("/PlayRollADice", PlayRollADice)
	mux.HandleFunc("/PlayTwoDice", PlayTwoDice)
	mux.HandleFunc("/PlayCoinFilp", PlayCoinFilp)

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
