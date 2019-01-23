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

package cmd

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/qu1queee/ledger/pkg/ledger"
	"github.com/qu1queee/ledger/pkg/transaction"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ledger",
	Short: "Ledger is a wise finance manager",
	Long: `Ledger is a nice book recorder of your
   different finances, from a typical
   monthly portfolio at home, till your
   ocassionally exchange trades`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		var client transaction.Executor

		cryptoTransaction := ledger.CryptoRecord{}

		billTransactions := ledger.BillRecord{
			Amount: 2,
			Place:  "joder",
			Date:   "no",
			ID:     1,
		}

		ledgers := []*ledger.Ledger{
			{User: "enrique", Transactions: nil},
			{User: "eduardo", Transactions: nil},
		}
		tools := transaction.Tools{
			Add:    transaction.Construct{},
			Remove: transaction.Construct{},
		}

		client.ModifyLedgers(tools, "eduardo", cryptoTransaction, ledgers)
		client.ModifyLedgers(tools, "enrique", billTransactions, ledgers)
		spew.Dump(ledgers)
	},
}

// Execute ..
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
