package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	userName string
	age      int
	favNum   float64
	shitter  bool
}

var read *bufio.Reader

func main() {
	read = bufio.NewReader(os.Stdin)
	var user User
	user.userName = ReadString("What's your name?")
	user.age = ReadInt("How old r u?")
	user.favNum = ReadFloat("Favourite Number?")
	//other ways
	// fmt.Println("Your name is", userName, "and you are", age, "years old")
	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", userName, age))
	fmt.Printf("Your name is %s. You are %d years old. Your favt number is %.2f.\n", user.userName, user.age, user.favNum)
}

func promt() {
	fmt.Print("->")
}

func ReadString(s string) string {
	for {
		fmt.Println(s)
		promt()
		userName, _ := read.ReadString('\n')
		userName = strings.Replace(userName, "\r\n", "", -1)
		userName = strings.Replace(userName, "\n", "", -1)
		if userName == "" {
			fmt.Println("fam pls")
		} else {
			return userName
		}
	}
}

func ReadInt(s string) int {
	fmt.Println(s)
	promt()
	userName, _ := read.ReadString('\n')
	userName = strings.Replace(userName, "\r\n", "", -1)
	userName = strings.Replace(userName, "\n", "", -1)

	num, err := strconv.Atoi(userName)
	if err != nil {
		fmt.Println("Enter number pls")
		return ReadInt(s)
	}
	return num
}

func ReadFloat(s string) float64 {
	fmt.Println(s)
	promt()
	userName, _ := read.ReadString('\n')
	userName = strings.Replace(userName, "\r\n", "", -1)
	userName = strings.Replace(userName, "\n", "", -1)

	num, err := strconv.ParseFloat(userName, 64)
	if err != nil {
		fmt.Println("Enter number pls")
		return ReadFloat(s)
	}
	return num
}
