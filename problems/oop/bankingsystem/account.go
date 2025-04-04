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

// Account interface defines methods that all account types must implement
type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
	GetAccountNumber() string
	GetAccountType() string
}

// BaseAccount contains common fields and methods for all account types
type BaseAccount struct {
	AccountNumber string
	Balance       float64
	OpenDate      time.Time
	OwnerID       string
}

// GetBalance returns the current balance
func (a *BaseAccount) GetBalance() float64 {
	return a.Balance
}

// GetAccountNumber returns the account number
func (a *BaseAccount) GetAccountNumber() string {
	return a.AccountNumber
}

// Deposit adds funds to the account
func (a *BaseAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	a.Balance += amount
	return nil
}

// CheckingAccount represents a checking account
type CheckingAccount struct {
	BaseAccount
	TransactionFee float64
}

// GetAccountType returns the account type
func (c *CheckingAccount) GetAccountType() string {
	return "Checking"
}

// Withdraw removes funds from the checking account and applies transaction fee
func (c *CheckingAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	totalWithdrawal := amount + c.TransactionFee

	if totalWithdrawal > c.Balance {
		return errors.New("insufficient funds for withdrawal and fee")
	}

	c.Balance -= totalWithdrawal
	return nil
}

// SavingsAccount represents a savings account
type SavingsAccount struct {
	BaseAccount
	InterestRate         float64
	WithdrawalsThisMonth int
	MaxWithdrawals       int
}

// GetAccountType returns the account type
func (s *SavingsAccount) GetAccountType() string {
	return "Savings"
}

// Withdraw removes funds from the savings account with withdrawal limits
func (s *SavingsAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	if amount > s.Balance {
		return errors.New("insufficient funds")
	}

	if s.WithdrawalsThisMonth >= s.MaxWithdrawals {
		return errors.New("maximum withdrawals for this month reached")
	}

	s.Balance -= amount
	s.WithdrawalsThisMonth++
	return nil
}

// ApplyMonthlyInterest applies interest to the savings account
func (s *SavingsAccount) ApplyMonthlyInterest() {
	interest := s.Balance * s.InterestRate
	s.Balance += interest
	// Reset monthly withdrawal counter
	s.WithdrawalsThisMonth = 0
}
