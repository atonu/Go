package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)
const promt = "and hit enter when ready."
func main() {
	// seed
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8)+2
	var secondNumber = rand.Intn(8)+2
	var substraction =  rand.Intn(8)+2
	answer := firstNumber * secondNumber - substraction

	mathGame(firstNumber,secondNumber,substraction, answer)
}

func mathGame(firstNumber,secondNumber,substraction,answer int){
	reader := bufio.NewReader(os.Stdin)
	//display welcome
	fmt.Println("Guess the number Game")
	fmt.Println("__________________")
	fmt.Println("")
	fmt.Println("Think of a number 1-10 ",promt)
	reader.ReadString('\n')
	// take them through game
	fmt.Println("Multiply youe number by", firstNumber, promt)
	reader.ReadString('\n')
	fmt.Println("Now multiply the result by", secondNumber, promt)
	reader.ReadString('\n')
	fmt.Println("Devide by the original number you thought of",promt)
	reader.ReadString('\n')
	fmt.Println("Now substract", substraction, promt)
	reader.ReadString('\n')

	// give the answer
	fmt.Println("The answer is", answer)
}