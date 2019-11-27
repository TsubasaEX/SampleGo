// statsHandler project statsHandler.go
package statsHandler

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func GetStatisticsResults(c echo.Context) error {
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
	file.Close()
	return c.Render(http.StatusOK, "statistics.html", data[1:])
}
