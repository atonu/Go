package staff

var OverPaidLimit = 70000
var underPaidLimit = 40000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) OverPaid() []Employee {
	var overPaid []Employee

	for _, i := range e.AllStaff {
		if i.Salary > OverPaidLimit {
			overPaid = append(overPaid, i)
		}
	}
	return overPaid
}

func (e *Office) UnderPaid() []Employee {
	var UnderPaid []Employee

	for _, i := range e.AllStaff {
		if i.Salary < underPaidLimit {
			UnderPaid = append(UnderPaid, i)
		}
	}
	return UnderPaid
}
