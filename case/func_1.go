/**
 *
 *
 *
 **/

package main
import (
 	"fmt"
)

type fibonacci struct {
	n int
	data []int
	maps map[int]int
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