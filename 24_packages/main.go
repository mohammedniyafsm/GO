package main

import (
	"fmt"
	"strings"
)

// ============================================
// 1. WHAT IS A PACKAGE?
// ============================================
//
// A package is a way to organize related Go code.
//
// Examples:
//
// fmt     = Formatting and printing.
// strings = Working with strings.
// os      = Working with the operating system.
// sync    = Synchronization between goroutines.
//
// In Go, every .go file starts by declaring its package.
//
// Example:
//
// package main
//
// IMPORTANT:
//
// package main is special.
// It is used to build an executable program.
//
// A package with another name, such as "payment",
// can be imported and reused by other packages.
//
// ============================================
// 2. WHY DO WE NEED PACKAGES?
// ============================================
//
// Imagine building a backend application with:
//
// - User management
// - Authentication
// - Payments
// - Database operations
// - Email services
//
// Putting everything into one huge file becomes
// difficult to maintain.
//
// Packages help organize related functionality.
//
// Example project structure:
//
// myapp/
//     go.mod
//     main.go
//     user/
//         user.go
//     payment/
//         payment.go
//     database/
//         database.go
//
// Each directory can contain a separate package.
//
// Benefits:
//
// 1. Better code organization.
// 2. Code reuse.
// 3. Easier maintenance.
// 4. Clear boundaries between components.
//
// ============================================
// 3. USING A PACKAGE
// ============================================
//
// We use the import keyword to use another package.
//
// Example:
//
// import "fmt"
//
// fmt is the package name.
//
// fmt.Println("Hello")
//
// Println is a function provided by fmt.
//
// Another example:
//
// import "strings"
//
// strings.ToUpper("golang")
//
// ToUpper converts a string to uppercase.
//
// ============================================
// 4. USING MULTIPLE PACKAGES
// ============================================

func demonstrateStandardPackages() {
	fmt.Println("\n--- 4. Standard Packages ---")

	// fmt is used to print output.
	fmt.Println("Hello from Go!")

	// strings is used to work with strings.
	fmt.Println(strings.ToUpper("golang"))

	// Contains checks whether one string
	// exists inside another string.
	fmt.Println(strings.Contains("golang", "go"))

	// TrimSpace removes whitespace from
	// the beginning and end of a string.
	fmt.Println(strings.TrimSpace("   Hello Go   "))
}

// Expected output:
//
// --- 4. Standard Packages ---
// Hello from Go!
// GOLANG
// true
// Hello Go

// ============================================
// 5. WHAT IS AN IMPORT ALIAS?
// ============================================
//
// An import alias gives an imported package
// a different name within the current file.
//
// Example:
//
// import str "strings"
//
// Now we can use:
//
// str.ToUpper("golang")
//
// Instead of:
//
// strings.ToUpper("golang")
//
// Aliases are useful when package names conflict
// or when a package needs a clearer local name.
//
// Avoid unnecessary aliases because they can
// make code harder to read.

func demonstrateImportAlias() {
	fmt.Println("\n--- 5. Import Alias ---")

	// This example uses the normal strings import.
	// The alias syntax is explained above.
	fmt.Println(strings.ToLower("HELLO GO"))
}

// Expected output:
//
// --- 5. Import Alias ---
// hello go

// ============================================
// 6. EXPORTED AND UNEXPORTED FUNCTIONS
// ============================================
//
// Go uses capitalization to control visibility
// outside a package.
//
// EXPORTED:
//
// Starts with an uppercase letter.
//
// Example:
//
// CalculateTotal()
//
// Another package can access this function.
//
// UNEXPORTED:
//
// Starts with a lowercase letter.
//
// Example:
//
// calculateDiscount()
//
// It can only be accessed within its own package.
//
// This is similar to JavaScript exports in purpose,
// but the syntax and rules are different.
//
// JavaScript:
//
// export function calculateTotal() {}
//
// Go:
//
// func CalculateTotal() {}
//
// In Go, the uppercase first letter makes the
// function exported automatically.

// Exported function.
func CalculateTotal(price int, quantity int) int {
	return price * quantity
}

// Unexported function.
func calculateDiscount(price int) int {
	return price / 10
}

func demonstrateVisibility() {
	fmt.Println("\n--- 6. Exported and Unexported ---")

	fmt.Println("Total:", CalculateTotal(100, 3))
	fmt.Println("Discount:", calculateDiscount(100))
}

// Expected output:
//
// --- 6. Exported and Unexported ---
// Total: 300
// Discount: 10
//
// IMPORTANT:
//
// Both functions can be called here because they
// are declared in the same package: main.
//
// A different package could call CalculateTotal,
// but it could not call calculateDiscount.

// ============================================
// 7. PACKAGE-LEVEL VARIABLES
// ============================================
//
// Variables can be declared at package level.
//
// Package-level variables can be accessed by
// functions in the same package.
//
// Their visibility outside the package also depends
// on whether their names are exported.
//
// Exported variable:
//
// var AppName = "MyApp"
//
// Unexported variable:
//
// var appVersion = "1.0"
//
// In real projects, avoid exposing mutable variables
// unnecessarily. Prefer functions or types that
// control how data is accessed and modified.

