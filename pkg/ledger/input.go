package ledger

import (
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// Ledger struct that defines how the main config file should look
type Ledger struct {
	Admin   string `yaml:"user"`
	Salary  int    `yaml:"salary"`
	Access  time.Time
	Clients []Borrower `yaml:"borrowers"`
	Bills   []Bill     `yaml:"bills"`
	Places  []string   `yaml:"frequent_places"`
}

// GetInitialConf will process the main config YAML
func GetInitialConf(path string) Ledger {
	var admin Ledger
	yamlFile, err := ioutil.ReadFile(path)
	check(err)
	err = yaml.Unmarshal(yamlFile, &admin)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return admin
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
