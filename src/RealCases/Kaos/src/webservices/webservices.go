// webservices project webservices.go
package webservices

import (
	"html/template"
	"io"
	"net/http"

	"webservices/handlers/reportHandler"

	"github.com/labstack/echo"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates *template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// t.templates.Execute()
	// fmt.Println(t.templates.ExecuteTemplate(w, name, data))
	return t.templates.ExecuteTemplate(w, name, data)
}

func initHandlers(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/reports", reportHandler.GetReports)
	e.GET("/reports/statistics", reportHandler.GetStatisticsReports)
}

func Start() {
	e := echo.New()
	// Instantiate a template registry and register all html files inside the view folder
	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("../webservices/views/*.html")),
	}
	initHandlers(e)
	e.Logger.Fatal(e.Start(":1323"))
}
