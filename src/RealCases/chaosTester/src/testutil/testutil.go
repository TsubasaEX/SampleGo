// testutil project testutil.go
package testutil

import (
	"fmt"
	"html"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type TestFunc interface {
	Test([]string,
		bool,
		*log.Logger)
}

var Rc = color.New(color.FgCyan, color.Bold)
var Pc = color.New(color.FgGreen, color.Bold)
var Fc = color.New(color.FgRed, color.Bold)
var Cc = color.New(color.FgYellow, color.Bold)

var logFile *os.File

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
const ICON_SNOWMAN = "\\U000026C4"

func Emoji(s string) string {
	// Hex String
	h := strings.ReplaceAll(s, "\\U", "0x")

	// Hex to Int
	i, _ := strconv.ParseInt(h, 0, 64)

	// Unescape the string (HTML Entity -> String).
	str := html.UnescapeString(string(i))
	return str
}

func GetTestString(s string) string {
	return fmt.Sprintln(Emoji(ICON_SCISSOR), TEST_AGAINST, s, DOT_STR, Emoji(ICON_HEART))
}

func GetSplitLine() string {
	str := ""
	for i := 0; i < 40; i++ {
		str += Emoji(ICON_ARROW)
	}
	return fmt.Sprintln(str)
}
func GetRoundString(n int) string {
	return fmt.Sprintln(Emoji(ICON_AIRPLANE), ROUND, n)
}

func GetStageString(s string) string {
	return fmt.Sprintln(Emoji(ICON_SPARKLE), s, DOT_STR, Emoji(ICON_HEART))
}

func GetStagePassString(s string) string {
	return fmt.Sprintln(Emoji(ICON_CHECK), PASS, s)
}

func GetStageFailString(s string) string {
	return fmt.Sprintln(Emoji(ICON_CROSS), FAIL, s)
}

func GetCompleteString(s string) string {
	return fmt.Sprintln(Emoji(ICON_SNOWMAN), COMPLETE, TEST_AGAINST, s)
}

func GetLogger() *log.Logger {
	t := time.Now()
	logFile, err := os.OpenFile(t.Format("20060102150405")+"_"+"logfile.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}

	logger := log.New(logFile, "[LOG]", log.Ldate|log.Lmicroseconds)
	return logger
}

func CloseLogger() {
	logFile.Close()
}
