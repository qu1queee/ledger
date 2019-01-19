package transaction

import (
	"github.com/qu1queee/ledger/pkg/ledger"
)

// Driver type will add a new transaction
type Driver interface {
	AddTransaction(amount float64, place string, ledger *ledger.Ledger)
}

// Remover type will remove an existing transaction
type Remover interface {
	RemoveTransaction(id int, ledger *ledger.Ledger)
}

// DriverRemover can add or delete transactions
type DriverRemover interface {
	Driver
	Remover
}

// Construct does not represent an state,
// useful to implement our interfaces
type Construct struct {
}

// AddTransaction based on the amount, place and date
func (Construct) AddTransaction(amount float64, place string, ledgerBook *ledger.Ledger) {
	newTransaction := ledger.TransactionBook{
		Amount: amount, Place: place, Date: "fakeasdate", Id: 2,
	}
	ledgerBook.Transactions = append(ledgerBook.Transactions, newTransaction)
}

// RemoveTransaction based on an id
func (Construct) RemoveTransaction(id int, ledgerBook *ledger.Ledger) {
	ledgerBook.Transactions = append(ledgerBook.Transactions[:id], ledgerBook.Transactions[id+1:]...)
}

// check
func check(e error) {
	if e != nil {
		panic(e)
	}
}
