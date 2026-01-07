package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name string `valid:"required~Name is required"`
	Salary float64 `valid:"range(15000|200000)~Salary must be between 15000 and 200000"`
	EmployeeCode string `valid:"required~EmployeeCode is required"`
}

func (e *Employee) Validation() error {
	ok, err := govalidator.ValidateStruct(e)
	if !ok {
		return err
	}
	return nil
}