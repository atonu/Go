package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "fish")
	animals = append(animals, "dog")
	animals = append(animals, "horse")
	animals = append(animals, "cat")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}
	fmt.Println("First 2 animals are:", animals[0:2])
	fmt.Println("slice is", len(animals), "elements long")
	fmt.Println("is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("is it sorted?", sort.StringsAreSorted(animals))

	fmt.Println(animals)
	animals = DeleteFromSlice(animals, 1)
	sort.Strings(animals)
	fmt.Println(animals)
}

func DeleteFromSlice(str []string, i int) []string {

	str[i] = str[len(str)-1]
	str = str[:len(str)-1]
	return str
}
