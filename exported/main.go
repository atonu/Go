package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Atonu", LastName: "Ahmed", Salary: 60000, FullTime: true},
	{FirstName: "Anika", LastName: "Tasnim", Salary: 60000, FullTime: true},
	{FirstName: "Nafees", LastName: "Farraz", Salary: 100000, FullTime: true},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	// log.Println(myStaff.All())
	// staff.OverPaidLimit = 60000
	log.Println(myStaff.OverPaid())
	log.Println(myStaff.UnderPaid())

}
