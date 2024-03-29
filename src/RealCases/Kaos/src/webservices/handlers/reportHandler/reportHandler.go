// reportHandler project reportHandler.go
package reportHandler

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"testutil"

	"github.com/labstack/echo"
)

func GetReports(c echo.Context) error {
	var files [][]string
	root := "./"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		var file []string
		name := info.Name()
		match, _ := regexp.MatchString("^([0-9]+)_rpt.csv$", name)
		if match {
			tf := info.ModTime().Format(testutil.TFORMAT)
			file = append(file, name)
			file = append(file, tf)
			files = append(files, file)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	// for _, file := range files {
	// 	fmt.Println(file)
	// }

	return c.JSON(http.StatusOK, map[string][][]string{
		"files": files})
}

func GetReportView(c echo.Context) error {
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
