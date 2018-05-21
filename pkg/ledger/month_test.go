package ledger

import (
	"reflect"
	"testing"
)

func generateFrequentPlaces(frequentPlace string) map[string][]Spend {
	var spend []Spend
	var place map[string][]Spend
	for _, test := range spend {
		test.Amount = 2
		test.Date = "May"
		test.Description = "Nothing"
	}

	place = make(map[string][]Spend)
	place[frequentPlace] = spend

	return place
}

func generateBills(billCompany string, billAmount int) []Bill {
	var bills []Bill
	var bill Bill
	bill.Company = billCompany
	bill.Type = "monthly"
	bill.Amount = billAmount
	bill.Description = "nothing"
	bills = append(bills, bill)
	return bills
}

func generateBorrowers(loan int, iteration string, person string) []Borrower {
	var borrowers []Borrower
	var borrower Borrower
	borrower.Loan = loan
	borrower.Type = iteration
	borrower.Person = person
	borrowers = append(borrowers, borrower)
	return borrowers
}

func generateMonthForTest(userName string, billCompany string, billAmount int, userFrequentPlace string) Month {
	var month Month
	month.User = userName
	month.Bills = generateBills(billCompany, billAmount)
	month.Expenses = generateFrequentPlaces(userFrequentPlace)
	return month
}

func generateLedgerForTest(userName string, userSalary int, borrowerLoan int, borrowerIteration string, borrowerPerson string, billCompany string, billAmount int, userFrequentPlace string) Ledger {
	var ledger Ledger
	ledger.Admin = userName
	ledger.Salary = userSalary
	ledger.Clients = generateBorrowers(borrowerLoan, borrowerIteration, borrowerPerson)
	ledger.Bills = generateBills(billCompany, billAmount)
	ledger.Places = []string{userFrequentPlace}
	return ledger
}

func TestInitializeCurrentMonth(t *testing.T) {
	var tests = []struct {
		ledger Ledger
		want   Month
	}{
		{generateLedgerForTest("user_X", 100, 20, "monthly", "friend_Y", "Netflix", 10, "taco_bell"), generateMonthForTest("user_X", "Netflix", 10, "taco_bell")},
	}
	for _, test := range tests {
		if test.ledger.Admin != test.want.User {
			t.Errorf("Expected user %v, got %v", test.ledger.Admin, test.want.User)
		} else if !reflect.DeepEqual(test.ledger.Bills, test.want.Bills) {
			t.Errorf("Expected bill %v, got %v", test.ledger.Bills, test.want.Bills)
		}

	}
}
