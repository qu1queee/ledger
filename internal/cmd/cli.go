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
		cli.StringFlag{
			Name:  "user, u",
			Usage: "Specify user to use `USER`",
		},
	}

	app.Name = "ledger"
	app.Usage = "Control your finanzas!"
	app.Action = func(c *cli.Context) error {
		var path string
		var desiredUser string
		if c.String("user") != "" || c.String("config") != "" {
			desiredUser = c.String("user")
			path = c.String("config")
			ledger.GetInitialConf(path, desiredUser)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
