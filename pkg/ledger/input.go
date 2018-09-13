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

const configFilePath = "config.yaml"

// LedgerDir sets root ledger repo dir
const LedgerDir = ".ledger"

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
	if _, err := os.Stat(os.Getenv("HOME") + "/" + LedgerDir); os.IsNotExist(err) {
		err = os.MkdirAll(os.Getenv("HOME")+"/"+LedgerDir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// IfConfigFileExist checks that a config.yaml exists under the root DIR ~/.ledger
func IfConfigFileExist(configFilePath string) bool {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return false
	}
	return true
}

// InitializeConfigFile creates a config.yaml if not exists
func InitializeConfigFile(user Ledger, configFile string) {
	b, err := yaml.Marshal(user)
	check(err)
	if !IfConfigFileExist(os.Getenv("HOME") + "/" + LedgerDir + "/" + configFilePath) {
		// TODO: Add debug mode
		// fmt.Printf("Adding %v into ~/.ledger\n", prettyRedBold("config.yaml"))
		errs := ioutil.WriteFile(os.Getenv("HOME")+"/"+LedgerDir+"/"+configFilePath, b, 0644)
		check(errs)
	}
}

// GetInitialConf will process the main config YAML and generated a copy
// under the ledger ROOT dir
func GetInitialConf(path string, user string, desiredMonth string) Ledger {
	if len(path) == 0 {
		path = os.Getenv("HOME") + "/" + LedgerDir + "/" + configFilePath
	}
	configExist := IfConfigFileExist(path)
	if !configExist {
		fmt.Println("No config.yaml found, use --config to provide it")
		os.Exit(1)
	}
	var ledger Ledger
	yamlFile, err := ioutil.ReadFile(path)
	check(err)
	err = yaml.Unmarshal(yamlFile, &ledger)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return ledger
}

// SetupUserFile ss
func SetupUserFile() {
	// InitializeLedgerCurrentMonthDir()
	// InitializeCurrentMonth(ledger, desiredMonth)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
