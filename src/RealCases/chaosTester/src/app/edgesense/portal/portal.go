// portal project portal.go
package portal

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"testutil"
	"time"
)

type TestEntry struct {
	IP     string
	Name   string
	Label  string
	Times  int
	Report bool
	Port   string
}

var b_Simple = false
var testLogger *log.Logger

func (testentry *TestEntry) status_check() bool {
	str := "kubectl get pods -l " + testentry.Label + " --field-selector=status.phase=Running"
	cmd := exec.Command("cmd", "/K", str)

	if !b_Simple {
		fmt.Println(testutil.Emoji(testutil.ICON_PENCIL), str)
		if testLogger != nil {
			testLogger.Println(testutil.Emoji(testutil.ICON_PENCIL), str)
		}
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Failed to exec command:", cmd)
		if testLogger != nil {
			testLogger.Println(err.Error())
			testLogger.Println("Failed to exec command:", cmd)
		}
	}

	str_out := string(out)
	if strings.Contains(str_out, testutil.NO_RESOURCES_FOUND) {
		return false
	}
	return true
}

func (testentry *TestEntry) delete_pod() bool {
	str := "kubectl delete pods -l " + testentry.Label + " --field-selector=status.phase=Running"
	cmd := exec.Command("cmd", "/K", str)

	if !b_Simple {
		fmt.Println(testutil.Emoji(testutil.ICON_PENCIL), str)
		if testLogger != nil {
			testLogger.Println(testutil.Emoji(testutil.ICON_PENCIL), str)
		}
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Failed to exec command:", cmd)
	}

	str_out := string(out)
	if strings.Contains(str_out, testutil.NO_RESOURCES_FOUND) {
		return false
	}
	return true
}

func (testentry *TestEntry) func_check() bool {
	str := "http://" + testentry.IP + ":" + testentry.Port

	if !b_Simple {
		fmt.Println(testutil.Emoji(testutil.ICON_PENCIL), str)
		if testLogger != nil {
			testLogger.Println(testutil.Emoji(testutil.ICON_PENCIL), str)
		}
	}
	_, err := http.Get(str)

	if err != nil {
		return false
	}
	return true
}

func (testentry *TestEntry) Test(args []string, simple bool, testlogger *log.Logger) {
	b_Simple = simple
	testLogger = testlogger

	testutil.Rc.Print(testutil.GetTestString(testentry.Name))
	if testLogger != nil {
		testLogger.Print(testutil.GetTestString(testentry.Name))
	}
	for i := 0; i < testentry.Times; i++ {
		count := 0
		fmt.Print(testutil.GetSplitLine())
		if testLogger != nil {
			testLogger.Print(testutil.GetSplitLine())
		}
		testutil.Rc.Print(testutil.GetRoundString(i + 1))
		if testLogger != nil {
			testLogger.Print(testutil.GetRoundString(i + 1))
		}
		for {
			fmt.Print(testutil.GetStageString(testutil.STAGE_ONE))
			if testLogger != nil {
				testLogger.Print(testutil.GetStageString(testutil.STAGE_ONE))
			}
			if testentry.status_check() {
				testutil.Pc.Print(testutil.GetStagePassString(testutil.STAGE_ONE))
				if testLogger != nil {
					testLogger.Print(testutil.GetStagePassString(testutil.STAGE_ONE))
				}
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Print(testutil.GetStageString(testutil.STAGE_TWO))
			if testLogger != nil {
				testLogger.Print(testutil.GetStageString(testutil.STAGE_TWO))
			}
			if testentry.delete_pod() {
				testutil.Pc.Print(testutil.GetStagePassString(testutil.STAGE_TWO))
				if testLogger != nil {
					testLogger.Print(testutil.GetStagePassString(testutil.STAGE_TWO))
				}
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Print(testutil.GetStageString(testutil.STAGE_THREE))
			if testLogger != nil {
				testLogger.Print(testutil.GetStageString(testutil.STAGE_THREE))
			}
			if testentry.status_check() {
				testutil.Pc.Print(testutil.GetStagePassString(testutil.STAGE_THREE))
				if testLogger != nil {
					testLogger.Print(testutil.GetStagePassString(testutil.STAGE_THREE))
				}
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		for {
			fmt.Print(testutil.GetStageString(testutil.STAGE_FOUR))
			if testLogger != nil {
				testLogger.Print(testutil.GetStageString(testutil.STAGE_FOUR))
			}
			if testentry.func_check() {
				testutil.Pc.Print(testutil.GetStagePassString(testutil.STAGE_FOUR))
				if testLogger != nil {
					testLogger.Print(testutil.GetStagePassString(testutil.STAGE_FOUR))
				}
				break
			}
			count++
			if count >= testutil.RETRY_TIMES {
				testutil.Fc.Print(testutil.GetStageFailString(testutil.STAGE_FOUR))
				if testLogger != nil {
					testLogger.Print(testutil.GetStageFailString(testutil.STAGE_FOUR))
				}
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}
	}
	testutil.Cc.Print(testutil.GetCompleteString(testentry.Name))
	if testLogger != nil {
		testLogger.Print(testutil.GetCompleteString(testentry.Name))
	}
}
