/**
 *
 *  Printf使用%v不能打印有多个返回值的函数
 **/

package main
import(
	"fmt"
)

func max(list []int) (maxV int){
	if len(list) == 0 {
		return
	}

	maxV = list[0]
	for _, value := range list {
		if (maxV < value) {
			maxV = value
		}
	}
	return
}

func min(list []int) (minV int) {
	if len(list) == 0 {
		return
	}

	minV = list[0]
	for _, value := range list {
		if (minV > value) {
			minV = value
		}
	}
	return
}

func end(list []int)(minV, maxV int) {
	if len(list) == 0 {
		return
	}
	
	minV, maxV = list[0], list[0]
	for _, value := range list {
		if (minV > value) {
			minV = value
		}
		if (maxV < value){
			maxV = value
		}
	}

	return
}

func main() {
	arr := [5]int{1,2,3,4,5}
	fmt.Printf("max value: %d\n", max(arr[:]))
	fmt.Printf("min value: %d\n", min(arr[:]))

	minV, maxV := end(arr[:])
	fmt.Printf("min and max value: %d, %d\n", minV, maxV)
}