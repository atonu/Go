package main

import "fmt"

type Employee struct {
	Name       string
	Age        int
	Salary     int
	isFullTime bool
}

func main() {
	jack := Employee{
		Name:       "Jack",
		Age:        25,
		Salary:     30000,
		isFullTime: false,
	}

	jill := Employee{
		Name:       "Jill",
		Age:        32,
		Salary:     55000,
		isFullTime: true,
	}

	fmt.Printf("%s is %d years old with %d salary\n", jack.Name, jack.Age, jack.Salary)
	fmt.Printf("%s is %d years old with %d salary\n", jill.Name, jill.Age, jill.Salary)

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, e := range employees {
		if e.Age > 30 {
			fmt.Println(e.Name, "is >30")
		} else {
			fmt.Println(e.Name, "is <30")
		}

		fmt.Printf("%s's salary is >50k: %t\n", e.Name, e.Salary > 50000)

		if e.Age > 30 && e.Salary > 50000 {
			fmt.Println(e.Name, "is >30 and salary is more than 50k")
		} else {
			fmt.Println(e.Name, "is'nt >30 or salary is not more than 50k")
		}

		if e.Age > 30 || e.Salary > 50000 {
			fmt.Println(e.Name, "is >30 or salary is more than 50k")
		} else {
			fmt.Println(e.Name, "is'nt >30 and salary is not more than 50k")
		}

		if (e.Age > 30 || e.Salary > 50000) && e.isFullTime {
			fmt.Println(e.Name, "matches this weird condition")
		} else {
			fmt.Println(e.Name, "doesn't match this weird condition")
		}
	}
}
