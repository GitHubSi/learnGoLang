package main
import (
	"fmt"
	"time"
)

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	/*
	for i:= 0; i < 100; i++ {
		go Add(i, i)
	}
	*/

	messages := make(chan string)
	go func () {
		fmt.Println("ping-0")
		messages <- "ping"
		fmt.Println("ping-2")
	}()
	
	time.Sleep(2 * time.Second) 
	// messages <- "pong"
	fmt.Println("ping-1")
	<-messages
	time.Sleep(2 * time.Second) 
	//fmt.Println(<-messages)
}