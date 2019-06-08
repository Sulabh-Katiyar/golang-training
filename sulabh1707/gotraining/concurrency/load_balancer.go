package main

import (
	"fmt"
	"time"
)

// type person struct {
//  name string
//  age  int
// }

// var println = fmt.Println

func server123(queue1 chan int) {
	limiter := time.Tick(200 * time.Millisecond)
	// time.Sleep(100)

	for que := range queue1 {
		<-limiter
		println(que)
	}

}

// func server2(queue chan int) {

// 	// time.Sleep(300)

// 	for {
// 		println(<-queue)
// 	}

// }

func main() {
	len := 16
	queue1 := make(chan int, len)
	//   queue2 := make(chan int, 2)
	go server123(queue1)
	// go server2(queue2)
	for i := 0; i < 16; i++ {
		queue1 <- i
	}
	fmt.Scanln()

}
