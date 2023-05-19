package employees

import (
	"database/sql"
	"log"
)

type EmployeeRepository interface {
	GetEmployees() ([]employee, error)
	GetEmployee(empolyeeID string) (*employee, error)
	CreateEmployee(emp *employee) error
	UpdateEmployee(employeeID string, emp *employee) error
	DeleteEmployee(employeeID string) error
}

type DefaultEmployeeRepository struct {
	db *sql.DB
}

func NewDefaultEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &DefaultEmployeeRepository{
		db: db,
	}
}

func (er *DefaultEmployeeRepository) GetEmployees() ([]employee, error) {
	query := "SELECT EmployeeID, EmployeeName, EmployeeGender, EmployeeCity FROM Employee"

	rows, err := er.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []employee
	for rows.Next() {
		var emp employee
		if err := rows.Scan(&emp.EmployeeID, &emp.EmployeeName, &emp.EmployeeGender, &emp.EmployeeCity); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (er *DefaultEmployeeRepository) GetEmployee(employeeID string) (*employee, error) {
	query := "SELECT EmployeeID, EmployeeName, EmployeeGender, EmployeeCity FROM Employee WHERE EmployeeID = @empID"

	stmt, err := er.db.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(sql.Named("empID", employeeID))

	var emp employee
	if err := row.Scan(&emp.EmployeeID, &emp.EmployeeName, &emp.EmployeeGender, &emp.EmployeeCity); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &emp, nil
}

func (er *DefaultEmployeeRepository) CreateEmployee(emp *employee) error {
	/*query := "INSERT INTO Employee (EmployeeID, EmployeeName, EmployeeGender, EmployeeCity) VALUES (?, ?, ?, ?)"

	_, err := er.db.Exec(query, emp.EmployeeID, emp.EmployeeName, emp.EmployeeGender, emp.EmployeeCity)
	if err != nil {
		return err
	}

	return nil*/

	query := "INSERT INTO Employee (EmployeeID, EmployeeName, EmployeeGender, EmployeeCity) VALUES (@EmployeeID, @EmployeeName, @EmployeeGender, @EmployeeCity)"

	_, err := er.db.Exec(query,
		sql.Named("EmployeeID", emp.EmployeeID),
		sql.Named("EmployeeName", emp.EmployeeName),
		sql.Named("EmployeeGender", emp.EmployeeGender),
		sql.Named("EmployeeCity", emp.EmployeeCity),
	)
	if err != nil {
		return err
	}

	return nil
}

func (er *DefaultEmployeeRepository) UpdateEmployee(employeeID string, emp *employee) error {
	query := "UPDATE Employee SET EmployeeName = @EmployeeName, EmployeeGender = @EmployeeGender, EmployeeCity = @EmployeeCity WHERE EmployeeID = @EmployeeID"

	_, err := er.db.Exec(query,
		sql.Named("EmployeeName", emp.EmployeeName),
		sql.Named("EmployeeGender", emp.EmployeeGender),
		sql.Named("EmployeeCity", emp.EmployeeCity),
		sql.Named("EmployeeID", employeeID),
	)

	if err != nil {
		return err
	}

	return nil
}

func (er *DefaultEmployeeRepository) DeleteEmployee(employeeID string) error {
	query := "DELETE FROM Employee WHERE EmployeeID = @EmployeeID"

	_, err := er.db.Exec(query,
		sql.Named("EmployeeID", employeeID),
	)

	if err != nil {
		return err
	}

	return nil
}
