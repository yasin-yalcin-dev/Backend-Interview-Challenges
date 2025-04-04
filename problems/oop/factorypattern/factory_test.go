/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package factorypattern

import (
	"testing"
)

func TestPaymentFactory(t *testing.T) {
	factory := NewPaymentFactory()

	// Test credit card payment creation
	t.Run("Create valid credit card payment", func(t *testing.T) {
		params := map[string]interface{}{
			"cardNumber": "4111-1111-1111-1111",
			"expiryDate": "12/2025",
			"cvv":        "123",
			"cardHolder": "John Doe",
			"amount":     100.0,
		}

		payment, err := factory.CreatePayment("creditcard", params)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if payment == nil {
			t.Fatal("Expected payment to be created, got nil")
		}

		if payment.GetAmount() != 100.0 {
			t.Errorf("Expected amount 100.0, got: %.2f", payment.GetAmount())
		}
	})

	// Test invalid credit card
	t.Run("Create invalid credit card payment", func(t *testing.T) {
		params := map[string]interface{}{
			"cardNumber": "invalid",
			"expiryDate": "12/2025",
			"cvv":        "123",
			"amount":     100.0,
		}

		_, err := factory.CreatePayment("creditcard", params)
		if err == nil {
			t.Error("Expected error for invalid card number, got none")
		}
	})

	// Test PayPal payment creation
	t.Run("Create valid PayPal payment", func(t *testing.T) {
		params := map[string]interface{}{
			"email":  "user@example.com",
			"amount": 75.50,
		}

		payment, err := factory.CreatePayment("paypal", params)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if payment == nil {
			t.Fatal("Expected payment to be created, got nil")
		}

		if payment.GetAmount() != 75.50 {
			t.Errorf("Expected amount 75.50, got: %.2f", payment.GetAmount())
		}
	})

	// Test invalid PayPal
	t.Run("Create invalid PayPal payment", func(t *testing.T) {
		params := map[string]interface{}{
			"email":  "invalid-email",
			"amount": 75.50,
		}

		_, err := factory.CreatePayment("paypal", params)
		if err == nil {
			t.Error("Expected error for invalid email, got none")
		}
	})

	// Test bank transfer payment creation
	t.Run("Create valid bank transfer payment", func(t *testing.T) {
		params := map[string]interface{}{
			"accountName":   "Jane Smith",
			"accountNumber": "123456789",
			"bankCode":      "ABCDEF",
			"amount":        500.0,
		}

		payment, err := factory.CreatePayment("banktransfer", params)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if payment == nil {
			t.Fatal("Expected payment to be created, got nil")
		}

		if payment.GetAmount() != 500.0 {
			t.Errorf("Expected amount 500.0, got: %.2f", payment.GetAmount())
		}
	})

	// Test unknown payment type
	t.Run("Create unknown payment type", func(t *testing.T) {
		_, err := factory.CreatePayment("bitcoin", map[string]interface{}{
			"amount": 42.0,
		})

		if err == nil {
			t.Error("Expected error for unknown payment type, got none")
		}
	})
}
