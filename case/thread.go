package main

import (
	"fmt"
	"sync"
	// "runtime"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	// lock.Lock()
	counter++
	fmt.Println(counter, "-")
	// lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}		//sync.Mutex is not an expression
	for i:=0; i<10; i++ {
		go Count(lock)
	}

	for {
		c := counter
		if (c >= 5 ){
			fmt.Println("total ", c)
			break;
		}
	}
}