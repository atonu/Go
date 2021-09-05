package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 10
	b := 0

	c, error := devide(a, b)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(c)
	}

}

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}
