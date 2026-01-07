package tests

import (
	"testing"
	"fmt"

	"sut-final-lab/entities"
	."github.com/onsi/gomega"
)

func TestEmployeeCodeEmployee(t *testing.T) {
	empInfo := entities.Employee{
		Name: "Worawut",
		Salary: 20000,
		EmployeeCode: "EMP6616",
	}

	fmt.Println("Case 3: EmployeeCode must be 2 uppercase English letters (A-Z) followed by and 4 digits (0-9)")
	t.Run("EmployeeCode must be 2 uppercase English letters (A-Z) followed by and 4 digits (0-9)", func(t *testing.T){
		g := NewGomegaWithT(t)
		obj := empInfo
		obj.EmployeeCode = ""
		err := obj.Validation()

		g.Expect(err).ToNot(BeNil())
	})
}