package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan string)
var ch2 = make(chan string)

func task1() {
	time.Sleep(time.Second * 10)
	ch1 <- "one"
}

func task2() {
	time.Sleep(time.Second * 15)
	ch2 <- "two"
}

func main() {
	go task1()
	go task2()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
