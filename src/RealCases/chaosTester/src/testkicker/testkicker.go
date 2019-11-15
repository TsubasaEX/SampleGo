// testkicker project testkicker.go
package testkicker

import (
	"app/edgesense/portal"
	"configutil"
	"fmt"
	"testutil"
)

func Kick(config configutil.Config, args []string) {

	var testfunc testutil.TestFunc

	for _, web := range config.Apps.Web {
		if web.Enable {
			switch web.Name {
			case "es-edgesense-portal":
				testfunc = &portal.TestEntry{config.IP, web.Name, web.Label, web.Times, web.Report, web.Port}
				testfunc.Test(args)
			// case "es-edgesense-worker":
			default:
				fmt.Println("default")
			}
		}
	}

	// for _, app := range config.Apps.App {
	// }

}
