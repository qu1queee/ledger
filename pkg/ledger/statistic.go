package ledger

import (
	"fmt"
)

//GenerateStatsPerMonth will calculate total expenses per current month for both bills and expenses
func GenerateStatsPerMonth(month Month) {
	var totalBills float32
	var totalExpenses float32
	fmt.Printf("%s %s\n", prettyBlueBold("Retrieving data for user"), prettyGreen(month.User))
	for _, bills := range month.Bills {
		totalBills += bills.Amount
	}
	for _, expenseValue := range month.Expenses {
		for _, expensePrice := range expenseValue {
			totalExpenses += expensePrice.Amount
		}
	}
	fmt.Printf("%s %v%s\n", prettyBlueBold("Total amount for Bills:"), prettyGreen(totalBills), prettyGreen("€"))
	fmt.Printf("%s %v%s\n", prettyBlueBold("Total amount for Expenses:"), prettyGreen(totalExpenses), prettyGreen("€"))
}
