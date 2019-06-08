package main

import "time"
import "fmt"

func news1(c1 chan<- string) {
	time.Sleep(1 * time.Second)
	c1 <- "one"
}

func news2(c2 chan<- string) {
	time.Sleep(1 * time.Second)
	c2 <- "two"
}

func image2(c3 chan<- string) {
	time.Sleep(4 * time.Second)
	c3 <- "three"
}

func image1(c4 chan<- string) {
	time.Sleep(2 * time.Second)
	c4 <- "four"
}

func site1(c5 chan<- string) {
	//time.Sleep(2 * time.Second)
	c5 <- "five"
}

func site2(c6 chan<- string) {
	time.Sleep(1 * time.Second)
	c6 <- "six"
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	c4 := make(chan string)
	c5 := make(chan string)
	c6 := make(chan string)

	go news1(c1)
	go news2(c2)
	go image1(c3)
	go image2(c4)
	go site1(c5)
	go site2(c6)

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	case <-time.After(10 * time.Second):
		fmt.Println("timeout 1")

	}

	select {
	case msg3 := <-c3:
		fmt.Println("received", msg3)
	case msg4 := <-c4:
		fmt.Println("received", msg4)
	case <-time.After(10 * time.Second):
		fmt.Println("timeout 1")

	}

	select {
	case msg5 := <-c5:
		fmt.Println("received", msg5)
	case msg6 := <-c6:
		fmt.Println("received", msg6)
	case <-time.After(10 * time.Second):
		fmt.Println("timeout 1")

	}
}
