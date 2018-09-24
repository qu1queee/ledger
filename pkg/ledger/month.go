package ledger

import (
	"io/ioutil"
	"log"
	"os"

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

// TODO logic needs to be reworked
// MarshallMonth func to generate the current month config YAML file
func MarshallMonth(month Month, path string) {
	b, err := yaml.Marshal(month)
	check(err)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		errs := ioutil.WriteFile(path, b, 0644)
		check(errs)
	}
}

// TODO logic needs to be reworked
//UpdateExistingMonth will update month config file
func UpdateExistingMonth(user Ledger, path string) Month {
	var month Month
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
	return month
}

// CreateMonth create Month file with configuration file properties
// TODO logic needs to be reworked
func CreateMonth(user Ledger) Month {
	var month Month
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

//InitializeMonth ... will init the YAML file for the current month, based on the main config YAML
func InitializeMonth(user Ledger, desiredMonth string) (Month, string) {
	var month Month
	path := os.Getenv("HOME") + ledgerConfigDirName + "/" + desiredMonth + "/" + desiredMonth + ".yml"
	if IfFileExists(path) {
		month = UpdateExistingMonth(user, path)
	} else {
		month = CreateMonth(user)
	}
	return month, path
}

//IfFileExists ... checks file existance
func IfFileExists(configDirPath string) bool {
	if _, err := os.Stat(configDirPath); err == nil {
		return true
	}
	return false
}
