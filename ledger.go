package main

/* LEDGER project
*  A collection of financial activities.
 */
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/qu1queee/ledger/user"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

/*
* Add cmd arguments
 */
var enrique user.Ledger

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
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
			initializeCurrentMonth(getInitialConf(path))

		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getCurrentDate() string {
	t := time.Now()
	return t.Format("Mon Jan _2 15:04:05 2006")
}

// Todo
func addExpenses() {}

// Todo
func initializeCurrentMonth(user user.Ledger) {
	fmt.Printf("Current user %s \n", user.Admin)
	for _, bills := range user.Bills {
		fmt.Printf("Bill: %v \n", bills)
	}
	for _, clients := range user.Clients {
		fmt.Printf("Borrower: %v \n", clients.Person)
	}
	// spew.Dump(user)
}

// Todo
func getInitialConf(path string) user.Ledger {
	var admin user.Ledger
	yamlFile, err := ioutil.ReadFile(path)
	check(err)
	err = yaml.Unmarshal(yamlFile, &admin)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return admin
}
