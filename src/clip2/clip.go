package main

import (
	"fmt"
	"github.com/davidjairala/clipboard-go"
)

func main() {
	resultFile, _ := clipboard.GetClipboard()
	fmt.Println(resultFile.Filename)
}