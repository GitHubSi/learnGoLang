package main

import (
	"fmt"
)

func main() {
	var errors = [...]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println(errors[1])
}