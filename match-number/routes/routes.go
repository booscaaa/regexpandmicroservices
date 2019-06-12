package routes

import (
	"../texto"
	"github.com/gorilla/mux"
)

//Routes .
func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v1/texto/match-text", texto.MatchText).Methods("POST", "OPTIONS")
	r.HandleFunc("/v1/texto/match-number", texto.MatchNumber).Methods("POST", "OPTIONS")
	r.HandleFunc("/v1/texto/find-number", texto.FindNumber).Methods("POST", "OPTIONS")

	return r
}
