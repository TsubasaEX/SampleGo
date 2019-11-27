// reportHandler project reportHandler.go
package reportHandler

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func GetReportResults(c echo.Context) error {
	name := c.QueryParam("file")
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	defer file.Close()

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"file":  name,
			"error": "not found"})
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return c.Render(http.StatusOK, "report.html", data[1:])
}
