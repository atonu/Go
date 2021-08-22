package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("My %s says %s\n", a.Name, a.Sound)
}
func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s Has %d legs\n", a.Name, a.NumberOfLegs)
}

func main() {
	// myTotal := sumMany(2, 3, 4, 5)
	// fmt.Println(myTotal)
	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meao",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()
}

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}
	return total
}
