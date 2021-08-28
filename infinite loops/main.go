package main

import (
	"bufio"
	"fmt"
	"myapp/myLogger"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	fmt.Println("Write something")
	go myLogger.ListenForLog(ch)

	for i := 0; i < 5; i++ {
		fmt.Print("->")

		inp, _ := reader.ReadString('\n')
		ch <- inp
		time.Sleep(time.Second / 5)
	}
}
