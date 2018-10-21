package ledger

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDataContents(t *testing.T) {
	Convey("When month files exist", t, func() {
		// define something all my conveys need
		dir, _ := os.Getwd()
		os.Setenv("HOME", dir+"/../../test-assets/fake-home")
		monthName := "complex-month"
		user := "complex-admin"
		configPath := dir + "/../../test-assets/config/config-complex.yml"

		InitializeLedgerRootDir()
		ledgerStruct := GetInitialConf(configPath, user, monthName)
		monthStruct, monthFile := InitializeMonth(ledgerStruct, monthName)
		MarshallMonth(monthStruct, monthFile)

		Convey("When json format require", func() {
			jsonData := GetJson(monthStruct)
			jsonFake := `{"User":"complex-admin","Bills":[{"Company":"a_company","Type":"monthly","Amount":100,"Description":"some description"}],"Expenses":{"shop_a":[{"Amount":17,"Date":"01.10.2018","Description":"cheap shop_a product"},{"Amount":21.5,"Date":"01.11.2018","Description":"expensive shop_a product"}],"shop_b":[{"Amount":0,"Date":"","Description":"Please provide a description (optional)"}],"shop_c":[{"Amount":0,"Date":"","Description":"Please provide a description (optional)"}],"shop_d":[{"Amount":0,"Date":"","Description":"Please provide a description (optional)"}],"shop_r":[{"Amount":0,"Date":"","Description":""}]}}`
			So(jsonFake, ShouldEqual, jsonData)
		})
	})
}
