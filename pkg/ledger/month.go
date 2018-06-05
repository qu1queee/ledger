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
	Amount      int    `yaml:"amount"`
	Date        string `yaml:"date"`
	Description string `yaml:"description"`
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
func UpdateExistingMonth(month Month, user Ledger, path string) {
	yamlMonthFile, err := ioutil.ReadFile(path)
	check(err)
	err = yaml.Unmarshal(yamlMonthFile, &month)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, frequentPlaces := range user.Places {
		if month.Expenses[frequentPlaces] == nil {
			var missingExpense Spend
			month.Expenses[frequentPlaces] = append(month.Expenses[frequentPlaces], missingExpense)
		}
	}
	b, err := yaml.Marshal(month)
	check(err)
	errs := ioutil.WriteFile(path, b, 0644)
	check(errs)
}

// InitializeCurrentMonth will init the YAML file for the current month, based on the main config YAML
func InitializeCurrentMonth(user Ledger) {
	var month Month
	var currentMonth time.Month
	_, currentMonth, _ = time.Now().UTC().Date()
	currentMonthLowerCase := strings.ToLower(currentMonth.String())
	path := os.Getenv("HOME") + ledgerConfigDirName + "/" + currentMonthLowerCase + "/" + currentMonthLowerCase + ".yml"

	if _, err := os.Stat(path); err == nil {
		fmt.Printf("File %s, exists. \n", path)
		UpdateExistingMonth(month, user, path)
	} else {
		//TODO improve code, too many lines
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
		MarshallMonth(month, path)
	}

}
