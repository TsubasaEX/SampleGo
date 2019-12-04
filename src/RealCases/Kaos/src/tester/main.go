// chaosTester project main.go
package main

import (
	"configutil"
	"fmt"
	"os"
	"sync"

	"testkicker"

	"io/ioutil"
	"log"
	"strings"

	"webservices"

	"gopkg.in/yaml.v3"
)

var wg sync.WaitGroup
var args []string

func startTest(ch chan string) {
	defer wg.Done()
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
	ch <- testkicker.Kick(config, args)
}

func main() {
	b_Test := true
	b_Web := false
	args = os.Args[1:]

	if len(args) != 0 {
		split := strings.Split(args[0], "-")
		if strings.Contains(split[1], "w") {
			b_Web = true
		}
		if strings.Contains(split[1], "n") {
			b_Test = false
		}
	}

	if b_Test {
		ch := make(chan string, 1)
		go startTest(ch)
		wg.Add(1)
		wg.Wait()
		close(ch)
		_, ok := <-ch
		if ok {
			webservices.Start()
		} else {
			fmt.Println("Starting webservices failed!!")
		}
	}

	if b_Web {
		webservices.Start()
	}

	if !b_Test && !b_Web {
		fmt.Println("Hey!! WTF do you actually want?!")
	}
}
