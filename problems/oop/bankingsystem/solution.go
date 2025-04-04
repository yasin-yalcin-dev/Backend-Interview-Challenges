/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package bankingsystem

import "fmt"

// RunExample demonstrates the banking system implementation
func RunExample() {
	// Create a bank
	bank := NewBank("Go Banking")

	// Add customers with accounts
	customer, err := bank.CreateCustomer("John", "Doe", "123 Go Street")
	if err != nil {
		fmt.Printf("Error creating customer: %v\n", err)
		return
	}

	checkingAccount, err := bank.OpenCheckingAccount(customer.ID, 1000.0)
	if err != nil {
		fmt.Printf("Error creating checking account: %v\n", err)
		return
	}

	savingsAccount, err := bank.OpenSavingsAccount(customer.ID, 5000.0)
	if err != nil {
		fmt.Printf("Error creating savings account: %v\n", err)
		return
	}

	fmt.Println("Banking System Example:")
	fmt.Println("------------------------")
	fmt.Printf("Customer: %s (ID: %s)\n", customer.GetFullName(), customer.ID)
	fmt.Printf("Initial Checking Balance: $%.2f\n", checkingAccount.GetBalance())
	fmt.Printf("Initial Savings Balance: $%.2f\n\n", savingsAccount.GetBalance())

	// Perform operations
	fmt.Println("Performing transactions:")
	fmt.Println("- Depositing $500 to Checking")
	checkingAccount.Deposit(500.0)

	fmt.Println("- Withdrawing $1000 from Savings")
	if err := savingsAccount.Withdraw(1000.0); err != nil {
		fmt.Printf("  Error: %v\n", err)
	}

	fmt.Println("- Withdrawing $200 from Checking (plus $5 fee)")
	if err := checkingAccount.Withdraw(200.0); err != nil {
		fmt.Printf("  Error: %v\n", err)
	}

	// Show updated balances
	fmt.Printf("\nUpdated Balances:\n")
	fmt.Printf("Checking balance: $%.2f\n", checkingAccount.GetBalance())
	fmt.Printf("Savings balance: $%.2f\n", savingsAccount.GetBalance())

	// Apply monthly interest
	fmt.Println("\nApplying monthly interest to Savings account...")
	savingsAccount.ApplyMonthlyInterest()
	fmt.Printf("Savings balance after interest: $%.2f\n", savingsAccount.GetBalance())

	// Try exceeding withdrawal limits
	fmt.Println("\nTesting withdrawal limits on Savings account:")
	for i := 1; i <= 7; i++ {
		err := savingsAccount.Withdraw(10.0)
		if err != nil {
			fmt.Printf("Withdrawal #%d: Error - %v\n", i, err)
			break
		}
		fmt.Printf("Withdrawal #%d: $10 successfully withdrawn. New balance: $%.2f\n",
			i, savingsAccount.GetBalance())
	}
}
