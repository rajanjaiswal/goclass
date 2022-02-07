package goroutines

import (
	"fmt"
)

func InitClosingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			msg, more := <-jobs
			if more {
				fmt.Println("Receiving jobs", msg)
			} else {
				fmt.Println("all jobs received")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
	//time.Sleep(2 * time.Second) // we also use this but this is fixed and also dont know how long it takes to stop so we used done line 10,19,31
}
