/**
 *	函数作为值传递的时候需要完整的声明，包括参数，返回值，不然会报比如“fun(value) used as value”这样的错误
 *	调用函数的时候只需要传递函数名
 *	[...]int{1,2,3}这种形式定义的是一个函数，不是切片
 *	go语言中类型都在参数的后面声明
 **/

package main
import "fmt"

func filter(fun func(int) int, params []int) (result []int){
	if len(params) == 0 {
		return
	}

	result = make([]int, len(params))
	for index, value := range params {
		result[index] = fun(value)
	}
	return result
} 

func double(v int) int {
	return 2 * v
}

func main() {
 	a := [...]int{1,2,3}
 	result := filter(double, a[:])
 	/*b := func (v int) int{
 		return 2 * v
 	}
 	result := filter(b, a[:])
 	*/
 	fmt.Printf("%v\n", result)
 }