# Banking System

## Problem
Design and implement a simple banking system using object-oriented principles in Go. The system should support different account types (Checking, Savings), allow for deposits and withdrawals, and apply appropriate fees and interest rates based on account type.

## Requirements
1. Create an Account interface with Deposit(), Withdraw(), and GetBalance() methods
2. Implement concrete account types: CheckingAccount and SavingsAccount
3. Create a Customer struct with personal information and account references
4. Create a Bank struct to manage customers and accounts
5. Implement proper validation and error handling
6. Apply different interest rates and fees based on account type:
   - Checking accounts have transaction fees but no interest
   - Savings accounts have interest but withdrawal limits

## Examples
```go
// Create a bank
bank := bankingsystem.NewBank("Go Banking")

// Add customers with accounts
customer, _ := bank.CreateCustomer("John", "Doe", "123 Go Street")
checkingAccount, _ := bank.OpenCheckingAccount(customer.ID, 1000.0)
savingsAccount, _ := bank.OpenSavingsAccount(customer.ID, 5000.0)

// Perform operations
checkingAccount.Deposit(500.0)
savingsAccount.Withdraw(1000.0)
checkingAccount.Withdraw(200.0)

// Check balances
fmt.Printf("Checking balance: $%.2f\n", checkingAccount.GetBalance())
fmt.Printf("Savings balance: $%.2f\n", savingsAccount.GetBalance())

// Calculate monthly interest (for savings account)
savingsAccount.ApplyMonthlyInterest()
fmt.Printf("Savings balance after interest: $%.2f\n", savingsAccount.GetBalance())

Expected Output
CopyChecking balance: $1295.00  // $1000 + $500 - $200 - $5 transaction fee
Savings balance: $4000.00   // $5000 - $1000
Savings balance after interest: $4020.00  // $4000 + 0.5% monthly interest