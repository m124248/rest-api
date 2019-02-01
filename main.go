package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Phrase struct {
	/*ID    string `json:"id"`*/
	Greeting string `json:"greeting"`
}

var phrase Phrase

func GetPhrase(w http.ResponseWriter, r *http.Request) {
	greetingPhrase = "Catherine!"
	phrase = Phrase{Greeting: greetingPhrase}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(phrase)
	if err != nil {
		log.Panic("Oh No")
	}
}

var greetingPhrase string

// fun main()
func main() {
	//init router
	router := mux.NewRouter()





	//route handles & endpoints
	router.HandleFunc("/phrase", GetPhrase).Methods("GET")

	//start server
	log.Fatal(http.ListenAndServe(":8000", cors.Default().Handler(router)))

}

//DEVL02TD080G8WL:srv catherineluning$ go run main.go
//todo localhost:8000/phrase
