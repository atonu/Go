package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var read *bufio.Reader

func main() {
	read = bufio.NewReader(os.Stdin)
	userName := ReadString("What's your name?")
	fmt.Println("Your name is", userName)
}

func promt() {
	fmt.Print("->")
}

func ReadString(s string) string {
	fmt.Println(s)
	promt()
	userName, _ := read.ReadString('\n')
	userName = strings.Replace(userName, "\r\n", "", -1)
	userName = strings.Replace(userName, "\n", "", -1)

	return userName
}
