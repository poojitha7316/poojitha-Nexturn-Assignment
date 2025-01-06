package main

import (
	"fmt"
	"strings"
	"errors"
)

const (
	HR = "HR"
	IT = "IT"
)

type Employee struct {
	ID        int
	Name      string
	Age       int
	Department string
}

var employees []Employee

// Add Employee: Accept input and store in the array.
func addEmployee(id int, name string, age int, department string) error {
	// Check if the ID already exists
	for _, e := range employees {
		if e.ID == id {
			return errors.New("error: ID already exists")
		}
	}

	// Validate age
	if age <= 18 {
		return errors.New("error: Age must be greater than 18")
	}

	// Add the employee to the list
	employees = append(employees, Employee{ID: id, Name: name, Age: age, Department: department})
	fmt.Println("Employee added successfully!")

	return nil
}

// Search Employee: Search by ID or name
func searchEmployee(query string) (*Employee, error) {
	for _, e := range employees {
		if strings.EqualFold(e.Name, query) || fmt.Sprintf("%d", e.ID) == query {
			return &e, nil
		}
	}
	return nil, errors.New("error: employee not found")
}

// List Employees by Department
func listEmployeesByDepartment(department string) []Employee {
	var departmentEmployees []Employee
	for _, e := range employees {
		if strings.EqualFold(e.Department, department) {
			departmentEmployees = append(departmentEmployees, e)
		}
	}
	return departmentEmployees
}

// Count Employees by Department
func countEmployeesByDepartment(department string) int {
	count := 0
	for _, e := range employees {
		if strings.EqualFold(e.Department, department) {
			count++
		}
	}
	return count
}

func main() {
	// Add employees
	if err := addEmployee(1, "Alice", 25, HR); err != nil {
		fmt.Println(err)
	}
	if err := addEmployee(2, "Bob", 19, IT); err != nil {
		fmt.Println(err)
	}
	if err := addEmployee(1, "Charlie", 28,IT); err != nil {
		fmt.Println(err)
	}
	

	// Search for an employee
	employee, err := searchEmployee("Alice")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Employee found: ID: %d, Name: %s, Age: %d, Department: %s\n", employee.ID, employee.Name, employee.Age, employee.Department)
	}

	// List employees by department
	fmt.Println("\nEmployees in IT Department:")
	itEmployees := listEmployeesByDepartment(IT)
	for _, e := range itEmployees {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Department: %s\n", e.ID, e.Name, e.Age, e.Department)
	}

	// Count employees in a department
	fmt.Printf("\nNumber of employees in HR: %d\n", countEmployeesByDepartment(HR))
	fmt.Printf("Number of employees in IT: %d\n", countEmployeesByDepartment(IT))

}
