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

// HOME defines home path of the environment
var HOME = os.Getenv("HOME")

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
	if _, err := os.Stat(HOME + ledgerConfigDirName + "/" + currentTimeLowerCase); os.IsNotExist(err) {
		err = os.MkdirAll(HOME+ledgerConfigDirName+"/"+currentTimeLowerCase, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// InitializeLedgerRootDir generate new .ledger dir under HOME path
func InitializeLedgerRootDir() {
	if _, err := os.Stat(HOME + ledgerConfigDirName); os.IsNotExist(err) {
		err = os.MkdirAll(HOME+ledgerConfigDirName, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// IfConfigFileExist checks that a config.yaml exists under Ledger HOME path
func IfConfigFileExist(configFilePath string) (bool, string) {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return false, configFilePath
	}
	return true, configFilePath
}

// InitializeConfigFile creates a config.yaml if not exists
func InitializeConfigFile(user Ledger, configFile string) {
	b, err := yaml.Marshal(user)
	check(err)
	configExist, configPath := IfConfigFileExist(configFile)
	if !configExist {
		fmt.Printf("Adding %v into ~/.ledger\n", prettyRedBold("config.yaml"))
		errs := ioutil.WriteFile(configPath, b, 0644)
		check(errs)
	}
}

// readConfigFile will unmarshall the config.yaml
func readConfigFile(configPath string) Ledger {
	var ledger Ledger
	yamlFile, err := ioutil.ReadFile(configPath)
	check(err)
	err = yaml.Unmarshal(yamlFile, &ledger)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return ledger
}

// GetInitialConf will process the main config YAML
func GetInitialConf(path string, user string, desiredMonth string) {
	var ledger Ledger
	var configFilePath = HOME + ledgerConfigDirName + "/" + configFilePath
	configExist, configPath := IfConfigFileExist(configFilePath)
	if configExist && path == "" {
		fmt.Printf("Going to use %v from ~/.ledger\n", prettyRedBold("config.yaml"))
		ledger = readConfigFile(configPath)
		if ledger.Admin != user {
			fmt.Printf("User %v, does not exists\n", prettyRedBold(user))
			os.Exit(127)
		}
	} else if path != "" {
		fmt.Printf("Going to use %v \n", prettyRedBold(path))
		ledger = readConfigFile(path)
		InitializeConfigFile(ledger, configFilePath)
		InitializeLedgerRootDir()
		InitializeLedgerCurrentMonthDir()
	} else {
		fmt.Println("No config.yaml found, use --config to provide it")
		os.Exit(127)
	}
	InitializeCurrentMonth(ledger, desiredMonth)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
