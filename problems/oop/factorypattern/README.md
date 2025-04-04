# Factory Pattern

## Problem
Implement the Factory design pattern to create different types of payment methods. The system should be able to create and process payments using Credit Card, PayPal, and Bank Transfer, with each payment method having its own specific properties and behaviors.

## Requirements
1. Create a Payment interface with common methods for all payment types
2. Implement concrete payment types: CreditCardPayment, PayPalPayment, BankTransferPayment
3. Create a PaymentFactory that generates the appropriate payment type based on input
4. Each payment type should have its own validation rules and processing behavior
5. Design the system to be easily extensible for new payment methods

## Examples
```go
// Create different payment methods using the factory
paymentFactory := factory.NewPaymentFactory()

// Create a credit card payment
creditCardPayment, err := paymentFactory.CreatePayment("creditcard", map[string]interface{}{
    "cardNumber": "4111-1111-1111-1111",
    "expiryDate": "12/2025",
    "cvv":        "123",
    "amount":     100.0,
})

// Create a PayPal payment
paypalPayment, err := paymentFactory.CreatePayment("paypal", map[string]interface{}{
    "email":  "user@example.com",
    "amount": 75.50,
})

// Process payments
creditCardPayment.Process()
paypalPayment.Process()
Expected Output
CopyProcessing Credit Card payment of $100.00 with card number: 4111-****-****-1111
Processing PayPal payment of $75.50 with email: user@example.com