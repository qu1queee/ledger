package ledger

import (
	"os"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/afero"
)

// TODO finish method comments
// TestMonthFileConfig
func TestMonthFilePresence(t *testing.T) {
	var ledger Ledger
	dir, _ := os.Getwd()
	os.Setenv("HOME", dir+"/../../test-assets/fake-home")

	Convey("When dealing with month files", t, func() {
		//Use afero wrapper around the native OS calls
		appfs := afero.NewOsFs()
		Convey("When file month does not exists and no month is provided, create it", func() {
			ledger = GetInitialConf("../../test-assets/config/config-test.yml", "admin", "")
			currentMonth := strings.ToLower(time.Now().Month().String())
			InitializeConfigFile(ledger, "../../test-assets/config/config-test.yml")
			InitializeLedgerCurrentMonthDir()
			monthStruct, monthFile := InitializeMonth(ledger, currentMonth)
			MarshallMonth(monthStruct, monthFile)
			So(monthStruct.User, ShouldEqual, "admin")
			err := appfs.RemoveAll(dir + "/../../test-assets/fake-home/.ledger/october")
			if err != nil {
				check(err)
			}
			So(len(monthStruct.Expenses), ShouldEqual, 3)

		})

		Convey("When file month does not exists and month is provided, create it", func() {
			ledger := GetInitialConf("../../test-assets/config/config-test-april.yml", "admin", "april")
			InitializeConfigFile(ledger, "../../test-assets/config/config-test-april.yml")
			InitializeLedgerCurrentMonthDir()
			monthStruct, monthFile := InitializeMonth(ledger, "april")
			MarshallMonth(monthStruct, monthFile)
			So(monthStruct.User, ShouldEqual, "adminofapril")
			err := appfs.Remove(dir + "/../../test-assets/fake-home/.ledger/april/april.yml")
			if err != nil {
				check(err)
			}
			So(len(monthStruct.Expenses), ShouldEqual, 3)
		})

	})

}

func TestMonthFileContents(t *testing.T) {
	var ledger Ledger
	dir, _ := os.Getwd()
	os.Setenv("HOME", dir+"/../../test-assets/fake-home/")

	//Use afero wrapper around the native OS calls
	appfs := afero.NewOsFs()

	Convey("For existing month file", t, func() {
		ledger = GetInitialConf("../../test-assets/config/config-another-test.yml", "admin", "may")
		InitializeConfigFile(ledger, "../../test-assets/config/config-another-test.yml")
		InitializeLedgerCurrentMonthDir()

		Convey("When config file is not changed", func() {
			monthStruct, monthFile := InitializeMonth(ledger, "may")
			MarshallMonth(monthStruct, monthFile)
			So(monthStruct.User, ShouldEqual, "owner")
			So(len(monthStruct.Expenses), ShouldEqual, 3)
		})

		Convey("When config file is updated", func() {
			ledger = GetInitialConf("../../test-assets/config/config-another-test-updated.yml", "admin", "may")
			monthStruct, _ := InitializeMonth(ledger, "may")
			So(len(monthStruct.Expenses), ShouldEqual, 5)
		})
		errFile := appfs.Remove(dir + "/../../test-assets/fake-home/.ledger/config.yaml")
		if errFile != nil {
			check(errFile)
		}
		errDir := appfs.RemoveAll(dir + "/../../test-assets/fake-home/.ledger/october")
		if errDir != nil {
			check(errDir)
		}
	})
}
