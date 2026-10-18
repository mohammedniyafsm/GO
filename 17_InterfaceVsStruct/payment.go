package main

import (
	"fmt"
)

// ============================================
// 1. PAYMENT REQUEST
// ============================================

// This contains the information required to make a payment.
//
// In a real application, this data might come from an API request.

type PaymentRequest struct {
	OrderID  string
	Amount   float64
	Currency string
}

// ============================================
// 2. PAYMENT GATEWAY INTERFACE
// ============================================

// Every payment provider must have a Pay() method.
//
// Razorpay, Stripe, and PayPal will each implement this interface.
//
// Go automatically recognizes that a type implements an interface
// when it has all the required methods with matching signatures.

type PaymentGateway interface {
	Pay(request PaymentRequest) error
}

// ============================================
// 3. RAZORPAY
// ============================================

// Imagine this represents a Razorpay integration.
//
// We are simulating the provider response instead of making
// a real HTTP request to Razorpay.

type Razorpay struct {
	KeyID string
}

func (r Razorpay) Pay(request PaymentRequest) error {
	fmt.Println("Processing payment through Razorpay...")
	fmt.Println("Razorpay Key ID:", r.KeyID)
	fmt.Println("Order ID:", request.OrderID)
	fmt.Printf("Amount: %.2f %s\n", request.Amount, request.Currency)
	fmt.Println("Razorpay payment successful!")

	return nil
}

// ============================================
// 4. STRIPE
// ============================================

// Stripe has its own implementation of Pay().
//
// Notice that the method signature is the same as Razorpay's,
// but the code inside the method can be different.

type Stripe struct {
	SecretKey string
}

func (s Stripe) Pay(request PaymentRequest) error {
	fmt.Println("Processing payment through Stripe...")
	fmt.Println("Stripe configured:", s.SecretKey != "")
	fmt.Println("Order ID:", request.OrderID)
	fmt.Printf("Amount: %.2f %s\n", request.Amount, request.Currency)
	fmt.Println("Stripe payment successful!")

	return nil
}

// ============================================
// 5. PAYPAL
// ============================================

// PayPal also implements the same PaymentGateway interface.

type PayPal struct {
	MerchantID string
}

func (p PayPal) Pay(request PaymentRequest) error {
	fmt.Println("Processing payment through PayPal...")
	fmt.Println("Merchant ID:", p.MerchantID)
	fmt.Println("Order ID:", request.OrderID)
	fmt.Printf("Amount: %.2f %s\n", request.Amount, request.Currency)
	fmt.Println("PayPal payment successful!")

	return nil
}

// ============================================
// 6. OUR E-COMMERCE CHECKOUT SERVICE
// ============================================

// This service is the important part.
//
// It accepts the PaymentGateway interface, NOT Razorpay,
// Stripe, or PayPal specifically.
//
// Therefore, the checkout service does not need to know
// which payment provider is being used.

type CheckoutService struct{}

func (c CheckoutService) Checkout(
	gateway PaymentGateway,
	request PaymentRequest,
) error {
	fmt.Println("\n========== CHECKOUT ==========")

	err := gateway.Pay(request)

	if err != nil {
		fmt.Println("Payment failed:", err)
		return err
	}

	fmt.Println("Checkout completed!")
	return nil
}

// ============================================
// 7. MAIN FUNCTION
// ============================================

func main() {
	// Create our checkout service.
	checkout := CheckoutService{}

	// Example order.
	order := PaymentRequest{
		OrderID:  "ORDER-101",
		Amount:   1500,
		Currency: "INR",
	}

	// ----------------------------------------
	// OPTION 1: PAY USING RAZORPAY
	// ----------------------------------------

	razorpay := Razorpay{
		KeyID: "rzp_test_example",
	}

	checkout.Checkout(razorpay, order)

	// ----------------------------------------
	// OPTION 2: PAY USING STRIPE
	// ----------------------------------------

	stripe := Stripe{
		SecretKey: "sk_test_example",
	}

	checkout.Checkout(stripe, order)

	// ----------------------------------------
	// OPTION 3: PAY USING PAYPAL
	// ----------------------------------------

	paypal := PayPal{
		MerchantID: "merchant-101",
	}

	checkout.Checkout(paypal, order)
}
