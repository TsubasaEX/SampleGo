// reportHandler project reportHandler.go
package reportHandler

import (
	"fmt"
	"log"
	"net/http"
	"testutil"

	"github.com/labstack/echo"
)

type Record struct {
	No    int
	Name  string
	Pass  int
	Fail  int
	Total int
	PR    float32
	Ts    string
}

func GetStatisticsReports(c echo.Context) error {

	teststReporter_R := testutil.GetStatisticsReporter_R()
	data, err := teststReporter_R.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// var records []Record

	for k, v := range data {
		fmt.Println(k, v[0])
		fmt.Println(k, v[1])
		fmt.Println(k, v[2])
		fmt.Println(k, v[3])
		fmt.Println(k, v[4])
	}
	// fmt.Println(records)

	// type User struct {
	// 	Id   int
	// 	Name string
	// }

	// type UserList []User
	// var myuserlist UserList = UserList{
	// 	{1, "a"},
	// 	{2, "b"},
	// 	{3, "c"},
	// }

	fmt.Print("-----------")
	testutil.CloseStatisticsReporter_R()
	return c.Render(http.StatusOK, "statistics.html", data)
	// return c.Render(http.StatusOK, "statistics.html", map[string]interface{}{
	// 	"name": "HOME",
	// 	"msg":  "Hello, Boatswain!",
	// })

	// return c.String(http.StatusOK, webutil.LatestFileName)
}
