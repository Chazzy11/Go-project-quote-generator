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