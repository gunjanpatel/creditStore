package creditStore

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Account
	credits float64
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.credits)
}

func CreateEmployee(firstName string, lastName string, credit float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credit}, nil
}

func (account *Account) ChangeFirstName(value string) {
	account.FirstName = value
}

func (account *Account) ChangeLastName(value string) {
	account.LastName = value
}

func (employee *Employee) AddCredits(amount float64) (float64, error) {
	if amount <= 0.0 {
		return employee.credits, errors.New("invalid credit amount")
	}

	employee.credits += amount

	return employee.credits, nil
}

func (employee *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount <= 0.0 {
		return 0.0, errors.New("you can't remove negative numbers")
	}

	if amount > employee.credits {
		return 0.0, errors.New("you can't remove more credits than the account has")
	}

	employee.credits -= amount

	return employee.credits, nil
}

func (employee Employee) CheckCredits() float64 {
	return employee.credits
}
