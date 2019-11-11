// testkicker project testkicker.go
package testkicker

import (
	"app/edgesense/portal"
	"configutil"
	"fmt"
)

func Kick(config configutil.Config) {

	for _, app := range config.Apps {
		switch app.Name {
		case "es-edgesense-portal":
			portal.Test()
		// case "es-edgesense-worker":
		default:
			fmt.Println("default")
		}
	}

}
