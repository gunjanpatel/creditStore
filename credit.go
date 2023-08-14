package creditStore

import "fmt"

type Account struct {
	FirtName string
	LastName string
}

type Employee struct {
	Account
	Credits float64
}

func (account *Account) ChangeFirstName(name string) {
	account.FirtName = name
}

func (account *Account) ChangeLastName(name string) {
	account.LastName = name
}

func (employee *Employee) AddCredits(credit float64) {
	employee.Credits += credit
}

func (employee *Employee) RemoveCredits(credit float64) {
	employee.Credits -= credit
}

func (employee Employee) CheckCredits() string {
	return fmt.Sprintf("%.2f DKK", employee.Credits)
}

func (employee Employee) String() string {
	return fmt.Sprintf("%s %s", employee.FirtName, employee.LastName)
}
