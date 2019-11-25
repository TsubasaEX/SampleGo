// testkicker project testkicker.go
package testkicker

import (
	"app/edgesense/portal"
	"configutil"
	"encoding/csv"
	"fmt"
	"log"
	"strings"
	"testutil"
)

func Kick(config configutil.Config, args []string) string {

	var testfunc testutil.TestFunc
	var testLogger *log.Logger
	var testReporter *csv.Writer
	var teststReporter_W *csv.Writer
	var b_Simple bool = false
	var b_Log bool = false
	var records [][]string
	var n = 1
	if len(args) != 0 {
		split := strings.Split(args[0], "-")
		if strings.Contains(split[1], "s") {
			b_Simple = true
		}
		if strings.Contains(split[1], "l") {
			b_Log = true
		}
	}
	if b_Log {
		testLogger = testutil.GetLogger()
	}
	testReporter = testutil.GetReporter()
	teststReporter_W = testutil.GetStatisticsReporter_W()
	for _, web := range config.Apps.Web {
		if web.Enable {
			switch web.Name {
			case "es-edgesense-portal":
				testfunc = &portal.TestEntity{config.IP, web.Name, web.Label, web.Times, web.Port}
				passNum := testfunc.Test(args, b_Simple, testLogger, testReporter)
				record := testutil.GetStatisticsRecord(n, web.Name, passNum, web.Times)
				records = append(records, record)
			// case "es-edgesense-worker":
			default:
				fmt.Println(testutil.GetNoTestCaseString(web.Name))
				if testLogger != nil {
					testLogger.Println(testutil.GetNoTestCaseString(web.Name))
				}
			}
		}
	}
	teststReporter_W.WriteAll(records)
	if err := teststReporter_W.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	if b_Log {
		testutil.CloseLogger()
	}
	testutil.CloseReporter()
	testutil.CloseStatisticsReporter_W()

	return testutil.GetStatisticsFileName()
	// for _, app := range config.Apps.App {
	// }

}
