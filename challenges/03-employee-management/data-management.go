package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

func (m *Manager) AddEmployee(e Employee) {
	m.Employees = append(m.Employees, e)
}

func (m *Manager) RemoveEmployee(id int) {
	for i, e := range m.Employees {
		if e.ID == id {
			m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
			return
		}
	}
}

func (m *Manager) GetAverageSalary() float64 {
	var total float64
	for _, e := range m.Employees {
		total += e.Salary
	}
	return total / float64(len(m.Employees))
}

func (m *Manager) FindEmployeeByID(id int) *Employee {
	for _, e := range m.Employees {
		if e.ID == id {
			return &e
		}
	}
	return nil
}

func main() {
	manager := Manager{}

	manager.AddEmployee(Employee{1, "Arpit", 22, 50000})
	manager.AddEmployee(Employee{2, "Rahul", 25, 60000})
	manager.AddEmployee(Employee{3, "Neha", 24, 70000})

	fmt.Println("Average Salary:", manager.GetAverageSalary())

	emp := manager.FindEmployeeByID(2)
	if emp != nil {
		fmt.Println("Found:", emp.Name)
	}

	manager.RemoveEmployee(2)

	fmt.Println("After Removal:")
	for _, emp := range manager.Employees {
		fmt.Println(emp.ID, emp.Name)
	}
}
