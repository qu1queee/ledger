package ledger

import (
	"fmt"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// Spend struct that defines a single spend action
type Spend struct {
	Amount      int    `yaml:"amount"`
	Date        string `yaml:"date"`
	Description string `yaml:"description"`
}

// Month struct that defines a month config YAML file
type Month struct {
	User     string             `yaml:"user"`
	Access   time.Time          `yaml:"access"`
	Bills    []Bill             `yaml:"bills"`
	Expenses map[string][]Spend `yaml:"frequent_places"`
}

// MarshallMonth func to generate the current month config YAML file
func MarshallMonth(month Month) {
	b, err := yaml.Marshal(month)
	check(err)
	fmt.Printf("--- t dump:\n%s\n\n", string(b))
	errs := ioutil.WriteFile("joder.yaml", b, 0644) // TODO: generate according to the current month
	check(errs)
}

// InitializeCurrentMonth will init the YAML file for the current month, based on the main config YAML
func InitializeCurrentMonth(user Ledger) {
	var mymap = make(map[string][]Spend)

	var spend = make([]Spend, 1)
	spend[0].Description = "Please provide a description (optional)"
	spend[0].Date = time.Now().Local().Format("2006-01-02") + " (optional)"

	var month Month
	month.Expenses = mymap
	month.Bills = user.Bills
	month.User = user.Admin

	for _, bills := range user.Places {
		mymap[bills] = spend
	}
	month.Expenses = mymap
	month.Access = time.Now().Local()
	MarshallMonth(month) //TODO: avoid this if file exist
}
