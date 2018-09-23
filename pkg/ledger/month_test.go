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
func TestMonthFileConfig(t *testing.T) {
	var ledger Ledger
	dir, _ := os.Getwd()
	os.Setenv("HOME", dir+"/../../test-assets/fake-home/")
	ledger = GetInitialConf("../../test-assets/config/config-test.yml", "admin", "")
	InitializeConfigFile(ledger, "../../test-assets/config/config-test.yml")
	_, month, _ := time.Now().Date()

	Convey("When dealing with month files, and no month is provided", t, func() {

		Convey("When file month does not exists, create it", func() {
			InitializeLedgerCurrentMonthDir()
			InitializeCurrentMonth(ledger, strings.ToLower(month.String()))
			flag := IfFileExists(dir + "/../../test-assets/fake-home/.ledger/" + strings.ToLower(month.String()) + "/" + strings.ToLower(month.String()) + ".yml")
			So(flag, ShouldBeTrue)
		})

		// TODO: add stdout match here
		// Convey("When file month exists, do nothing", func() {
		// 	InitializeCurrentMonth(ledger, "april")
		// 	flag := IfFileExists(dir + "/../../test-assets/fake-home/.ledger/" + "april/april.yml")
		// 	So(flag, ShouldBeTrue)
		// })
	})
}

func IfFileExists(configDirPath string) bool {
	if _, err := os.Stat(configDirPath); err == nil {
		return true
	}
	return false
}
