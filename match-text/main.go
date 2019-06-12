package main

import (
	"fmt"
	"net/http"

	"./texto"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/texto/match-text", texto.MatchText).Methods("POST", "OPTIONS")

	fmt.Println("5000")
	fmt.Println(http.ListenAndServe(":5000", handlers.CompressHandler(r)))
}
