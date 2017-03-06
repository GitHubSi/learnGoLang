/**
 * 1. 变参其实传递的是一个slice
 * 2. ... 符号不要忘记在参数中带上
 *
 **/
package main
import (
	"fmt"
)

func displayParam(param ...int) {
	for _, v := range param {
		fmt.Printf("%d\n", v);
	}
}

func main() {
	var arr [3]int = [3]int{1,2,3}
	displayParam(arr[:]...)
}

