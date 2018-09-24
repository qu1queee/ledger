package ledger

import (
	"os"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

// TODO finish method comments
// TestMonthFileConfig
func TestMonthFilePresence(t *testing.T) {
	var ledger Ledger
	dir, _ := os.Getwd()
	os.Setenv("HOME", dir+"/../../test-assets/fake-home/")

	Convey("When dealing with month files", t, func() {

		Convey("When file month does not exists and no month is provided, create it", func() {
			ledger = GetInitialConf("../../test-assets/config/config-test.yml", "admin", "")
			currentMonth := strings.ToLower(time.Now().Month().String())
			InitializeConfigFile(ledger, "../../test-assets/config/config-test.yml")
			InitializeLedgerCurrentMonthDir()
			monthStruct, _ := InitializeMonth(ledger, currentMonth)
			So(monthStruct.User, ShouldEqual, "admin")
			So(len(monthStruct.Expenses), ShouldEqual, 3)
		})

		Convey("When file month does not exists and month is provided, create it", func() {
			ledger2 := GetInitialConf("../../test-assets/config/config-test-april.yml", "admin", "april")
			InitializeConfigFile(ledger2, "../../test-assets/config/config-test-april.yml")
			InitializeLedgerCurrentMonthDir()
			monthStruct, _ := InitializeMonth(ledger2, "april")
			So(monthStruct.User, ShouldEqual, "adminofapril")
			So(len(monthStruct.Expenses), ShouldEqual, 3)
		})

	})

}

func TestMonthFileContents(t *testing.T) {
	var ledger Ledger
	dir, _ := os.Getwd()
	os.Setenv("HOME", dir+"/../../test-assets/fake-home/")

	Convey("For existing month file", t, func() {
		ledger = GetInitialConf("../../test-assets/config/config-another-test.yml", "admin", "may")
		InitializeConfigFile(ledger, "../../test-assets/config/config-another-test.yml")
		InitializeLedgerCurrentMonthDir()
		monthStruct, _ := InitializeMonth(ledger, "may")

		Convey("When config file is not changed", func() {
			So(monthStruct.User, ShouldEqual, "owner")
			So(len(monthStruct.Expenses), ShouldEqual, 3)
		})

		Convey("When config file is updated", func() {
			ledger = GetInitialConf("../../test-assets/config/config-another-test-updated.yml", "admin", "may")
			monthStruct, _ := InitializeMonth(ledger, "may")
			So(len(monthStruct.Expenses), ShouldEqual, 5)
		})

	})
}
