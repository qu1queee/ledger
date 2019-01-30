// Copyright © 2019 Enrique Encalada
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transaction

import (
	"github.com/qu1queee/ledger/pkg/ledger"
)

// Add will put transactions into a Ledger book
type Add interface {
	// AddTransaction(amount float64, place string, ledger *ledger.Ledger)
	AddTransaction(transaction ledger.Records, ledger *ledger.Ledger)
}

// Remove will delete recorded transactions from a Ledger book
type Remove interface {
	RemoveTransaction(id int, ledger *ledger.Ledger)
}

// AddRemove implements both single method interfaces
type AddRemove interface {
	Add
	Remove
}

// Construct implements Add/Remove/AddRemove interfaces
type Construct struct {
}

// AddTransaction based on the amount, place and date
func (Construct) AddTransaction(transaction ledger.Records, ledgerBook *ledger.Ledger) {
	// newTransaction := ledger.TransactionBook{
	// 	Amount: amount, Place: place, Date: "fakeasdate", Id: 2,
	// }
	ledgerBook.Transactions = append(ledgerBook.Transactions, transaction)
}

// RemoveTransaction based on an id
func (Construct) RemoveTransaction(id int, ledgerBook *ledger.Ledger) {
	ledgerBook.Transactions = append(ledgerBook.Transactions[:id], ledgerBook.Transactions[id+1:]...)
}

// Executor ...
type Executor struct {
}

// Tools ...
type Tools struct {
	Add
	Remove
}

// ExecuteAddTransaction ...
func (Executor) ExecuteAddTransaction(td Add, transaction ledger.Records, ledgerBook *ledger.Ledger) {
	td.AddTransaction(transaction, ledgerBook)
}

// ExecuteRemoveTransaction ...
func (Executor) ExecuteRemoveTransaction(tr Remove, id int, ledgerBook *ledger.Ledger) {
	tr.RemoveTransaction(id, ledgerBook)
}

// ModifyLedgers ...
func (Executor) ModifyLedgers(trd AddRemove, ledgerName string, transaction ledger.Records, ledgerBook []*ledger.Ledger) {
	for i := range ledgerBook {
		if ledgerBook[i].User == ledgerName {
			l := ledgerBook[i]
			trd.AddTransaction(transaction, l)
		}
	}
}
