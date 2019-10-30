// yamlParser project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
// type T struct {
// 	A string
// 	B struct {
// 		RenamedC int   `yaml:"c"`
// 		D        []int `yaml:",flow"`
// 	}
// }

type TestDetail struct {
	Apps struct {
		Portal struct {
			Times int `yaml:"times"`
		} `yaml:"portal"`
		Worker struct {
			Times int `yaml:"times"`
		} `yaml:"worker"`
	} `yaml:"apps"`
}

func main() {

	log.SetFlags(log.Lshortfile)

	data, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	o := TestDetail{}
	err = yaml.Unmarshal(data, &o)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(o)
}
