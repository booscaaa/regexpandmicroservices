package texto

import (
	"encoding/json"
	"log"
	"net/http"
)

//MatchText .
func MatchText(w http.ResponseWriter, r *http.Request) {
	w = setOrigins(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		var texto Texto
		if err := json.NewDecoder(r.Body).Decode(&texto); err != nil {
			log.Println("Nenhum parametro enviado")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		} else {
			defer r.Body.Close()
			response := texto.MatchText()
			w.Write(response)
		}
	}
}

//MatchNumber .
func MatchNumber(w http.ResponseWriter, r *http.Request) {
	w = setOrigins(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		var texto Texto
		if err := json.NewDecoder(r.Body).Decode(&texto); err != nil {
			log.Println("Nenhum parametro enviado")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		} else {
			defer r.Body.Close()
			response := texto.MatchNumber()
			w.Write(response)
		}
	}
}

//FindNumber .
func FindNumber(w http.ResponseWriter, r *http.Request) {
	w = setOrigins(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		var texto Texto
		if err := json.NewDecoder(r.Body).Decode(&texto); err != nil {
			log.Println("Nenhum parametro enviado")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		} else {
			defer r.Body.Close()
			response := texto.FindNumber()
			w.Write(response)
		}
	}
}

func setOrigins(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, PUT")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	return w
}
