package ledger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

// TestRootConfigPath will verify if the root
// config directory exists
func TestRootConfigPath(t *testing.T) {
	Convey("Simply by calling the ledger binary", t, func() {
		Convey("When invoking the binary", func() {
			var flag bool
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			os.Setenv("HOME", dir+"/../../test-assets/fake-home")
			fmt.Println(os.Getenv("HOME"))
			InitializeLedgerRootDir()
			if _, err := os.Stat(os.Getenv("HOME") + "/.ledger"); err == nil {
				flag = true
			}
			So(flag, ShouldBeTrue)
		})
	})
}

// TestConfigFile will verify if the configuration
// file used by ledger exists.
func TestConfigFile(t *testing.T) {
	// About the configuration file
	Convey("Given some configuration file", t, func() {
		// If the path to configuration file is provided
		Convey("When the configuration file is provided", func() {
			var ledger Ledger
			ledger = GetInitialConf("../../test-assets/config/config-test.yml", "admin", "")
			marshalledLedger, err := json.Marshal(ledger)
			if err != nil {
				fmt.Println(err)
				return
			}
			expected := `{"Admin":"admin","Salary":5000,"Clients":null,"Bills":[{"Company":"a_company","Type":"monthly","Amount":100,"Description":"some description"}],"Places":["shop_a","shop_b","shop_c"]}`
			So(string(marshalledLedger), ShouldEqual, expected)
		})
		// If the path to configuration file is provided,
		// but there is no copy generated under the ROOT ledger dir.
		// Also, HOME env var, needs to be modify, to enable test-assets folder.
		Convey("When the configuration file is provided, and it does not exist in ledger root", func() {
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			os.Setenv("HOME", dir+"/../../test-assets")
			InitializeConfigFile(GetInitialConf("../../test-assets/config/config-test.yml", "admin", ""), "../../test-assets/config/config-test.yml")
			flag := IfConfigFileExist(dir + "/../../test-assets/.ledger/config.yaml")
			So(flag, ShouldBeTrue)
		})

		// If the path to configuration file is not provided,
		// but there is a copy generated under the ROOT ledger dir
		Convey("When the configuration file is not provided, and already exist in ledger root", func() {
			var ledger Ledger
			ledger = GetInitialConf("", "admin", "")
			So(ledger.Admin, ShouldHaveSameTypeAs, "idleuser")
		})
	})
}

// TestMonthConfig will verify proper creation of
// the month dir under .ledger
func TestMonthConfig(t *testing.T) {
	var ledger Ledger
	dir, _ := os.Getwd()
	os.Setenv("HOME", dir+"/../../test-assets/")
	ledger = GetInitialConf("../../test-assets/config/config-test.yml", "admin", "")
	InitializeConfigFile(ledger, "../../test-assets/config/config-test.yml")
	_, month, _ := time.Now().Date()

	Convey("When dealing with month records", t, func() {

		Convey("When no month is provided and no month dir exists, create the month config", func() {
			InitializeLedgerCurrentMonthDir()
			flag := IfDirFileExists(dir + "/../../test-assets/.ledger/" + strings.ToLower(month.String()))
			So(flag, ShouldBeTrue)
		})

		Convey("When no month is provided and the month dir exists, get the current month", func() {
			os.Setenv("HOME", dir+"/../../test-assets/fake-home")
			InitializeLedgerCurrentMonthDir()
			flag := IfDirFileExists(dir + "/../../test-assets/fake-home/.ledger/april")
			So(flag, ShouldBeTrue)
		})
	})
}

func IfDirFileExists(configDirPath string) bool {
	if _, err := os.Stat(configDirPath); err == nil {
		return true
	}
	return false
}
