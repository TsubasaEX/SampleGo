// reportHandler project reportHandler.go
package reportHandler

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
)

func GetReports(c echo.Context) error {
	var files []string
	root := "./"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	return c.JSON(http.StatusNotFound, map[string][]string{
		"files": files})
}

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
