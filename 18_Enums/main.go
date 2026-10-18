package main

import "fmt"

// ============================================
// 1. WHAT IS AN ENUM?
// ============================================

// An enum represents a fixed set of possible values.
//
// Example: A payment can have these statuses:
// Pending, Success, Failed, Refunded.
//
// Go does not have a built-in enum keyword.
// We create enum-like values using const + iota.
//
// Similar idea in TypeScript:
// enum PaymentStatus {
//     Pending,
//     Success,
//     Failed
// }

// ============================================
// 2. CREATE ENUM USING const + iota
// ============================================

// PaymentStatus is a custom type based on int.
//
// We use a custom type so that payment statuses
// have a meaningful type instead of being plain integers.

type PaymentStatus int

const (
	Pending  PaymentStatus = iota // 0
	Success                       // 1
	Failed                        // 2
	Refunded                      // 3
)

// iota is an automatic counter inside a const block.
//
// Pending = 0
// Success = 1
// Failed  = 2
// Refunded = 3
//
// Each new line increments iota automatically.

// ============================================
// 3. START ENUM VALUES FROM 1
// ============================================

// Sometimes we want the first value to be 1 instead of 0.
//
// Example:
//
// const (
//     Pending PaymentStatus = iota + 1 // 1
//     Success                          // 2
//     Failed                           // 3
// )
//
// iota + 1 makes the first value 1.
//
// We are not using this second declaration here because
// PaymentStatus is already declared above.

// ============================================
// 4. CONVERT ENUM VALUES INTO READABLE STRINGS
// ============================================

// By default, printing Success prints 1.
//
// We can create a String() method to display a readable name.
//
// fmt.Println(Success) will print "Success"
// because fmt uses this method for the value.

func (status PaymentStatus) String() string {
	switch status {
	case Pending:
		return "Pending"
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	case Refunded:
		return "Refunded"
	default:
		return "Unknown"
	}
}

// ============================================
// 5. USE ENUM WITH A FUNCTION
// ============================================

// A function can accept PaymentStatus as its parameter.
//
// This makes it clear that the function expects a payment
// status, not just any integer.

func printPaymentStatus(status PaymentStatus) {
	switch status {
	case Pending:
		fmt.Println("Payment is pending")
	case Success:
		fmt.Println("Payment successful")
	case Failed:
		fmt.Println("Payment failed")
	case Refunded:
		fmt.Println("Payment refunded")
	default:
		fmt.Println("Unknown payment status")
	}
}

// ============================================
// 6. MAIN FUNCTION
// ============================================

func main() {

	// ----------------------------------------
	// A. PRINT ENUM VALUES
	// ----------------------------------------

	fmt.Println("===== Payment Status Names =====")

	fmt.Println(Pending)
	fmt.Println(Success)
	fmt.Println(Failed)
	fmt.Println(Refunded)

	// Output:
	// Pending
	// Success
	// Failed
	// Refunded

	// ----------------------------------------
	// B. PRINT ENUM NUMERIC VALUES
	// ----------------------------------------

	// Convert PaymentStatus into int to see its number.

	fmt.Println("\n===== Numeric Values =====")

	fmt.Println(int(Pending))
	fmt.Println(int(Success))
	fmt.Println(int(Failed))
	fmt.Println(int(Refunded))

	// Output:
	// 0
	// 1
	// 2
	// 3

	// ----------------------------------------
	// C. USE ENUM WITH A FUNCTION
	// ----------------------------------------

	fmt.Println("\n===== Payment Processing =====")

	printPaymentStatus(Pending)
	printPaymentStatus(Success)
	printPaymentStatus(Failed)
	printPaymentStatus(Refunded)

	// Output:
	// Payment is pending
	// Payment successful
	// Payment failed
	// Payment refunded

	// ----------------------------------------
	// D. USE ENUM IN A PAYMENT SYSTEM
	// ----------------------------------------

	fmt.Println("\n===== Payment Provider Example =====")

	razorpayStatus := Success
	stripeStatus := Pending
	paypalStatus := Failed

	fmt.Println("Razorpay:", razorpayStatus)
	fmt.Println("Stripe:", stripeStatus)
	fmt.Println("PayPal:", paypalStatus)

	// Output:
	// Razorpay: Success
	// Stripe: Pending
	// PayPal: Failed

	// ----------------------------------------
	// E. UNKNOWN / ZERO VALUE
	// ----------------------------------------

	// A variable of type PaymentStatus defaults to 0.
	//
	// Here, 0 represents Pending because Pending = 0.
	//
	// If we had started our constants from 1,
	// the zero value would not match any named status.

	var status PaymentStatus

	fmt.Println("\n===== Zero Value =====")
	fmt.Println(status)

	// Output:
	// Pending
}
