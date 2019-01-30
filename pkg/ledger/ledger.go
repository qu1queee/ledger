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

package ledger

// Records ...
type Records interface {
	GetState()
}

// CryptoRecord should illustrate a trade position, based on
// - amount of coins used
// - type of crypto coin
// - state of the trade
// - date of execution
type CryptoRecord struct {
	Amount float64
	Coin   string
	Date   string
	State  string
}

// BillRecord should illustrate a normall bill, based on
// - amount, aka price
// - place where the buy took place
// - date of payment
// - ID randomly generated number associated to the buy
type BillRecord struct {
	Amount float64
	Place  string
	Date   string
	ID     int
}

// GetState to implement the interface LedgerContents
func (CryptoRecord) GetState() {
}

// GetState to implement the interface LedgerContents
func (BillRecord) GetState() {
}

// Ledger where all transaction should be recorded
type Ledger struct {
	Name         string
	User         string
	Transactions []Records
}
