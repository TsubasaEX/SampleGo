// HelloWorld project main.go
package main

// need to modify GOAPTH to D:\GIT\SampleGo
// windows - setx GOPATH D:\GIT\SampleGo
import (
	"exportfunc"
	"fmt"
)

func main() {
	news := exportfunc.News{"abc", "123"}
	fmt.Println(news)
	fmt.Println(exportfunc.Pi)
	fmt.Println(exportfunc.Vi)
	fmt.Println(exportfunc.Print())
}
