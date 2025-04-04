/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package factorypattern

import "fmt"

// RunExample demonstrates the factory pattern implementation
func RunExample() {
	fmt.Println("Factory Pattern Example:")
	fmt.Println("------------------------")

	// Create a payment factory
	paymentFactory := NewPaymentFactory()

	// Create and process a credit card payment
	creditCardParams := map[string]interface{}{
		"cardNumber": "4111-1111-1111-1111",
		"expiryDate": "12/2025",
		"cvv":        "123",
		"cardHolder": "John Doe",
		"amount":     100.0,
	}

	fmt.Println("Creating a credit card payment...")
	creditCardPayment, err := paymentFactory.CreatePayment("creditcard", creditCardParams)
	if err != nil {
		fmt.Printf("Error creating credit card payment: %v\n", err)
	} else {
		creditCardPayment.Process()
	}

	// Create and process a PayPal payment
	payPalParams := map[string]interface{}{
		"email":  "user@example.com",
		"amount": 75.50,
	}

	fmt.Println("\nCreating a PayPal payment...")
	payPalPayment, err := paymentFactory.CreatePayment("paypal", payPalParams)
	if err != nil {
		fmt.Printf("Error creating PayPal payment: %v\n", err)
	} else {
		payPalPayment.Process()
	}

	// Create and process a bank transfer payment
	bankTransferParams := map[string]interface{}{
		"accountName":   "Jane Smith",
		"accountNumber": "123456789",
		"bankCode":      "ABCDEF",
		"amount":        500.0,
	}

	fmt.Println("\nCreating a bank transfer payment...")
	bankTransferPayment, err := paymentFactory.CreatePayment("banktransfer", bankTransferParams)
	if err != nil {
		fmt.Printf("Error creating bank transfer payment: %v\n", err)
	} else {
		bankTransferPayment.Process()
	}

	// Try creating an invalid payment type
	fmt.Println("\nTrying to create an invalid payment type...")
	_, err = paymentFactory.CreatePayment("bitcoin", map[string]interface{}{
		"amount": 42.0,
	})
	fmt.Printf("Result: %v\n", err)

	// Try creating a payment with invalid parameters
	fmt.Println("\nTrying to create a payment with invalid parameters...")
	_, err = paymentFactory.CreatePayment("creditcard", map[string]interface{}{
		"cardNumber": "invalid",
		"amount":     100.0,
	})
	fmt.Printf("Result: %v\n", err)
}
