// reportHandler project reportHandler.go
package reportHandler

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func GetReports(c echo.Context) error {

	file, err := os.OpenFile("20191125154904_rpt.csv", os.O_RDONLY, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return c.Render(http.StatusOK, "report.html", data[1:])
}

func GetStatisticsReports(c echo.Context) error {

	file, err := os.OpenFile("20191125154904_stats.csv", os.O_RDONLY, 0666)
	defer file.Close()

	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	return c.Render(http.StatusOK, "statistics.html", data[1:])
}
