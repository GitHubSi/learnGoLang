/**
 * 返回函数的类型func(a int) int
 *	返回的函数是一个匿名函数
 **/
package main
import (
	"fmt"
)

func plusTwo(x int) func(int) int {
	return func (a int) int{
		return a + x
	}
}

func main() {
	p := plusTwo(5);
	fmt.Printf("%v\n", p(3))
}