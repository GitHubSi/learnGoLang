package main
import (
	"fmt"
	"time"
)

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

// 在main goroutine线，期望从管道中获得一个数据，而这个数据必须是其他goroutine线放入管道的 
// 但是其他goroutine线都已经执行完了(all goroutines are asleep)，那么就永远不会有数据放入管道。 
// 所以，main goroutine线在等一个永远不会来的数据，那整个程序就永远等下去了。 
// 统自动在除了主协程之外的协程都关闭后，做的检查，继而报出的错误， 
func main() {
	/*
	for i:= 0; i < 100; i++ {
		go Add(i, i)
	}
	*/

	/*
	messages := make(chan string)	//channel缓冲区为0
	messages <- "ping"				//在主goroutine中写入，因为没有读取，会造成deadlock
	fmt.Println(<-messages)
	*/

	/*
	messages := make(chan string, 1)	//channel缓冲区为1
	messages <- "ping"					//在主goroutine中写入，存在缓存区，正常读取
	fmt.Println(<-messages)
	*/

	/*
	messages := make(chan string)
	go func () {						//启用另一个goroutine写入channel
		fmt.Println("ping-0")
		messages <- "ping"
		fmt.Println("ping-2")
	}()
	
	time.Sleep(2 * time.Second) 
	fmt.Println("ping-1")
	<-messages
	time.Sleep(2 * time.Second) 
	*/
}