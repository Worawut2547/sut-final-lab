package tests

import (
	"testing"
	"fmt"

	"sut-final-lab/entities"
	."github.com/onsi/gomega"
)

func TestSalaryEmployee(t *testing.T) {
	empInfo := entities.Employee{
		Name: "Worawut",
		Salary: 20000,
		EmployeeCode: "EMP6616",
	}

	fmt.Println("Case 2: Salary must be between 15000 and 200000")
	t.Run("Salary must be between 15000 and 200000", func(t *testing.T){
		g := NewGomegaWithT(t)
		obj := empInfo
		obj.Salary = 150
		err := obj.Validation()

		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error).ToNot(Equal("Salary must be between 15000 and 200000"))
	})
}