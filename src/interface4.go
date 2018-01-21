package main

import (
	"flag"
	"time"
	"fmt"
)

/*实现flag的value接口，查看接口的定义如下：
type  Value interface {
	String() string
	Set(string) error
}
*/
func main() {
	var period = flag.Duration("period", 1 * time.Second, "sleep period")
	flag.CommandLine.Var()
			flag.Parse()

	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
