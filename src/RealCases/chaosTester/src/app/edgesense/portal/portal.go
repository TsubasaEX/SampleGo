// portal project portal.go
package portal

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type TestEntry struct {
	IP     string
	Label  string
	Times  int
	Report bool
}

const NO_RESOURCES_FOUND string = "No resources found."
const PASS string = "[PASS]"

func (testentry *TestEntry) status_check() bool {
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

func (testentry *TestEntry) delete_pod() bool {
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

func (testentry *TestEntry) func_check() bool {
	return true
}

func (testentry *TestEntry) Test() {
	for {
		fmt.Println("Stage 1.....")
		if testentry.status_check() {
			fmt.Println(PASS, "Stage 1")
			break
		}
		time.Sleep(10000 * time.Millisecond) //10 sec per status check
	}

	for {
		fmt.Println("Stage 2......")
		if testentry.delete_pod() {
			fmt.Println(PASS, "Stage 2")
			break
		}
		time.Sleep(10000 * time.Millisecond) //10 sec per status check
	}

	for {
		fmt.Println("Stage 3......")
		if testentry.status_check() {
			fmt.Println(PASS, "Stage 3")
			break
		}
		time.Sleep(10000 * time.Millisecond) //10 sec per status check
	}
	fmt.Println("Done")
}
