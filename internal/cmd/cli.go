package cmd

import (
	"log"
	"os"

	"github.com/qu1queee/ledger/pkg/ledger"
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
	}

	app.Name = "ledger"
	app.Usage = "Control your finanzas!"
	app.Action = func(c *cli.Context) error {
		if c.String("config") != "" {
			path := c.String("config")
			ledger.InitializeCurrentMonth(ledger.GetInitialConf(path))
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
