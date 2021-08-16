package packageone

import "fmt"

var PackageVar = 3

func PrintMe(var1, var2 int) {
	fmt.Print(var1,var2,PackageVar)
}