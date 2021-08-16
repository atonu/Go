package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil{
		log.Fatal(err)
	}
	defer func ()  {
		_ = keyboard.Close()
	}()

	fmt.Println("Press any key. Hit ESC")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil{
			log.Fatal(err)
		}
		if key != 0 {
			fmt.Println("u pressed", char,key)
			} else {
			fmt.Println("u pressed", char)
		}

		if key == keyboard.KeyEsc {
			break
		}

		fmt.Println("Byee....")
	}
}