// worker project worker.go
package worker

import (
	"encoding/csv"
	"log"
)

type TestEntity struct {
	IP    string
	Name  string
	Label string
	Times int
}

func (testentity *TestEntity) Test(args []string, simple bool, testlogger *log.Logger, testreporter *csv.Writer) int {
	return 0
}
