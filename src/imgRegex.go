package main

import (
	"regexp"
	"fmt"
)

var imgFilter = "<img\\s*src=(?:\"|')(.+?)(?:\"|')"
//var imgFilter = "<img\\s*src=(\"|')"

func main() {
	imgRegex := regexp.MustCompile(imgFilter)

	exampleString := "<img src=\"//upload.jianshu.io/users/upload_avatars/3548266/cb93426b-cf0e-4ed9-b936-59997f4607dd.jpg?imageMogr2/auto-orient/strip|imageView2/1/w/144/h/144\" alt=\"144\" />"
	dd := imgRegex.FindAllStringSubmatch(exampleString, 1)

	for index, v := range dd[0] {
		fmt.Println(index, v)
	}
}