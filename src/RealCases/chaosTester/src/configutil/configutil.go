// configutil project configutil.go
package configutil

type Config struct {
	Apps []struct {
		Name   string `yaml:"name"`
		Enable bool   `yaml:"enable"`
		Label  string `yaml:"label"`
		Times  int    `yaml:"times"`
		Report bool   `yaml:"report"`
	} `yaml:"apps"`
}

func abc() string {
	return "123"
}
