package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

type User struct {
	userName string
	age      int
	favNum   float64
	OwnsADog bool
}

var read *bufio.Reader

func main() {
	read = bufio.NewReader(os.Stdin)
	var user User
	user.userName = ReadString("What's your name?")
	user.age = ReadInt("How old r u?")
	user.favNum = ReadFloat("Favourite Number?")
	user.OwnsADog = ReadBool("Do you own a dog? (y/n)")

	//other ways
	// fmt.Println("Your name is", userName, "and you are", age, "years old")
	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", userName, age))
	fmt.Printf("Your name is %s. You are %d years old. Your favt number is %.2f and own a dog: %t.\n", user.userName, user.age, user.favNum, user.OwnsADog)
}

func promt() {
	fmt.Print("->")
}

func ReadString(s string) string {
	for {
		fmt.Println(s)
		promt()
		userInput, _ := read.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("fam pls")
		} else {
			return userInput
		}
	}
}

func ReadInt(s string) int {
	fmt.Println(s)
	promt()
	userInput, _ := read.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	num, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Enter number pls")
		return ReadInt(s)
	}
	return num
}

func ReadFloat(s string) float64 {
	fmt.Println(s)
	promt()
	userInput, _ := read.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	num, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		fmt.Println("Enter number pls")
		return ReadFloat(s)
	}
	return num
}

func ReadBool(s string) bool {
	fmt.Println(s)
	promt()
	keyboard.Open()
	Rune, _, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}

	defer keyboard.Close()

	key := fmt.Sprintf("%c", Rune)
	if key == "y" || key == "Y" {
		return true
	} else if key == "n" || key == "N" {
		return false
	} else {
		fmt.Println("y/n pls")
		return ReadBool(s)
	}
}
