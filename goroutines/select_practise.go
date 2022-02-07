package goroutines

import (
	"fmt"
	"time"
)

func Init2() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "One" // reciiving message in channel
	}()

	go func() {
		time.Sleep(8 * time.Second)
		c2 <- "Two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1: // this is blocking way
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-time.After(5 * time.Second): // this is for time out if the second loop doesnt complete in 5 second it willtime out check the output
			fmt.Println("Timeout")

			//	default: // this in non blocking way
			//		fmt.Println("default")
		}

	}

}
