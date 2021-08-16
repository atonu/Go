package main

import (
	"myapp/packageone"
)

var myVar = 1
func main() {
	blockVar := 2
	packageone.PrintMe(myVar,blockVar)
}