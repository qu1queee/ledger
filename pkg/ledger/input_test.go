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
	"github.com/spf13/afero"
)

// TestRootConfigPath will verify if the root
// config directory exists
func TestRootConfigPath(t *testing.T) {
	Convey("Simply by calling the ledger binary", t, func() {
		//Use afero wrapper around the native OS calls
		appfs := afero.NewOsFs()
		Convey("When invoking the binary", func() {
			var flag bool
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			os.Setenv("HOME", dir+"/../../test-assets")
			InitializeLedgerRootDir()
			if _, err1 := os.Stat(os.Getenv("HOME") + "/.ledger"); err1 == nil {
				flag = true
			}
			So(flag, ShouldBeTrue)
			os.Unsetenv("HOME")
			err = appfs.RemoveAll("../../test-assets/.ledger")
			if os.IsNotExist(err) {
				check(err)
			}
		})
	})
}

// TestConfigFile will verify if the configuration
// file used by ledger exists.
func TestConfigFile(t *testing.T) {
	// About the configuration file
	Convey("Given some configuration file", t, func() {
		dir, errRootedPath := os.Getwd()
		if errRootedPath != nil {
			log.Fatal(errRootedPath)
		}
		//This simulates the binary call to ledger, by passing a
		//configFile and the desired user
		ledger := GetInitialConf("../../test-assets/config/config-test.yml", "admin", "")

		//Use afero memory filesystem mock to manage filesystem calls
		appFS := afero.NewMemMapFs()
		//Use afero wrapper around the native OS calls
		appfs := afero.NewOsFs()

		// If the path to configuration file is provided
		// make sure the config can be extracted
		Convey("When the configuration file is provided", func() {
			marshalledLedger, marshallError := json.Marshal(ledger)
			if marshallError != nil {
				fmt.Println(marshallError)
				return
			}
			expected := `{"Admin":"admin","Salary":5000,"Clients":null,"Bills":[{"Company":"a_company","Type":"monthly","Amount":100,"Description":"some description"}],"Places":["shop_a","shop_b","shop_c"]}`
			So(string(marshalledLedger), ShouldEqual, expected)
		})
		// If the path to configuration file is provided,
		// but there is no copy generated under the ROOT ledger dir.
		// Also, HOME env var, needs to be modify, to enable test-assets folder.
		Convey("When the configuration file is provided, and it does not exist in ledger root", func() {
			os.Setenv("HOME", dir+"/../../test-assets")
			InitializeConfigFile(ledger, "../../test-assets/config/config-test.yml")

			//Read provisioned file via afero, and save it under
			//the in memory mocked filesystem
			defaultConfig, aferoReadConfigError := afero.ReadFile(appfs, "../../test-assets/.ledger/config.yaml")
			if os.IsNotExist(aferoReadConfigError) {
				check(aferoReadConfigError)
			}
			afero.WriteFile(appFS, "../../test-assets/.ledger/config.yaml", defaultConfig, 0644)

			// Verify that the config file exist in memory backed filesystem
			_, err := appFS.Stat("../../test-assets/.ledger/config.yaml")
			if os.IsNotExist(err) {
				check(err)
			}
			aferoFile, aferoReadMainConfigError := afero.ReadFile(appFS, "../../test-assets/.ledger/config.yaml")
			if os.IsNotExist(aferoReadMainConfigError) {
				check(aferoReadMainConfigError)
			}
			//Compare content between existing provisioned file
			//and the in memory one generated with afero
			So(string(defaultConfig), ShouldEqual, string(aferoFile))
			os.Unsetenv("HOME")
		})

		// If the path to configuration file is not provided,
		// but there is a copy generated under the ROOT ledger dir
		Convey("When the configuration file is not provided, and already exist in ledger root", func() {
			os.Setenv("HOME", dir+"/../../test-assets")
			//This simulates the binary call to ledger, by ommitting the
			//configFile and the desired user
			ledger := GetInitialConf("", "admin", "")
			var fakeLedger Ledger
			//Extract existing in memory file, and use it.
			bytes, _ := afero.ReadFile(appFS, "../../test-assets/.ledger/config.yml")
			json.Unmarshal(bytes, &fakeLedger)
			So(fakeLedger.Admin, ShouldHaveSameTypeAs, ledger.Admin)
			os.Unsetenv("HOME")
			//Todo: find a way to mock this(input.go itself)
			err := appfs.RemoveAll("../../test-assets/.ledger")
			if os.IsNotExist(err) {
				check(err)
			}
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

	//Use afero wrapper around the native OS calls
	appfs := afero.NewOsFs()

	Convey("When dealing with month records", t, func() {

		Convey("When no month is provided and no month dir exists, create the month config", func() {
			InitializeLedgerCurrentMonthDir()
			monthDirExists, err := afero.DirExists(appfs, dir+"/../../test-assets/.ledger/"+strings.ToLower(month.String()))
			if err != nil {
				log.Fatal(err)
			}
			So(monthDirExists, ShouldBeTrue)
			err = appfs.RemoveAll(dir + "/../../test-assets/.ledger")
			if os.IsNotExist(err) {
				check(err)
			}
			os.Unsetenv("HOME")
		})
	})
}
