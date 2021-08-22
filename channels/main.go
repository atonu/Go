package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)
	go listenForKeyPress()

	keyboard.Open()

	defer func() {
		keyboard.Close()
	}()

	for {
		key, _, _ := keyboard.GetSingleKey()

		if key == 'q' || key == 'Q' {
			break
		}
		keyPressChan <- key
	}

}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("u pressed", string(key))
	}
}
