package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0
	PAPER        = 1
	SCISSORS     = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

var messages = [][]string{
	{"Good job!", "Congrats!", "Cool!"},
	{"Gotcha!", "Take a Bow!", "I > U"},
	{"Heck!", "We think alike", "We're the same"},
}

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	Result         string `json:"round_result"`
}

func PlayRound(playerValue int) Round {

	rand.Seed(time.Now().UnixNano())

	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw"
		winner = messages[2][rand.Intn(len(messages[2]))]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player Wins!"
		winner = messages[0][rand.Intn(len(messages[0]))]

	} else {
		roundResult = "ComputerWins"
		winner = messages[1][rand.Intn(len(messages[1]))]
	}

	round := Round{
		Message:        winner,
		ComputerChoice: computerChoice,
		Result:         roundResult,
	}
	return round

}
