module chaosTester

go 1.13

replace configutil => ../configutil

replace testkicker => ../testkicker

replace app/edgesense/portal => ../app/edgesense/portal

replace testutil => ../testutil

replace webservices => ../webservices

replace webservices/webutil => ../webservices/webutil

replace webservices/handlers/reportHandler => ../webservices/handlers/reportHandler

replace webservices/handlers/statsHandler => ../webservices/handlers/statsHandler

require (
	app/edgesense/portal v0.0.0-00010101000000-000000000000 // indirect
	configutil v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/wcharczuk/go-chart v2.0.1+incompatible // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
	gopkg.in/yaml.v3 v3.0.0-20191026110619-0b21df46bc1d
	testkicker v0.0.0-00010101000000-000000000000
	testutil v0.0.0-00010101000000-000000000000 // indirect
	webservices v0.0.0-00010101000000-000000000000
	webservices/handlers/reportHandler v0.0.0-00010101000000-000000000000 // indirect
	webservices/handlers/statsHandler v0.0.0-00010101000000-000000000000 // indirect
	webservices/webutil v0.0.0-00010101000000-000000000000 // indirect
)
