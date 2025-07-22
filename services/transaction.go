package services

import (
	"context"
	"pakun-api-poc/firebase"
	"time"

	"cloud.google.com/go/firestore"
)

type ListTransactions struct {
	Amount    float64 `json:"amount" firestore:"amount"`
	Note      string `json:"note" firestore:"note"`
	Type      string `json:"type" firestore:"type"`
	Timestamp time.Time `json:"timestamp" firestore:"timestamp"`
}

func GetTransactions(identifier string) ([]ListTransactions, error) {
	ctx := context.Background()
	iter := firebase.Client.Collection("transactions").
		Where("sender", "==", identifier).
		OrderBy("timestamp", firestore.Desc).
		Documents(ctx)

	var transactions []ListTransactions

	for {
		doc, err := iter.Next()
		if err != nil {
			break;
		}
		var txn ListTransactions
		if err := doc.DataTo(&txn); err == nil {
			transactions = append(transactions, txn)
		}
	}

	return transactions, nil
}