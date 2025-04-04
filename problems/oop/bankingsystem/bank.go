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
	"fmt"
	"time"
)

// Bank represents the banking institution
type Bank struct {
	Name                  string
	Customers             map[string]*Customer
	Accounts              map[string]Account
	NextCustomerID        int
	NextAccountNumber     int
	CheckingFeeRate       float64
	SavingsInterestRate   float64
	MaxSavingsWithdrawals int
}

// NewBank creates a new bank instance
func NewBank(name string) *Bank {
	return &Bank{
		Name:                  name,
		Customers:             make(map[string]*Customer),
		Accounts:              make(map[string]Account),
		NextCustomerID:        1000,
		NextAccountNumber:     10000,
		CheckingFeeRate:       5.0,   // $5 transaction fee
		SavingsInterestRate:   0.005, // 0.5% monthly interest
		MaxSavingsWithdrawals: 6,     // 6 withdrawals per month
	}
}

// CreateCustomer creates a new customer at the bank
func (b *Bank) CreateCustomer(firstName, lastName, address string) (*Customer, error) {
	id := fmt.Sprintf("C%d", b.NextCustomerID)
	b.NextCustomerID++

	customer, err := NewCustomer(id, firstName, lastName, address)
	if err != nil {
		return nil, err
	}

	b.Customers[id] = customer
	return customer, nil
}

// OpenCheckingAccount opens a new checking account for a customer
func (b *Bank) OpenCheckingAccount(customerID string, initialDeposit float64) (*CheckingAccount, error) {
	if initialDeposit < 0 {
		return nil, errors.New("initial deposit cannot be negative")
	}

	customer, exists := b.Customers[customerID]
	if !exists {
		return nil, errors.New("customer not found")
	}

	accountNumber := fmt.Sprintf("AC%d", b.NextAccountNumber)
	b.NextAccountNumber++

	account := &CheckingAccount{
		BaseAccount: BaseAccount{
			AccountNumber: accountNumber,
			Balance:       initialDeposit,
			OpenDate:      time.Now(),
			OwnerID:       customerID,
		},
		TransactionFee: b.CheckingFeeRate,
	}

	b.Accounts[accountNumber] = account
	customer.AddAccount(accountNumber)

	return account, nil
}

// OpenSavingsAccount opens a new savings account for a customer
func (b *Bank) OpenSavingsAccount(customerID string, initialDeposit float64) (*SavingsAccount, error) {
	if initialDeposit < 0 {
		return nil, errors.New("initial deposit cannot be negative")
	}

	customer, exists := b.Customers[customerID]
	if !exists {
		return nil, errors.New("customer not found")
	}

	accountNumber := fmt.Sprintf("AS%d", b.NextAccountNumber)
	b.NextAccountNumber++

	account := &SavingsAccount{
		BaseAccount: BaseAccount{
			AccountNumber: accountNumber,
			Balance:       initialDeposit,
			OpenDate:      time.Now(),
			OwnerID:       customerID,
		},
		InterestRate:         b.SavingsInterestRate,
		WithdrawalsThisMonth: 0,
		MaxWithdrawals:       b.MaxSavingsWithdrawals,
	}

	b.Accounts[accountNumber] = account
	customer.AddAccount(accountNumber)

	return account, nil
}

// GetAccount retrieves an account by account number
func (b *Bank) GetAccount(accountNumber string) (Account, error) {
	account, exists := b.Accounts[accountNumber]
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}

// GetCustomer retrieves a customer by ID
func (b *Bank) GetCustomer(customerID string) (*Customer, error) {
	customer, exists := b.Customers[customerID]
	if !exists {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

// GetCustomerAccounts gets all accounts for a customer
func (b *Bank) GetCustomerAccounts(customerID string) ([]Account, error) {
	customer, err := b.GetCustomer(customerID)
	if err != nil {
		return nil, err
	}

	accounts := make([]Account, 0, len(customer.AccountNumbers))

	for _, accNum := range customer.AccountNumbers {
		if acc, exists := b.Accounts[accNum]; exists {
			accounts = append(accounts, acc)
		}
	}

	return accounts, nil
}

// ApplyMonthlyInterest applies interest to all savings accounts
func (b *Bank) ApplyMonthlyInterest() {
	for _, acc := range b.Accounts {
		if savingsAcc, ok := acc.(*SavingsAccount); ok {
			savingsAcc.ApplyMonthlyInterest()
		}
	}
}
