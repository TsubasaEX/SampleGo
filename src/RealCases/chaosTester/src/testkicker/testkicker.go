// testkicker project testkicker.go
package testkicker

import (
	"app/edgesense/portal"
	"configutil"
	"fmt"
	"log"
	"strings"
	"testutil"
)

func Kick(config configutil.Config, args []string) {

	var testfunc testutil.TestFunc
	var testLogger *log.Logger
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

	for _, web := range config.Apps.Web {
		if web.Enable {
			switch web.Name {
			case "es-edgesense-portal":
				testfunc = &portal.TestEntry{config.IP, web.Name, web.Label, web.Times, web.Report, web.Port}
				testfunc.Test(args, b_Simple, testLogger)
			// case "es-edgesense-worker":
			default:
				fmt.Println("default")
			}
		}
	}

	if b_Log {
		testutil.CloseLogger()
	}

	// for _, app := range config.Apps.App {
	// }

}
