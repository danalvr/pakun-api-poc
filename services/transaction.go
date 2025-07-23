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

type MonthlySummary struct {
	Month string `json:"month" firestore:"month"`
	Income float64 `json:"income" firestore:"income"`
	Expense float64 `json:"expense" firestore:"expense"`
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

func GetMonthlySummary(identifier string, from, to time.Time) ([]MonthlySummary, error) {
	ctx := context.Background()
	iter := firebase.Client.Collection("transactions").
		Where("sender", "==", identifier).
		Where("timestamp", ">=", from).
		Where("timestamp", "<", to.AddDate(0, 1, 0)).
		Documents(ctx)

	summaryMap := make(map[string]*MonthlySummary)

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var txn ListTransactions
		if err := doc.DataTo(&txn); err != nil {
			continue
		}

		monthKey := txn.Timestamp.Format("2006-01")

		if _, exists := summaryMap[monthKey]; !exists {
			summaryMap[monthKey] = &MonthlySummary{
				Month: monthKey,
			}
		}

		if txn.Type == "income" {
			summaryMap[monthKey].Income += txn.Amount
		} else if txn.Type == "expense" {
			summaryMap[monthKey].Expense += txn.Amount
		}
	}

	var summaries []MonthlySummary
	for d := from; !d.After(to); d = d.AddDate(0, 1, 0) {
		key := d.Format("2006-01")
		if sum, ok := summaryMap[key]; ok {
			summaries = append(summaries, *sum)
		} else {
			summaries = append(summaries, MonthlySummary{
				Month: key,
				Income: 0,
				Expense: 0,
			})
		}
	}

	return summaries, nil
}