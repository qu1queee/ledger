package ledger

import "fmt"

//GenerateStatsPerMonth ssas
func GenerateStatsPerMonth(month Month) {
	var totalBills float32
	var totalExpenses float32
	fmt.Printf("Retrieving date for user %s\n", month.User)
	for _, bills := range month.Bills {
		totalBills += bills.Amount
	}
	for _, expenseValue := range month.Expenses {
		for _, expensePrice := range expenseValue {
			totalExpenses += expensePrice.Amount
		}
	}
	fmt.Printf("Total amount for Bills: %v €\n", totalBills)
	fmt.Printf("Total amount for Expenses: %v €\n", totalExpenses)
}
