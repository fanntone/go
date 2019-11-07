package main

import (
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.HandleFunc("/testrand", testrand)
		http.HandleFunc("/PlayEtheRoll", PlayEtheRoll)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}