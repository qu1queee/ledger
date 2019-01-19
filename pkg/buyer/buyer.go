package buyer

import (
	"github.com/qu1queee/ledger/pkg/ledger"
	"github.com/qu1queee/ledger/pkg/transaction"
)

// Buyer ...
type Buyer struct {
}

// Tools ...
type Tools struct {
	transaction.Driver
	transaction.Remover
}

// ExecuteDriverTransaction ...
func (Buyer) ExecuteDriverTransaction(td transaction.Driver, amount float64, place string, ledgerBook *ledger.Ledger) {
	td.AddTransaction(amount, place, ledgerBook)
}

// ExecuteRemovalTransaction ...
func (Buyer) ExecuteRemovalTransaction(tr transaction.Remover, id int, ledgerBook *ledger.Ledger) {
	tr.RemoveTransaction(id, ledgerBook)
}

// ModifyLedgers ...
func (Buyer) ModifyLedgers(trd transaction.DriverRemover, amount float64, place string, ledgerBook []*ledger.Ledger) {
	for i := range ledgerBook {
		l := ledgerBook[i]
		trd.AddTransaction(amount, place, l)
	}
}
