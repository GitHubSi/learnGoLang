/*
 *	注释调default当循环执行两次之后发生：fatal error: all goroutines are asleep - deadlock!
 *	因为系统检测到没有goroutine向channel中些数据啦
 */

package main 
import (
	"fmt"
	"time"
)

func main() {
	channelA := make(chan string)
	channelB := make(chan string)

	go func () {
		channelA <- "a"
	}()

	go func () {
		channelB <- "b"
	}()

	time.Sleep(time.Second * 1)

	count := 0
	for {
		select {
		case a := <- channelA:
			fmt.Println("result:", a)
		case b := <- channelB:
			fmt.Println("result:", b)
		default:
		 	fmt.Println("over")
		}
		count ++

		if count == 6 {
			break
		}
	}
}