var AppName = "MyGoApplication"
var appVersion = "1.0"

func demonstratePackageVariables() {
	fmt.Println("\n--- 7. Package-Level Variables ---")

	fmt.Println("App name:", AppName)
	fmt.Println("Version:", appVersion)
}

// Expected output:
//
// --- 7. Package-Level Variables ---
// App name: MyGoApplication
// Version: 1.0

// ============================================
// 8. WHAT IS A MODULE?
// ============================================
//
// A module is a collection of related Go packages.
//
// A Go module is defined by a go.mod file.
//
// Example go.mod:
//
// module example.com/myapp
//
// go 1.27.0
//
// The exact Go version depends on your installation.
//
// Create a module using:
//
// go mod init example.com/myapp
//
// The module path identifies the module and forms
// part of the import path for its packages.
//
// For example, a package inside:
//
// myapp/payment/
//
// can be imported as:
//
// import "example.com/myapp/payment"
//
// provided the module path is example.com/myapp.
//
// IMPORTANT:
//
// Package = Organizes Go source code.
// Module  = Groups packages and manages dependencies.
// Import  = Makes another package available to your code.

// ============================================
// 9. PACKAGE VS MODULE VS IMPORT
// ============================================
//
// PACKAGE:
//
// package main
//
// Declares the package for a Go source file.
//
// MODULE:
//
// Defined by go.mod.
//
// Groups related packages and manages dependencies.
//
// IMPORT:
//
// import "fmt"
//
// Makes another package available in the current file.
//
// Example:
//
// module example.com/myapp
//
// Directory:
// myapp/payment/
//
// Import path:
// example.com/myapp/payment
//
// Package declaration inside payment/payment.go:
//
// package payment

// ============================================
// 10. IMPORTANT RULES ABOUT PACKAGES
// ============================================
//
// RULE 1:
//
// Every Go source file declares its package.
//
// RULE 2:
//
// Ordinary Go files in the same directory must
// use the same package name.
//
// RULE 3:
//
// package main is used to build an executable.
//
// RULE 4:
//
// An executable main package needs a main()
// function as its entry point.
//
// RULE 5:
//
// Uppercase names are exported.
//
// Lowercase names are unexported.
//
// RULE 6:
//
// Imports use package import paths, but code
// normally uses the package's declared name.
//
// RULE 7:
//
// Packages cannot import each other in a cycle.
//
// For example:
//
// package A imports package B.
// package B imports package A.
//
// This circular dependency is not allowed.
//
// RULE 8:
//
// Imported packages that are not used usually
// cause a compilation error.
//
// ============================================
// 11. REAL-WORLD BACKEND EXAMPLE
// ============================================
//
// Imagine a payment backend.
//
// Possible project structure:
//
// payment-app/
//     go.mod
//     main.go
//     payment/
//         payment.go
//     user/
//         user.go
//
// main.go:
//
// Starts the application.
//
// payment/payment.go:
//
// Contains payment-related logic.
//
// user/user.go:
//
// Contains user-related logic.
//
// In payment/payment.go:
//
// package payment
//
// func ProcessPayment(amount int) string {
//     return "Payment processed"
// }
//
// In main.go, with module path
// example.com/payment-app:
//
// import "example.com/payment-app/payment"
//
// func main() {
//     result := payment.ProcessPayment(500)
//     fmt.Println(result)
// }
//
// Expected output:
//
// Payment processed
//
// IMPORTANT:
//
// This is an illustration of how separate packages
// work together. You cannot declare both package main
// and package payment in this single source file.
//
// ============================================
// 12. WHEN SHOULD YOU USE PACKAGES?
// ============================================
//
// Use separate packages when related functionality
// needs a clear boundary.
//
// Examples:
//
// payment package:
// Payment processing logic.
//
// user package:
// User-related operations.
//
// database package:
// Database connection and queries.
//
// middleware package:
// HTTP middleware.
//
// In small learning programs, one package is fine.
//
// In larger backend applications, organize code
// based on responsibilities and dependencies.
//
// Do not create a new package for every tiny function.
// A package should represent a useful, cohesive unit.

// ============================================
// 13. MAIN FUNCTION
// ============================================

func main() {
	fmt.Println("Go Packages Learning")

	demonstrateStandardPackages()
	demonstrateImportAlias()
	demonstrateVisibility()
	demonstratePackageVariables()
}

// ============================================
// EXPECTED OUTPUT
// ============================================
//
// Go Packages Learning
//
// --- 4. Standard Packages ---
// Hello from Go!
// GOLANG
// true
// Hello Go
//
// --- 5. Import Alias ---
// hello go
//
// --- 6. Exported and Unexported ---
// Total: 300
// Discount: 10
//
// --- 7. Package-Level Variables ---
// App name: MyGoApplication
// Version: 1.0
