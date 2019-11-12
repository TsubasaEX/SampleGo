// testkicker project testkicker.go
package testkicker

import (
	"app/edgesense/portal"
	"configutil"
	"fmt"
	"testutil"
)

func Kick(config configutil.Config) {

	var testfunc testutil.TestFunc
	for _, app := range config.Apps {
		if app.Enable {
			switch app.Name {
			case "es-edgesense-portal":
				testfunc = &portal.TestEntry{config.IP, app.Name, app.Label, app.Times, app.Report, app.Port}
				testfunc.Test()
			// case "es-edgesense-worker":
			default:
				fmt.Println("default")
			}
		}
	}

}
