/**
 *	将参数按照正序或者逆序排列返回
 *	将可变参数列表直接返回成切片数组(...float64 -> []float64)
 *	通过修改index，使用插入排序来处理数组
 *	声明切片的方式[]float64{1}
 */
 
package main

import "fmt"

func getOrder(number ...float64)(result []float64) {
	length := len(number)
	for i := 1; i < length; i++ {
		temp := i
		for j := temp-1; j >= 0; j-- {
			if number[j] > number[temp] {
				number[temp], number[j] = number[j], number[temp]
				temp--				
			} else {
				break
			}
		}
	}

	return number;
}

func main() {
	slices := []float64{3,2,5,1,6,4}
	result := getOrder(slices...)
	fmt.Printf("%f\n", result) 	
}