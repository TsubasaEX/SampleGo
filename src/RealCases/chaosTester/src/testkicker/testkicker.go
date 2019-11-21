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

func Kick(config configutil.Config, args []string) {

	var testfunc testutil.TestFunc
	var testLogger *log.Logger
	var testReporter *csv.Writer
	var b_Simple bool = false
	var b_Log bool = false
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

	for _, web := range config.Apps.Web {
		if web.Enable {
			switch web.Name {
			case "es-edgesense-portal":
				testfunc = &portal.TestEntity{config.IP, web.Name, web.Label, web.Times, web.Port}
				testfunc.Test(args, b_Simple, testLogger, testReporter)
			// case "es-edgesense-worker":
			default:
				fmt.Println("default")
			}
		}
	}

	if b_Log {
		testutil.CloseLogger()
	}
	testutil.CloseReporter()

	// for _, app := range config.Apps.App {
	// }

}
