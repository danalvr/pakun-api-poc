package models

import "time"

type Transaction struct {
	Amount float64 `firestore:"amount"`
	Note string `firestore:"note"`
	Sender string `firestore:"sender"`
	Source string `firestore:"source"`
	Timestamp time.Time `firestore:"timestamp"`
	Type string `firestore:"type"`
}