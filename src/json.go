package main

import (
	"fmt"
	"encoding/json"
)

type Point struct {
	Lat float64
	Lng float64
}

func main() {
	a := make(map[string]*Point, 6)
	fmt.Println(len(a))

	data := map[string]interface{}{
		"errno": 1,
		"errmsg": "this is error",
		"data": 1,
	}

	//result, _ := json.Marshal(data)
	result, _ := json.MarshalIndent(data, "", "	")
	fmt.Println(result)

	var m map[string]int
	fmt.Println(len(m))
}
