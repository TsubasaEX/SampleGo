// portal project portal.go
package portal

import (
	"fmt"
	"html"
	"net/http"
	"os/exec"
	"strconv"
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

const ICON_CHECK = "\\U00002705"
const ICON_CROSS = "\\U0000274E"

func emoji(s string) string {
	// Hex String
	h := strings.ReplaceAll(s, "\\U", "0x")

	// Hex to Int
	i, _ := strconv.ParseInt(h, 0, 64)

	// Unescape the string (HTML Entity -> String).
	str := html.UnescapeString(string(i))
	return str
}

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
	fmt.Println(emoji(ICON_CHECK))
	fmt.Println(emoji(ICON_CROSS))
	rc.Println(TEST_AGAINST, testentry.Name, DOT_STR)
	for i := 0; i < testentry.Times; i++ {
		count := 0
		rc.Println(ROUND, i+1)
		fmt.Println(DASH_STR)
		for {
			fmt.Println(STAGE_ONE, DOT_STR)
			if testentry.status_check() {
				pc.Println(emoji(ICON_CHECK), PASS, STAGE_ONE)
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Println(STAGE_TWO, DOT_STR)
			if testentry.delete_pod() {
				pc.Println(emoji(ICON_CHECK), PASS, STAGE_TWO)
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Println(STAGE_THREE, DOT_STR)
			if testentry.status_check() {
				pc.Println(emoji(ICON_CHECK), PASS, STAGE_THREE)
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Println(STAGE_FOUR, DOT_STR)
			if testentry.func_check() {
				pc.Println(emoji(ICON_CHECK), PASS, STAGE_FOUR, "\n")
				break
			}
			count++
			if count >= RETRY_TIMES {
				fc.Println(emoji(ICON_CROSS), FAIL, STAGE_FOUR)
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}
	}
	pc.Println(emoji(ICON_CHECK), DONE, TEST_AGAINST, testentry.Name)
}
