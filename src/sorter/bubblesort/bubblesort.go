package bubblesort

func BubbleSort(value []int) {
	len := len(value)
	for i:=0; i<len-1; i++ {
		for j:=0; j<len-i-1; j++ {
			if value[j] > value[j+1] {
				value[j], value[j+1] = value[j+1], value[j]
			}
		}
	}
}