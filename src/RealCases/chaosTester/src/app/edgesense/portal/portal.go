// portal project portal.go
package portal

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const NO_RESOURCES_FOUND string = "No resources found."
const PASS string = "[PASS]"

func Pr() string {
	return "portal"
}

func status_check() bool {
	str := "kubectl get pods -l rmmModule=rmm-portal --field-selector=status.phase=Running"
	cmd := exec.Command("cmd", "/K", str)

	fmt.Println(str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to exec command:", cmd)
	}

	str_out := string(out)
	if strings.Contains(str_out, NO_RESOURCES_FOUND) {
		return false
	}
	return true
}

func delete_pod() bool {
	str := "kubectl delete pods -l rmmModule=rmm-portal --field-selector=status.phase=Running"
	cmd := exec.Command("cmd", "/K", str)

	fmt.Println(str)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to exec command:", cmd)
	}

	str_out := string(out)
	if strings.Contains(str_out, NO_RESOURCES_FOUND) {
		return false
	}
	return true
}

func func_check() bool {
	return true
}

func Test() {
	for {
		fmt.Println("Stage 1.....")
		if status_check() {
			fmt.Println(PASS, "Stage 1")
			break
		}
		time.Sleep(10000 * time.Millisecond) //10 sec per status check
	}

	for {
		fmt.Println("Stage 2......")
		if delete_pod() {
			fmt.Println(PASS, "Stage 2")
			break
		}
		time.Sleep(10000 * time.Millisecond) //10 sec per status check
	}

	for {
		fmt.Println("Stage 3......")
		if status_check() {
			fmt.Println(PASS, "Stage 3")
			break
		}
		time.Sleep(10000 * time.Millisecond) //10 sec per status check
	}
	fmt.Println("Done")
}
