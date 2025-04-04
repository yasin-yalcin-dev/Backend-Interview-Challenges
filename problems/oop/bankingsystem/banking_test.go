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
	"testing"
)

func TestCustomerCreation(t *testing.T) {
	customer, err := NewCustomer("C1001", "John", "Doe", "123 Go Street")
	if err != nil {
		t.Errorf("Expected no error creating valid customer, got: %v", err)
	}

	if customer.GetFullName() != "John Doe" {
		t.Errorf("Expected full name 'John Doe', got: %s", customer.GetFullName())
	}

	// Test invalid customer
	_, err = NewCustomer("C1002", "", "Doe", "123 Go Street")
	if err == nil {
		t.Errorf("Expected error for empty first name, got none")
	}
}

func TestBankAndAccountOperations(t *testing.T) {
	bank := NewBank("Test Bank")

	// Create customer
	customer, err := bank.CreateCustomer("Jane", "Smith", "456 Go Lane")
	if err != nil {
		t.Errorf("Expected no error creating customer, got: %v", err)
	}

	// Test checking account
	checking, err := bank.OpenCheckingAccount(customer.ID, 1000.0)
	if err != nil {
		t.Errorf("Expected no error opening checking account, got: %v", err)
	}

	if checking.GetBalance() != 1000.0 {
		t.Errorf("Expected balance of 1000.0, got: %.2f", checking.GetBalance())
	}

	// Test deposit
	err = checking.Deposit(500.0)
	if err != nil {
		t.Errorf("Expected no error on deposit, got: %v", err)
	}

	if checking.GetBalance() != 1500.0 {
		t.Errorf("Expected balance of 1500.0 after deposit, got: %.2f", checking.GetBalance())
	}

	// Test withdrawal with fee
	err = checking.Withdraw(200.0)
	if err != nil {
		t.Errorf("Expected no error on withdrawal, got: %v", err)
	}

	expectedBalance := 1500.0 - 200.0 - checking.TransactionFee
	if checking.GetBalance() != expectedBalance {
		t.Errorf("Expected balance of %.2f after withdrawal and fee, got: %.2f",
			expectedBalance, checking.GetBalance())
	}

	// Test insufficient funds
	err = checking.Withdraw(2000.0)
	if err == nil {
		t.Errorf("Expected error on withdrawal with insufficient funds, got none")
	}

	// Test savings account
	savings, err := bank.OpenSavingsAccount(customer.ID, 2000.0)
	if err != nil {
		t.Errorf("Expected no error opening savings account, got: %v", err)
	}

	// Test interest
	initialBalance := savings.GetBalance()
	savings.ApplyMonthlyInterest()
	expectedWithInterest := initialBalance * (1.0 + savings.InterestRate)

	if savings.GetBalance() != expectedWithInterest {
		t.Errorf("Expected balance of %.2f after interest, got: %.2f",
			expectedWithInterest, savings.GetBalance())
	}

	// Test withdrawal limits
	for i := 0; i < savings.MaxWithdrawals; i++ {
		err = savings.Withdraw(10.0)
		if err != nil {
			t.Errorf("Expected no error on withdrawal #%d, got: %v", i+1, err)
		}
	}

	// This should exceed the limit
	err = savings.Withdraw(10.0)
	if err == nil {
		t.Errorf("Expected error when exceeding withdrawal limit, got none")
	}

	// Test customer accounts retrieval
	accounts, err := bank.GetCustomerAccounts(customer.ID)
	if err != nil {
		t.Errorf("Expected no error getting customer accounts, got: %v", err)
	}

	if len(accounts) != 2 {
		t.Errorf("Expected 2 accounts for customer, got: %d", len(accounts))
	}
}
