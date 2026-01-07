package tests

import (
	"testing"
	"fmt"

	"sut-final-lab/entities"
	."github.com/onsi/gomega"
)

func TestEmployeeValidation(t *testing.T) {
	empInfo := entities.Employee{
		Name: "Worawut",
		Salary: 20000,
		EmployeeCode: "EMP6616",
	}

	fmt.Println("Case 1: Data is collected")
	t.Run("Data is collected", func(t *testing.T){
		g := NewGomegaWithT(t)
		obj := empInfo
		err := obj.Validation()

		g.Expect(err).To(BeNil())
	})
}