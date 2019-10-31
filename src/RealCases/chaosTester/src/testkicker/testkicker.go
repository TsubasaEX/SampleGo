// testkicker project testkicker.go
package testkicker

import (
	"configutil"
	"fmt"
	"os/exec"
)

func Kick(config configutil.Config) {

	fmt.Println(config)
	fmt.Println(len(config.Apps))
	// for i := 0; i < 2; i++ {
	// 	fmt.Println(config.Apps[i])
	// }
	for _, app := range config.Apps {
		fmt.Println(app.Name)
	}
	//kubectl get pods -l rmmModule=rmm-portal --field-selector=status.phase=Running
	str := "kubectl get pods -l rmmModule=rmm-portal --field-selector=status.phase=Running"
	cmd := exec.Command("cmd", "/K", str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to exec command:", cmd)
	}
	fmt.Println(string(out))
}
