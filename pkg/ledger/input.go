package ledger

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Ledger struct that defines how the main config file should look
type Ledger struct {
	Admin   string     `yaml:"user"`
	Salary  int        `yaml:"salary"`
	Clients []Borrower `yaml:"borrowers"`
	Bills   []Bill     `yaml:"bills"`
	Places  []string   `yaml:"frequent_places"`
}

// GetInitialConf will process the main config YAML
func GetInitialConf(path string) {
	var admin Ledger
	yamlFile, err := ioutil.ReadFile(path)
	check(err)
	err = yaml.Unmarshal(yamlFile, &admin)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	InitializeCurrentMonth(admin)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
