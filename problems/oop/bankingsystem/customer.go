/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package bankingsystem

import (
	"errors"
	"time"
)

// Customer represents a bank customer
type Customer struct {
	ID             string
	FirstName      string
	LastName       string
	Address        string
	JoinDate       time.Time
	AccountNumbers []string
}

// NewCustomer creates a new customer with validation
func NewCustomer(id, firstName, lastName, address string) (*Customer, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("first name and last name are required")
	}

	if address == "" {
		return nil, errors.New("address is required")
	}

	return &Customer{
		ID:             id,
		FirstName:      firstName,
		LastName:       lastName,
		Address:        address,
		JoinDate:       time.Now(),
		AccountNumbers: []string{},
	}, nil
}

// AddAccount adds an account to the customer
func (c *Customer) AddAccount(accountNumber string) {
	c.AccountNumbers = append(c.AccountNumbers, accountNumber)
}

// GetFullName returns the customer's full name
func (c *Customer) GetFullName() string {
	return c.FirstName + " " + c.LastName
}
