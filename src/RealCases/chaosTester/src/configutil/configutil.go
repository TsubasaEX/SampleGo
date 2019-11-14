// configutil project configutil.go
package configutil

type Config struct {
	APIVersion string `yaml:"apiVersion"`
	IP         string `yaml:"ip"`
	Apps       struct {
		Web []struct {
			Name   string `yaml:"name"`
			Port   string `yaml:"port"`
			Enable bool   `yaml:"enable"`
			Label  string `yaml:"label"`
			Times  int    `yaml:"times"`
			Report bool   `yaml:"report"`
		} `yaml:"web"`
		App []struct {
			Name   string `yaml:"name"`
			Enable bool   `yaml:"enable"`
			Label  string `yaml:"label"`
			Times  int    `yaml:"times"`
			Report bool   `yaml:"report"`
		} `yaml:"app"`
	} `yaml:"apps"`
}
