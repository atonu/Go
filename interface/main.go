package main

import "fmt"

type Animals interface {
	MakesSound() string
	HasLegs() int
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (c *Cat) MakesSound() string {
	return c.Sound
}
func (c *Cat) HasLegs() int {
	return c.NumberOfLegs
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) MakesSound() string {
	return d.Sound
}
func (d *Dog) HasLegs() int {
	return d.NumberOfLegs
}

func main() {
	var dog Dog
	dog.Name = "dog"
	dog.NumberOfLegs = 4
	dog.Sound = "woof"

	cat := Cat{
		Name:         "Cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	Riddle(&cat)
	Riddle(&dog)
}

func Riddle(a Animals) {
	fmt.Printf(`What animal makes "%s" noise and have %d legs?`, a.MakesSound(), a.HasLegs())
	fmt.Println()
}
