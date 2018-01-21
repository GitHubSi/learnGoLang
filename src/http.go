package main
import "fmt"

func test6(a interface{}) string {
	return string(a)
}

func main()  {
	test6("fuhui")
}
/*
import (
	"io"
	"log"
	"net/http"
	"strings"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")

	var action string = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}
	io.WriteString(w,action)
}

func main() {
	http.HandleFunc("/", helloHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
*/
