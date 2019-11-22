// portal project portal.go
package portal

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"testutil"
	"time"

	"github.com/briandowns/spinner"
)

type TestEntity struct {
	IP    string
	Name  string
	Label string
	Times int
	Port  string
}

var b_Simple = false
var testLogger *log.Logger

const PASS = "PASS"
const FAIL = "FAIL"

func (testentity *TestEntity) status_check() bool {
	str := "kubectl get pods -l " + testentity.Label + " --field-selector=status.phase=Running"
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

func (testentity *TestEntity) delete_pod() bool {
	str := "kubectl delete pods -l " + testentity.Label + " --field-selector=status.phase=Running"
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

func (testentity *TestEntity) func_check() bool {
	str := "http://" + testentity.IP + ":" + testentity.Port

	if !b_Simple {
		fmt.Println(testutil.Emoji(testutil.ICON_PENCIL), str)
		if testLogger != nil {
			testLogger.Println(testutil.Emoji(testutil.ICON_PENCIL), str)
		}
	}
	_, err := http.Get(str)

	if err != nil {
		fmt.Println(err.Error())
		if testLogger != nil {
			testLogger.Println(err.Error())
		}
		return false
	}
	return true
}

func (testentity *TestEntity) Test(args []string, simple bool, testlogger *log.Logger, testreporter *csv.Writer) int {
	b_Simple = simple
	testLogger = testlogger
	var sp *spinner.Spinner
	testutil.Rc.Print(testutil.GetTestString(testentity.Name))
	if testLogger != nil {
		testLogger.Print(testutil.GetTestString(testentity.Name))
	}
	if b_Simple {
		sp = spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	}
	testReporter := testreporter
	var records [][]string

	passNum := 0
	for i := 0; i < testentity.Times; i++ {
		count := 0
		fmt.Print(testutil.GetSplitLine())
		if testLogger != nil {
			testLogger.Print(testutil.GetSplitLine())
		}
		testutil.Rc.Print(testutil.GetRoundString(i + 1))
		if testLogger != nil {
			testLogger.Print(testutil.GetRoundString(i + 1))
		}
		if sp != nil {
			sp.Suffix = testutil.GetStageString(testutil.STAGE_ONE)
			sp.Start()
		}
		for {
			if sp == nil {
				fmt.Println(testutil.GetStageString(testutil.STAGE_ONE))
			}
			if testLogger != nil {
				testLogger.Print(testutil.GetStageString(testutil.STAGE_ONE))
			}
			if testentity.status_check() {
				if sp != nil {
					sp.Stop()
				}
				testutil.Pc.Print(testutil.GetStagePassString(testutil.STAGE_ONE))
				if testLogger != nil {
					testLogger.Print(testutil.GetStagePassString(testutil.STAGE_ONE))
				}
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		if sp != nil {
			sp.Suffix = testutil.GetStageString(testutil.STAGE_TWO)
			sp.Start()
		}
		for {
			if sp == nil {
				fmt.Println(testutil.GetStageString(testutil.STAGE_TWO))
			}
			if testLogger != nil {
				testLogger.Print(testutil.GetStageString(testutil.STAGE_TWO))
			}
			if testentity.delete_pod() {
				if sp != nil {
					sp.Stop()
				}
				testutil.Pc.Print(testutil.GetStagePassString(testutil.STAGE_TWO))
				if testLogger != nil {
					testLogger.Print(testutil.GetStagePassString(testutil.STAGE_TWO))
				}
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}

		if sp != nil {
			sp.Suffix = testutil.GetStageString(testutil.STAGE_THREE)
			sp.Start()
		}
		for {
			if sp == nil {
				fmt.Println(testutil.GetStageString(testutil.STAGE_THREE))
			}
			if testLogger != nil {
				testLogger.Print(testutil.GetStageString(testutil.STAGE_THREE))
			}
			if testentity.status_check() {
				if sp != nil {
					sp.Stop()
				}
				testutil.Pc.Print(testutil.GetStagePassString(testutil.STAGE_THREE))
				if testLogger != nil {
					testLogger.Print(testutil.GetStagePassString(testutil.STAGE_THREE))
				}
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}
		if sp != nil {
			sp.Suffix = testutil.GetStageString(testutil.STAGE_FOUR)
			sp.Start()
		}
		for {
			if sp == nil {
				fmt.Println(testutil.GetStageString(testutil.STAGE_FOUR))
			}
			if testLogger != nil {
				testLogger.Print(testutil.GetStageString(testutil.STAGE_FOUR))
			}
			if testentity.func_check() {
				if sp != nil {
					sp.Stop()
				}
				testutil.Pc.Print(testutil.GetStagePassString(testutil.STAGE_FOUR))
				if testLogger != nil {
					testLogger.Print(testutil.GetStagePassString(testutil.STAGE_FOUR))
				}

				record := testutil.GetRecord(i+1, testentity.Name, PASS)
				records = append(records, record)
				passNum++
				break
			}
			count++
			if count >= testutil.RETRY_TIMES {
				if sp != nil {
					sp.Stop()
				}
				testutil.Fc.Print(testutil.GetStageFailString(testutil.STAGE_FOUR))
				if testLogger != nil {
					testLogger.Print(testutil.GetStageFailString(testutil.STAGE_FOUR))
				}
				record := testutil.GetRecord(i+1, testentity.Name, FAIL)
				records = append(records, record)
				break
			}
			time.Sleep(10000 * time.Millisecond) //10 sec per status check
		}
	}

	testReporter.WriteAll(records) // calls Flush internally

	if err := testReporter.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	testutil.Cc.Print(testutil.GetCompleteString(testentity.Name))
	if testLogger != nil {
		testLogger.Print(testutil.GetCompleteString(testentity.Name))
	}
	return passNum
}
