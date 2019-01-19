package cmd

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/qu1queee/ledger/pkg/buyer"

	"github.com/qu1queee/ledger/pkg/ledger"
	"github.com/qu1queee/ledger/pkg/transaction"
	"github.com/urfave/cli"
)

// Execute uses the urfave cli package to process input data
func Execute() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from yaml `FILE`",
		},
		cli.StringFlag{
			Name:  "user, u",
			Usage: "Specify user to use `USER`",
		},
		cli.StringFlag{
			Name:  "month, m",
			Usage: "Specify a month to collect(default: current)",
		},
		cli.StringFlag{
			Name:  "list, l",
			Usage: "List all collected months",
		},
	}

	app.Name = "ledger"
	app.Usage = "Control your finanzas!"
	app.Action = func(c *cli.Context) error {
		// var AllTransactions []ledger.TransactionBook
		var someone buyer.Buyer

		ledgers := []*ledger.Ledger{
			{User: "enrique", Transactions: nil},
		}
		bt := buyer.Tools{
			Driver:  transaction.Construct{},
			Remover: transaction.Construct{},
		}

		someone.ModifyLedgers(&bt, 100, "sams", ledgers)
		someone.ModifyLedgers(&bt, 40, "edeka", ledgers)
		spew.Dump(ledgers)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
