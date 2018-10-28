package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func GoRoutines() {
	go say("world")
	go say("hello")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	//GoRoutines()
	//UseChannels()
	bufferedChannel()
}

func UseChannels(){
	s := []int{7, 2, 8, -9, 4, 0}

	// create channel
	c := make(chan int)

	// sum half channels
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)


	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func bufferedChannel(){
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}


