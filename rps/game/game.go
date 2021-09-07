package game

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

var reader = bufio.NewReader(os.Stdin)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan  chan string
	RoundChannel chan int
	Round        Round
}
type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

func (g *Game) Rounds() {
	for {
		select {
		case round := <-g.RoundChannel:
			g.Round.RoundNumber += round
			g.RoundChannel <- 0
		case print := <-g.DisplayChan:
			fmt.Print(print)
			g.DisplayChan <- ""
		}
	}
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	g.DisplayChan <- fmt.Sprintf(`
Round: %d
---------
`, g.Round.RoundNumber)
	<-g.DisplayChan

	g.Instruction()
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\r", "", -1)
	playerValue := -1
	computerValue := rand.Intn(2)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()

	g.DisplayChan <- fmt.Sprintf("Player chosen %s\n", playerChoice)
	<-g.DisplayChan

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose ROCK"
		<-g.DisplayChan
	case PAPER:
		g.DisplayChan <- "Computer chose PAPER"
		<-g.DisplayChan
	case SCISSORS:
		g.DisplayChan <- "Computer chose SCISSORS"
		<-g.DisplayChan
	default:
	}

	fmt.Println()

	if playerValue == computerValue {
		g.Draw()
		return false
	} else {
		if playerValue == -1 {
			g.DisplayChan <- "Invalid!"
			<-g.DisplayChan
			return false
		} else if playerValue == (computerValue+1)%3 {
			g.PlayerWins()
		} else {
			g.ComputerWins()
		}

		// switch playerValue {
		// case ROCK:
		// 	if computerValue == PAPER {
		// 		g.ComputerWins()
		// 	} else {
		// 		g.PlayerWins()
		// 	}
		// case PAPER:

		// 	if computerValue == SCISSORS {
		// 		g.ComputerWins()
		// 	} else {
		// 		g.PlayerWins()
		// 	}
		// case SCISSORS:

		// 	if computerValue == ROCK {
		// 		g.ComputerWins()
		// 	} else {
		// 		g.PlayerWins()
		// 	}
		// default:
		// 	g.DisplayChan <- "Invalid!"
		// 	<-g.DisplayChan
		// 	return false
		// }
	}
	fmt.Println()
	return true
}

func (g *Game) Instruction() {
	g.DisplayChan <- "Please enter rock, paper, or scissors ->"
	<-g.DisplayChan
}
func (g *Game) ClearScreen() {
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

func (g *Game) Results() {
	fmt.Println()
	g.DisplayChan <- fmt.Sprintf(`
--- --- ---
Final score: Player: %d and Computer %d
`, g.Round.PlayerScore, g.Round.ComputerScore)
	<-g.DisplayChan

	if g.Round.ComputerScore > g.Round.PlayerScore {
		g.ComputerWins()
	} else {
		g.PlayerWins()
	}
}
func (g *Game) PrintIntro() {
	g.DisplayChan <- `
Rock, Paper & Scissors
--- --- ---
Game is played for three rounds, and best of three wins the game. Good luck!
`
	<-g.DisplayChan
}
func (g *Game) PlayerWins() {
	g.DisplayChan <- "Player wins!\n\n"
	<-g.DisplayChan
	g.Round.PlayerScore++
}

func (g *Game) ComputerWins() {
	g.DisplayChan <- "Computer wins!\n\n"
	<-g.DisplayChan
	g.Round.ComputerScore++
}

func (g *Game) Draw() {
	g.DisplayChan <- "It's a draw\n\n"
	<-g.DisplayChan
}
