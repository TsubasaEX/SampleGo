// portal project portal.go
package portal

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

type TestEntry struct {
	IP     string
	Name   string
	Label  string
	Times  int
	Report bool
	Port   string
}

const NO_RESOURCES_FOUND = "No resources found."
const TEST_AGAINST string = "Tests Against"
const ROUND string = "Round"
const STAGE_ONE string = "Stage 1"
const STAGE_TWO string = "Stage 2"
const STAGE_THREE string = "Stage 3"
const STAGE_FOUR string = "Stage 4"
const PASS string = "[PASS]"
const FAIL string = "[FAIL]"
const DONE string = "[DONE]"
const DOT_STR = "............"
const DASH_STR = "----------------------------------------"
const RETRY_TIMES = 3

func (testentry *TestEntry) status_check() bool {

	str := "kubectl get pods -l " + testentry.Label + " --field-selector=status.phase=Running"
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
	str := "kubectl delete pods -l " + testentry.Label + " --field-selector=status.phase=Running"
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
	str := "http://" + testentry.IP + ":" + testentry.Port
	fmt.Println(str)

	_, err := http.Get(str)

	if err != nil {
		return false
		// fmt.Println("1")
		// log.Fatal(err.Error())
	}
	return true
}

func (testentry *TestEntry) Test() {
	rc := color.New(color.FgCyan, color.Bold)
	pc := color.New(color.FgGreen, color.Bold)
	fc := color.New(color.FgRed, color.Bold)

	rc.Println(TEST_AGAINST, testentry.Name, DOT_STR)
	for i := 0; i < testentry.Times; i++ {
		count := 0
		rc.Println(ROUND, i+1)
		fmt.Println(DASH_STR)
		for {
			fmt.Println(STAGE_ONE, DOT_STR)
			if testentry.status_check() {
				pc.Println(PASS, STAGE_ONE)
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Println(STAGE_TWO, DOT_STR)
			if testentry.delete_pod() {
				pc.Println(PASS, STAGE_TWO)
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Println(STAGE_THREE, DOT_STR)
			if testentry.status_check() {
				pc.Println(PASS, STAGE_THREE)
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Println(STAGE_FOUR, DOT_STR)
			if testentry.func_check() {
				pc.Println(PASS, STAGE_FOUR, "\n")
				break
			}
			count++
			if count >= RETRY_TIMES {
				fc.Println(FAIL, STAGE_FOUR)
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}
	}
	pc.Println(DONE, TEST_AGAINST, testentry.Name)
}
