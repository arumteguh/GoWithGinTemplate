package employees

type EmployeeService interface {
	GetEmployees() ([]employee, error)
	GetEmployee(employeeID string) (*employee, error)
	CreateEmployee(emp *employee) error
	UpdateEmployee(employeeID string, emp *employee) error
	DeleteEmployee(employeeID string) error
}

type DefaultEmployeeService struct {
	repo EmployeeRepository
}

func NewDefaultEmployeeService(repo EmployeeRepository) EmployeeService {
	return &DefaultEmployeeService{
		repo: repo,
	}
}

func (es *DefaultEmployeeService) GetEmployees() ([]employee, error) {
	return es.repo.GetEmployees()
}

func (es *DefaultEmployeeService) GetEmployee(employeeID string) (*employee, error) {
	return es.repo.GetEmployee(employeeID)
}

func (es *DefaultEmployeeService) CreateEmployee(emp *employee) error {
	
	
	return es.repo.CreateEmployee(emp)
}

func (es *DefaultEmployeeService) UpdateEmployee(employeeID string, emp *employee) error {
	return es.repo.UpdateEmployee(employeeID, emp)
}

func (es *DefaultEmployeeService) DeleteEmployee(employeeID string) error {
	return es.repo.DeleteEmployee(employeeID)
}
