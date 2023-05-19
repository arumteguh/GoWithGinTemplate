package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service EmployeeService
}

func NewEmployeeController(service EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: service,
	}
}

// GetEmployees godoc
// @Summary Get all employees
// @Description Get a list of all employees
// @Tags Employees
// @Produce json
// @Success 200 {array} employee
// @Router /employee [get]
func (ec *EmployeeController) GetEmployees(c *gin.Context) {
	employees, err := ec.service.GetEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// GetEmployee godoc
// @Summary Get an employee by ID
// @Description Get an employee by ID
// @Tags Employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} employee
// @Router /employee/{id} [get]
func (ec *EmployeeController) GetEmployee(c *gin.Context) {
	employeeID := c.Param("id")

	employee, err := ec.service.GetEmployee(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if employee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// CreateEmployee godoc
// @Summary Create a new employee
// @Description Create a new employee
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body employee true "employee object"
// @Success 200 {object} employee
// @Router /employee [post]
func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	var emp employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ec.service.CreateEmployee(&emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, emp)
}
