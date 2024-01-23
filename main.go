// import dependencies

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

// create variables for upper lower bounds of credit score

const quoteNumMin = 1
const quoteNumMax = 9

// create struct for credit score data

type quotenumber struct {
	QuoteNumber int `json:"quotenumber"`
}

// create function to generate random credit score
// this function is also http handler function - use http.ResponseWriter and http.Request
// inside function, credit score is generated, passed to struct, and written to HTTP response as json

func generateQuoteNumber(w http.ResponseWriter, r *http.Request) {
	var quote = quotenumber{
		QuoteNumber: (rand.Intn(quoteNumMax-quoteNumMin) + quoteNumMin),
	} 

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quote)

	}

// create main function to create emdpoint and listen on port 8080

func handleRequests() {
	http.Handle("/quotenumber", http.HandlerFunc(generateQuoteNumber))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
