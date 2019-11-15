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
const STAGE_ONE string = "Stage 1: Checking Running Status"
const STAGE_TWO string = "Stage 2: Deleting Pod"
const STAGE_THREE string = "Stage 3: Checking Running Status"
const STAGE_FOUR string = "Stage 4: Checking Functions"
const PASS string = "[PASS]"
const FAIL string = "[FAIL]"
const COMPLETE string = "[COMPLETE]"
const DOT_STR = "............"
const DASH_STR = "----------------------------------------"
const RETRY_TIMES = 3

const ICON_CHECK = "\\U00002705"
const ICON_CROSS = "\\U0000274E"
const ICON_AIRPLANE = "\\U00002708"
const ICON_SPARKLE = "\\U00002728"
const ICON_HEART = "\\U00002764"
const ICON_PENCIL = "\\U0000270F"
const ICON_SCISSOR = "\\U00002702"
const ICON_ARROW = "\\U000027A1"

var simple = false

func emoji(s string) string {
	// Hex String
	h := strings.ReplaceAll(s, "\\U", "0x")

	// Hex to Int
	i, _ := strconv.ParseInt(h, 0, 64)

	// Unescape the string (HTML Entity -> String).
	str := html.UnescapeString(string(i))
	return str
}

func getTestString(s string) string {
	return fmt.Sprintln(emoji(ICON_SCISSOR), TEST_AGAINST, s, DOT_STR, emoji(ICON_HEART))
}

func getSplitLine() string {
	str := ""
	for i := 0; i < 20; i++ {
		str += emoji(ICON_ARROW)
	}
	return fmt.Sprintln(str)
}
func getRoundString(n int) string {
	return fmt.Sprintln(emoji(ICON_AIRPLANE), ROUND, n)
}

func getStageString(s string) string {
	return fmt.Sprintln(emoji(ICON_SPARKLE), s, DOT_STR, emoji(ICON_HEART))
}

func getStagePassString(s string) string {
	return fmt.Sprintln(emoji(ICON_CHECK), PASS, s)
}

func getStageFailString(s string) string {
	return fmt.Sprintln(emoji(ICON_CROSS), FAIL, s)
}

func getCompleteString(s string) string {
	return fmt.Sprintln(emoji(ICON_CHECK), COMPLETE, TEST_AGAINST, s)
}

func (testentry *TestEntry) status_check() bool {

	str := "kubectl get pods -l " + testentry.Label + " --field-selector=status.phase=Running"
	cmd := exec.Command("cmd", "/K", str)

	if !simple {
		fmt.Println(emoji(ICON_PENCIL), str)
	}
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

	if !simple {
		fmt.Println(emoji(ICON_PENCIL), str)
	}
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

	if !simple {
		fmt.Println(emoji(ICON_PENCIL), str)
	}
	_, err := http.Get(str)

	if err != nil {
		return false
		// fmt.Println("1")
		// log.Fatal(err.Error())
	}
	return true
}

func (testentry *TestEntry) Test(args []string) {
	rc := color.New(color.FgCyan, color.Bold)
	pc := color.New(color.FgGreen, color.Bold)
	fc := color.New(color.FgRed, color.Bold)
	cc := color.New(color.FgYellow, color.Bold)

	if len(args) != 0 {
		if strings.Compare(args[0], "-s") == 0 {
			simple = true
		}
	}

	rc.Print(getTestString(testentry.Name))
	for i := 0; i < testentry.Times; i++ {
		count := 0
		fmt.Print(getSplitLine())
		rc.Print(getRoundString(i + 1))
		for {
			fmt.Print(getStageString(STAGE_ONE))
			if testentry.status_check() {
				pc.Print(getStagePassString(STAGE_ONE))
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Print(getStageString(STAGE_TWO))
			if testentry.delete_pod() {
				pc.Print(getStagePassString(STAGE_TWO))
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Print(getStageString(STAGE_THREE))
			if testentry.status_check() {
				pc.Print(getStagePassString(STAGE_THREE))
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Print(getStageString(STAGE_FOUR))
			if testentry.func_check() {
				pc.Print(getStagePassString(STAGE_FOUR))
				break
			}
			count++
			if count >= RETRY_TIMES {
				fc.Print(getStageFailString(STAGE_FOUR))
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}
	}
	cc.Print(getCompleteString(testentry.Name))
}
