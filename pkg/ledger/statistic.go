package ledger

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

//GetJson shows data in json format
func GetJson(month Month) string {
	jsonMonth, err := json.Marshal(month)
	if err != nil {
		check(err)
	}
	return string(jsonMonth)
}

//GenerateStatsPerMonth will calculate total expenses per current month for both bills and expenses
func GenerateStatsPerMonth(month Month) {
	var tableDataBills, tableDataContent [][]string
	var billsFiltered, expensesFiltered []string
	var totalBills, totalExpense float32

	for key, value := range month.Expenses {
		for _, expenses := range value {
			if expenses.Amount != 0 {
				expensesFiltered = []string{expenses.Date, expenses.Description, key, (strconv.FormatFloat(float64(expenses.Amount), 'f', 1, 32))}
				tableDataContent = append(tableDataContent, expensesFiltered)
				totalExpense += expenses.Amount
			}
		}
	}
	for _, value := range month.Bills {
		billsFiltered = []string{value.Type, value.Description, value.Company, (strconv.FormatFloat(float64(value.Amount), 'f', 1, 32))}
		tableDataBills = append(tableDataBills, billsFiltered)
		totalBills += value.Amount
	}

	fmt.Println("")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Description", "Location", "Amount"})
	table.SetFooter([]string{"", "", "Total", strconv.FormatFloat(float64(totalExpense), 'f', 1, 32)})
	table.SetBorder(true)
	table.AppendBulk(tableDataContent)
	table.Render()
	fmt.Println("")
	fmt.Println("")
	tableBills := tablewriter.NewWriter(os.Stdout)
	tableBills.SetHeader([]string{"Type", "Description", "Location", "Amount"})
	tableBills.SetFooter([]string{"", "", "Total", strconv.FormatFloat(float64(totalBills), 'f', 1, 32)})
	tableBills.SetBorder(true)
	tableBills.AppendBulk(tableDataBills)
	tableBills.Render()
	fmt.Println("")
}
