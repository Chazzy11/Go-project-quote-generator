// import dependencies

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

// create variables for upper lower bounds of credit score

const creditScoreMin = 500
const creditScoreMax = 900

// create struct for credit score data

type credit_score struct {
	Score int `json:"credit_score"`
}

// create function to generate random credit score
// this function is also http handler function - use http.ResponseWriter and http.Request
// inside function, credit score is generated, passed to struct, and written to HTTP response as json

func generateCreditScore(w http.ResponseWriter, r *http.Request) {
	var creditScore = credit_score{
		CreditScore: (rand.Intn(creditScoreMax-creditScoreMin) + creditScoreMin)
	} 

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(creditScore)

	}

// create main function to create emdpoint and listen on port 8080
