package main

import "fmt"
import "time"

// type person struct {
//  name string
//  age  int
// }

var println = fmt.Println

func server1(news chan string) {
	time.Sleep(1 * time.Second)
	news <- "News"
}

func server2(news chan string) {
	time.Sleep(1 * time.Second)
	news <- "News"
}

func server3(news chan string) {
	time.Sleep(1 * time.Second)
	news <- "News"
}

func server4(image chan string) {
	time.Sleep(1 * time.Second)
	image <- "image"
}

func server5(image chan string) {
	time.Sleep(1 * time.Second)
	image <- "image"
}

func server6(image chan string) {
	time.Sleep(1 * time.Second)
	image <- "image"
}

func server7(video chan string) {
	time.Sleep(1 * time.Second)
	video <- "video"
}

func server8(video chan string) {
	time.Sleep(1 * time.Second)
	video <- "video"
}

func server9(video chan string) {
	time.Sleep(1 * time.Second)
	video <- "video"
}

func main() {
	news := make(chan string, 5)
	image := make(chan string, 5)
	video := make(chan string, 5)

	go server1(news)
	go server2(news)
	go server3(news)

	select {
	case res := <-news:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	go server4(image)
	go server5(image)
	go server6(image)
	select {
	case res := <-image:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 2")
	}
	go server7(video)
	go server8(video)
	go server9(video)
	select {
	case res := <-video:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 3")
	}

	// for i := 0; i < 16; i++ {
	// 	new <- i
	// 	image <- i + 4
	// 	video <- i + 5
	// }

}
