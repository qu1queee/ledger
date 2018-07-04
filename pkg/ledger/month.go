package ledger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

const ledgerConfigDirName string = "/.ledger"

// Spend struct that defines a single spend action
type Spend struct {
	Amount      float32 `yaml:"amount"`
	Date        string  `yaml:"date"`
	Description string  `yaml:"description"`
}

// Month struct that defines a month config YAML file
type Month struct {
	User     string             `yaml:"user"`
	Bills    []Bill             `yaml:"bills"`
	Expenses map[string][]Spend `yaml:"frequent_places"`
}

// MarshallMonth func to generate the current month config YAML file
func MarshallMonth(month Month, path string) {
	b, err := yaml.Marshal(month)
	check(err)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		errs := ioutil.WriteFile(path, b, 0644)
		check(errs)
	}
}

//UpdateExistingMonth will update month config file
func UpdateExistingMonth(month Month, user Ledger, path string) Month {
	yamlMonthFile, err := ioutil.ReadFile(path)
	check(err)
	err = yaml.Unmarshal(yamlMonthFile, &month)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, frequentPlaces := range user.Places {
		if month.Expenses[frequentPlaces] == nil {
			fmt.Printf("New frequent place detected: %v, updating. \n", frequentPlaces)
			var missingExpense Spend
			month.Expenses[frequentPlaces] = append(month.Expenses[frequentPlaces], missingExpense)
		}
	}
	b, err := yaml.Marshal(month)
	check(err)
	errs := ioutil.WriteFile(path, b, 0644)
	check(errs)
	return month
}

// CreateMonth create Month file with configuration file properties
func CreateMonth(user Ledger, month Month) Month {
	var mymap = make(map[string][]Spend)
	var spend = make([]Spend, 1)
	spend[0].Description = "Please provide a description (optional)"
	month.Expenses = mymap
	month.Bills = user.Bills
	month.User = user.Admin
	for _, bills := range user.Places {
		mymap[bills] = spend
	}
	month.Expenses = mymap
	return month
}

// InitializeCurrentMonth will init the YAML file for the current month, based on the main config YAML
func InitializeCurrentMonth(user Ledger, desiredMonth string) {
	var month Month
	var currentMonthLowerCase string
	var currentMonth time.Month
	_, currentMonth, _ = time.Now().UTC().Date()
	if desiredMonth != "" {
		currentMonthLowerCase = strings.ToLower(desiredMonth)
	} else {
		currentMonthLowerCase = strings.ToLower(currentMonth.String())
	}
	path := os.Getenv("HOME") + ledgerConfigDirName + "/" + currentMonthLowerCase + "/" + currentMonthLowerCase + ".yml"

	if _, err := os.Stat(path); err == nil {
		fmt.Printf("%s %s %s\n", prettyBlueBold("File"), prettyGreen(path), prettyBlueBold("exists."))
		month = UpdateExistingMonth(month, user, path)
	} else {
		month = CreateMonth(user, month)
		MarshallMonth(month, path)
	}
	GenerateStatsPerMonth(month)
}
