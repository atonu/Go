package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1
	computerValue := rand.Intn(2)
	reader := bufio.NewReader(os.Stdin)
	playerScore, computerScore := 0, 0

	clearScreen()

	for i := 0; i < 3; i++ {

		fmt.Print("Please enter rock, paper, or scissors ->")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)
		playerChoice = strings.Replace(playerChoice, "\r", "", -1)

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		} else {
			playerValue = -1
		}

		fmt.Println()

		fmt.Println("Player chose", playerChoice)

		switch computerValue {
		case 0:
			fmt.Println("Computer chose ROCK")
		case 1:
			fmt.Println("Computer chose PAPER")
		case 2:
			fmt.Println("Computer chose SCISSORS")
		default:
		}

		fmt.Println()

		if playerValue == computerValue {
			fmt.Println("It's a draw")
		} else {
			switch playerValue {

			case ROCK:

				if computerValue == PAPER {
					fmt.Println("Computer wins!")
					computerScore++
				} else {
					fmt.Println("Player wins!")
					playerScore++
				}
			case PAPER:

				if computerValue == SCISSORS {
					fmt.Println("Computer wins!")
					computerScore++
				} else {
					fmt.Println("Player wins!")
					playerScore++
				}
			case SCISSORS:

				if computerValue == ROCK {
					fmt.Println("Computer wins!")
					computerScore++
				} else {
					fmt.Println("Player wins!")
					playerScore++
				}

			default:
				fmt.Println("Invalid entry")
				i--
			}
		}
		fmt.Println()
	}

	fmt.Printf("Final score: Player: %d and Computer %d\n", playerScore, computerScore)
	if computerScore > playerScore {
		fmt.Println("Computer wins!")
	} else if computerScore < playerScore {
		fmt.Println("Player wins!")
	} else {
		fmt.Println("It's a draw")
	}

}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
