module chaosTester

go 1.13

replace configutil => ../configutil

replace testkicker => ../testkicker

replace app/edgesense/portal => ../app/edgesense/portal

replace testutil => ../testutil

replace webservices => ../webservices

replace webservices/handlers/reportHandler => ../webservices/handlers/reportHandler

require (
	app/edgesense/portal v0.0.0-00010101000000-000000000000 // indirect
	configutil v0.0.0-00010101000000-000000000000
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/wcharczuk/go-chart v2.0.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20191119213627-4f8c1d86b1ba // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
	gopkg.in/yaml.v3 v3.0.0-20191026110619-0b21df46bc1d
	testkicker v0.0.0-00010101000000-000000000000
	testutil v0.0.0-00010101000000-000000000000 // indirect
	webservices v0.0.0-00010101000000-000000000000 // indirect
	webservices/handlers/reportHandler v0.0.0-00010101000000-000000000000 // indirect
)
