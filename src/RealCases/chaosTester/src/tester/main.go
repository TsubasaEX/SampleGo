// chaosTester project main.go
package main

import (
	"configutil"
	"os"
	"sync"

	"testkicker"

	"io/ioutil"
	"log"

	"webservices"

	"gopkg.in/yaml.v3"
)

var wg sync.WaitGroup

func startTest(ch chan string) {
	defer wg.Done()
	argsWithoutProg := os.Args[1:]
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
	ch <- testkicker.Kick(config, argsWithoutProg)
}

func main() {
	// ch := make(chan string, 1)
	// go startTest(ch)
	// wg.Add(1)
	// wg.Wait()
	// close(ch)
	// _, ok := <-ch
	// if ok {
	// 	webservices.Start("123")
	// } else {
	// 	// fmt.Println("Starting webservices failed!!")
	// }

	webservices.Start("20191125154904_stats.csv")
}
