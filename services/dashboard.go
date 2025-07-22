package services

import (
	"context"
	"pakun-api-poc/firebase"
	"pakun-api-poc/models"
)

type DashboardData struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
}

func GetDashboardData(identifier string) (DashboardData, error) {
	ctx := context.Background()
	iter := firebase.Client.Collection("transactions").
		Where("sender", "==", identifier).
		Documents(ctx)

	var income, expense float64

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var txn models.Transaction
		if err := doc.DataTo(&txn); err != nil {
			continue
		}

		if txn.Type == "income" {
			income += txn.Amount
		} else if txn.Type == "expense" {
			expense += txn.Amount
		}
	}

	return DashboardData{
		TotalIncome: income,
		TotalExpense: expense,
	}, nil
}