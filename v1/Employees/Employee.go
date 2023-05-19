package employees

type employee struct {
	EmployeeID     string `json:"id"`
	EmployeeName   string `json:"name"`
	EmployeeGender string `json:"gender"`
	EmployeeCity   string `json:"city"`
}

func NewEmployee() *employee {
	return &employee{
		EmployeeID:     "",
		EmployeeName:   "",
		EmployeeGender: "",
		EmployeeCity:   "",
	}
}
