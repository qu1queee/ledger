package ledger

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

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

// InitializeLedgerCurrentMonthDir generate new .ledger/month dir under HOME path
func InitializeLedgerCurrentMonthDir() {
	var currentTime time.Month
	_, currentTime, _ = time.Now().UTC().Date()
	currentTimeLowerCase := strings.ToLower(currentTime.String())
	if _, err := os.Stat(os.Getenv("HOME") + ledgerConfigDirName + "/" + currentTimeLowerCase); os.IsNotExist(err) {
		err = os.MkdirAll(os.Getenv("HOME")+ledgerConfigDirName+"/"+currentTimeLowerCase, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// InitializeLedgerRootDir generate new .ledger dir under HOME path
func InitializeLedgerRootDir() {
	if _, err := os.Stat(os.Getenv("HOME") + ledgerConfigDirName); os.IsNotExist(err) {
		err = os.MkdirAll(os.Getenv("HOME")+ledgerConfigDirName, 0755)
		if err != nil {
			panic(err)
		}
	}
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
	InitializeLedgerRootDir()
	InitializeLedgerCurrentMonthDir()
	InitializeCurrentMonth(admin)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
