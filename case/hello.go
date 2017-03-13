package main 

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "world"
	fmt.Printf("%v", os.Args[1:])
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[0:], " ")
	}
	fmt.Println("Hello", who)
}