package goroutines

import (
	"fmt"
	"time"
)

func Init() {
	//routineMain() //uncomment if i wan to use routine main
	//Channels

	mChannel := make(chan string, 5) // 3 is used to make more then one channel
	mChannel <- "mainThread"         // this is main thread
	mChannel <- "MainThread"         //this is recieve only channel from line 49

	go func() {
		//	fmt.Println("testing")
		mChannel <- " anynomous function goroutine"

	}()

	go routine1(mChannel) // this is send only channel
	go routine2(mChannel) // this is send only channel

	msg := <-mChannel
	fmt.Println(msg)
	msg = <-mChannel //while using more then one channel use only = instead of := as in line 20
	fmt.Println(msg)
	msg = <-mChannel
	fmt.Println(msg)
	msg = <-mChannel
	fmt.Println(msg)
	msg = <-mChannel
	fmt.Println(msg)

}

func routine1(msg chan string) { // this is bi directional channel

	msg <- "routine1"

}

func routine2(msg chan<- string) { // this is send only channel

	msg <- "routine2"

}

// below program is for routine
func routineMain() {

	fmt.Println("First") //main thread
	routine()
	fmt.Println("Second") //main thread
	go routine()          // routine thread
	go func() {           // routine thread
		fmt.Println("Third")
	}()

	go func(msg string) { // to print we can also used this way instead of line 13 and 14  // routine thread
		fmt.Println(msg)
	}("testing")

	time.Sleep(time.Second) // can also use Millisecond or any other time
	fmt.Println("Done")
}

func routine() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

//func routine() {  // this is for example to see time execution complexity.. inorder to run this i have to put line 17 in comment
//for i := 0; i < 1000; i++ {
//for j := 0; i < 1000; j++ {
//for k := 0; i < 1000; k++ {
//fmt.Println(i, j, k)
//		}
//		}
//	}
//}
