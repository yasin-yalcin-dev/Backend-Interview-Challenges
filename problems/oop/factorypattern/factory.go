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
	"errors"
	"fmt"
	"strings"
)

// PaymentFactory creates payment objects based on the requested type
type PaymentFactory struct {
	// Could contain configuration or dependencies here
}

// NewPaymentFactory creates a new payment factory
func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{}
}

// CreatePayment creates a payment of the specified type with the given parameters
func (f *PaymentFactory) CreatePayment(paymentType string, params map[string]interface{}) (Payment, error) {
	paymentType = strings.ToLower(paymentType)

	switch paymentType {
	case "creditcard":
		return f.createCreditCardPayment(params)
	case "paypal":
		return f.createPayPalPayment(params)
	case "banktransfer":
		return f.createBankTransferPayment(params)
	default:
		return nil, fmt.Errorf("unknown payment type: %s", paymentType)
	}
}

// createCreditCardPayment creates a credit card payment with parameters
func (f *PaymentFactory) createCreditCardPayment(params map[string]interface{}) (Payment, error) {
	payment := &CreditCardPayment{
		BasePayment: BasePayment{},
	}

	// Extract required parameters
	for k, v := range params {
		switch k {
		case "amount":
			if amount, ok := v.(float64); ok {
				payment.Amount = amount
			} else {
				return nil, errors.New("amount must be a float64")
			}
		case "cardNumber":
			if cardNumber, ok := v.(string); ok {
				payment.CardNumber = cardNumber
			} else {
				return nil, errors.New("cardNumber must be a string")
			}
		case "expiryDate":
			if expiryDate, ok := v.(string); ok {
				payment.ExpiryDate = expiryDate
			} else {
				return nil, errors.New("expiryDate must be a string")
			}
		case "cvv":
			if cvv, ok := v.(string); ok {
				payment.CVV = cvv
			} else {
				return nil, errors.New("cvv must be a string")
			}
		case "cardHolder":
			if cardHolder, ok := v.(string); ok {
				payment.CardHolder = cardHolder
			} else {
				return nil, errors.New("cardHolder must be a string")
			}
		}
	}

	// Validate the payment
	if err := payment.Validate(); err != nil {
		return nil, err
	}

	return payment, nil
}

// createPayPalPayment creates a PayPal payment with parameters
func (f *PaymentFactory) createPayPalPayment(params map[string]interface{}) (Payment, error) {
	payment := &PayPalPayment{
		BasePayment: BasePayment{},
	}

	// Extract required parameters
	for k, v := range params {
		switch k {
		case "amount":
			if amount, ok := v.(float64); ok {
				payment.Amount = amount
			} else {
				return nil, errors.New("amount must be a float64")
			}
		case "email":
			if email, ok := v.(string); ok {
				payment.Email = email
			} else {
				return nil, errors.New("email must be a string")
			}
		}
	}

	// Validate the payment
	if err := payment.Validate(); err != nil {
		return nil, err
	}

	return payment, nil
}

// createBankTransferPayment creates a bank transfer payment with parameters
func (f *PaymentFactory) createBankTransferPayment(params map[string]interface{}) (Payment, error) {
	payment := &BankTransferPayment{
		BasePayment: BasePayment{},
	}

	// Extract required parameters
	for k, v := range params {
		switch k {
		case "amount":
			if amount, ok := v.(float64); ok {
				payment.Amount = amount
			} else {
				return nil, errors.New("amount must be a float64")
			}
		case "accountName":
			if accountName, ok := v.(string); ok {
				payment.AccountName = accountName
			} else {
				return nil, errors.New("accountName must be a string")
			}
		case "accountNumber":
			if accountNumber, ok := v.(string); ok {
				payment.AccountNumber = accountNumber
			} else {
				return nil, errors.New("accountNumber must be a string")
			}
		case "bankCode":
			if bankCode, ok := v.(string); ok {
				payment.BankCode = bankCode
			} else {
				return nil, errors.New("bankCode must be a string")
			}
		}
	}

	// Validate the payment
	if err := payment.Validate(); err != nil {
		return nil, err
	}

	return payment, nil
}
