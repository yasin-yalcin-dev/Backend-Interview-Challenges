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
	"regexp"
	"strings"
)

// Payment interface defines common methods for all payment types
type Payment interface {
	Validate() error
	Process() error
	GetAmount() float64
}

// BasePayment contains common fields and methods for all payment types
type BasePayment struct {
	Amount float64
}

// GetAmount returns the payment amount
func (p *BasePayment) GetAmount() float64 {
	return p.Amount
}

// CreditCardPayment represents a credit card payment method
type CreditCardPayment struct {
	BasePayment
	CardNumber string
	ExpiryDate string
	CVV        string
	CardHolder string
}

// Validate checks if the credit card payment information is valid
func (p *CreditCardPayment) Validate() error {
	if p.Amount <= 0 {
		return errors.New("payment amount must be positive")
	}

	// Basic credit card number validation (just checking format here)
	cardNumber := strings.ReplaceAll(p.CardNumber, "-", "")
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	if len(cardNumber) < 13 || len(cardNumber) > 19 {
		return errors.New("invalid card number length")
	}

	// Check if card number contains only digits
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(cardNumber) {
		return errors.New("card number must contain only digits")
	}

	// Basic expiry date validation (MM/YYYY format)
	reDate := regexp.MustCompile(`^(0[1-9]|1[0-2])\/20\d{2}$`)
	if !reDate.MatchString(p.ExpiryDate) {
		return errors.New("expiry date must be in MM/YYYY format")
	}

	// Basic CVV validation (3 or 4 digits)
	if len(p.CVV) < 3 || len(p.CVV) > 4 {
		return errors.New("CVV must be 3 or 4 digits")
	}

	return nil
}

// Process handles the credit card payment processing
func (p *CreditCardPayment) Process() error {
	// Mask the card number for display
	cardNumber := strings.ReplaceAll(p.CardNumber, "-", "")
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	maskedCard := cardNumber[:4] + "-****-****-" + cardNumber[len(cardNumber)-4:]

	fmt.Printf("Processing Credit Card payment of $%.2f with card number: %s\n",
		p.Amount, maskedCard)

	// In a real system, we would integrate with a payment gateway here
	return nil
}

// PayPalPayment represents a PayPal payment method
type PayPalPayment struct {
	BasePayment
	Email string
}

// Validate checks if the PayPal payment information is valid
func (p *PayPalPayment) Validate() error {
	if p.Amount <= 0 {
		return errors.New("payment amount must be positive")
	}

	// Basic email validation
	reEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !reEmail.MatchString(p.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

// Process handles the PayPal payment processing
func (p *PayPalPayment) Process() error {
	fmt.Printf("Processing PayPal payment of $%.2f with email: %s\n",
		p.Amount, p.Email)

	// In a real system, we would integrate with PayPal's API here
	return nil
}

// BankTransferPayment represents a bank transfer payment method
type BankTransferPayment struct {
	BasePayment
	AccountName   string
	AccountNumber string
	BankCode      string
}

// Validate checks if the bank transfer information is valid
func (p *BankTransferPayment) Validate() error {
	if p.Amount <= 0 {
		return errors.New("payment amount must be positive")
	}

	if p.AccountName == "" {
		return errors.New("account name is required")
	}

	// Basic account number validation
	if len(p.AccountNumber) < 5 || len(p.AccountNumber) > 20 {
		return errors.New("invalid account number length")
	}

	// Basic bank code validation
	if len(p.BankCode) < 3 || len(p.BankCode) > 11 {
		return errors.New("invalid bank code length")
	}

	return nil
}

// Process handles the bank transfer payment processing
func (p *BankTransferPayment) Process() error {
	fmt.Printf("Processing Bank Transfer of $%.2f to account: %s (Account number: %s, Bank code: %s)\n",
		p.Amount, p.AccountName, p.AccountNumber, p.BankCode)

	// In a real system, we would integrate with a banking API here
	return nil
}
