package main
import (
	"fmt"
)

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

// 在main goroutine线，期望从管道中获得一个数据，而这个数据必须是其他goroutine线放入管道的 
// 但是其他goroutine线都已经执行完了(all goroutines are asleep)，那么就永远不会有数据放入管道。 
// 所以，main goroutine线在等一个永远不会来的数据，那整个程序就永远等下去了。 
// 统自动在除了主协程之外的协程都关闭后，做的检查，继而报出的错误， 

func consumer (data chan int, done chan bool) {
	for x := range data {
		println("receive:", x)
	}

	fmt.Println("consumer chan %v", &data)
	done <- true
}

func procucer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
	}

	fmt.Println("producer chan %v", &data)
	close(data)
}


func main() {
	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)
	go procucer(data)

	<-done

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
	time.Sleep(time.Second) 
	messages := make(chan string)
	messages <- "ping"				
	select {						// all goroutines are asleep - deadlock!  只能检测到之前的，检测不到之后的
		case value := <- messages:
			fmt.Println("read ", value)
		default:
			fmt.Println("read null")
	}
	*/

	/*
	time.Sleep(time.Second) 
	messages := make(chan int, 1)
	// num := 1
	for {
		select {
		case messages <- 0 :		//？？？
		case messages <- 1 :
		}

		i := <- messages
		fmt.Println(i)

		break
		num++
		if num == 8 {
			break
		}
	}
	*/

	/*
	messages := make(chan string)

	timeout := make(chan bool)
	go func () {
		time.Sleep(time.Second * 2)
		timeout <- true
	}()

	go func () {
		time.Sleep(time.Second * 1)
		messages <- "test"
	}()

	select {
		case value := <-messages:
			fmt.Println("read messages:", value)
		case <- timeout:
	}

	go func () {
		messages <- "msg"
	}()
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