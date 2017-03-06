/**
 *	在main函数中如果不是用make，则data的lenght和cap都为0
 *	在add方法中明确需要返回s.data[n]	
 *	add方法中fibonacci中参数需要指定为指针，在main中声明的时候Go会转换成指针对象
 **/

package main
import (
 	"fmt"
)

type fibonacci struct {
	n int
	data []int
}

func(s *fibonacci) add(n int)(v int){
	if n == 0 || n == 1{
		s.data[n] = 1
		return 1
	}	

	s.data[n] = s.add(n-1) + s.add(n-2)
	return s.data[n]
}

func main() {
	var sa fibonacci
	sa.data = make([]int, 10)
	sa.add(7)
	fmt.Printf("%v\n", sa.data)
}