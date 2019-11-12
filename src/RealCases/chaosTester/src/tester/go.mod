module chaosTester

go 1.13

replace configutil => ../configutil

replace testkicker => ../testkicker

replace app/edgesense/portal => ../app/edgesense/portal

replace testutil => ../testutil

require (
	app/edgesense/portal v0.0.0-00010101000000-000000000000 // indirect
	configutil v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.7.0 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	gopkg.in/yaml.v3 v3.0.0-20191026110619-0b21df46bc1d
	testkicker v0.0.0-00010101000000-000000000000
	testutil v0.0.0-00010101000000-000000000000 // indirect
)
