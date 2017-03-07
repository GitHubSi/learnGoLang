package main

import "fmt"

func sort(list []int) []int {
	length := len(list)
	for i:= 0; i< length; i++ {
		temp := 0
		for j := 0; j< length; j++ {
			if list[j] < list[temp] {
				list[j], list[temp] = list[temp], list[j]	
			} 
			temp = j
		}
	}
	return list
}

func main() {
	arr := []int{2,4,1,7,3,5}
	newArr := sort(arr)
	fmt.Printf("%v\n", newArr)
}