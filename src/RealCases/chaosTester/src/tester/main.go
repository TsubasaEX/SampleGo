// chaosTester project main.go
package main

import (
	"configutil"
	"fmt"

	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	data, err := ioutil.ReadFile("Config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	config := configutil.Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config)
}
