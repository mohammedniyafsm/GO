package main

import "fmt"

// ============================================
// 1. WHY DO WE NEED STRUCTS?
// ============================================

// A struct groups related data into one type.
//
// Imagine we are building an e-commerce backend.
//
// A user has:
//   Name
//   Email
//   Age
//
// Without a struct, we might create separate variables
// for every user's name, email, and age.
//
// When the application grows, this becomes difficult
// to organize and maintain.
//
// A struct allows us to group related fields together.
//
// JavaScript comparison:
//
// const user = {
//     name: "Niyaf",
//     email: "niyaf@example.com",
//     age: 24
// }

// ============================================
// 2. DEFINING A STRUCT
// ============================================

// User is the name of our new struct type.
//
// Name is a string.
// Email is a string.
// Age is an integer.
//
// IMPORTANT:
// A struct describes the data that a value contains.

type User struct {
	Name  string
	Email string
	Age   int
}

// ============================================
// 3. CREATING STRUCT VALUES
// ============================================

// We can create multiple users using the same
// User struct definition.
//
// Each value has its own field data.

func createUsers() {

	user1 := User{
		Name:  "Niyaf",
		Email: "niyaf@example.com",
		Age:   24,
	}

	user2 := User{
		Name:  "Rahul",
		Email: "rahul@example.com",
		Age:   25,
	}

	fmt.Println(user1.Name)
	fmt.Println(user2.Name)

	// OUTPUT:
	// Niyaf
	// Rahul

	// Both values have the same struct type.
	// However, their field values are different.
}

// ============================================
// 4. WHAT IS AN INTERFACE?
// ============================================

// An interface defines a set of method requirements.
//
// A struct describes data.
//
// An interface describes behavior.
//
// Imagine that different animals can make sounds.
//
// A Dog makes a "Woof!" sound.
// A Cat makes a "Meow!" sound.
//
// We want one function to work with both animals.
//
// Instead of writing one function for Dog and another
// function for Cat, we can define an Animal interface.
//
// IMPORTANT:
// An interface does not require a struct to contain
// particular data fields.
//
// It requires matching methods.

// ============================================
// 5. DEFINING AN INTERFACE
// ============================================

// Animal requires a method named Sound.
//
// Sound() string means:
//   The method takes no arguments.
//   The method returns a string.
//
// Any type with this matching method satisfies Animal.

type Animal interface {
	Sound() string
}

// ============================================
// 6. STRUCTS THAT SATISFY AN INTERFACE
// ============================================

// Dog and Cat are concrete struct types.
//
// They do not need to declare:
//     implements Animal
//
// Go checks whether they have the required methods.

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// ============================================
// 7. ADDING METHODS TO STRUCTS
// ============================================

// Dog has a Sound method.
//
// The receiver (d Dog) associates the method
// with the Dog type.
//
// The method signature matches Animal.Sound().
//
// Therefore, Dog satisfies Animal.

func (d Dog) Sound() string {
	return "Woof!"
}

// Cat also provides the required Sound method.
//
// Therefore, Cat also satisfies Animal.

func (c Cat) Sound() string {
	return "Meow!"
}

// ============================================
// 8. USING AN INTERFACE AS A PARAMETER
// ============================================

// This function accepts Animal, not specifically Dog
// or Cat.
//
// It can accept any value whose type satisfies Animal.
//
// The function only needs to know that Sound()
// is available.
//
// It does not need to know how each animal
// implements its sound.

func makeSound(animal Animal) {
	fmt.Println(animal.Sound())
}

// ============================================
// 9. WHY IS THIS USEFUL?
// ============================================

// We can pass different concrete types to the same
// function because both satisfy Animal.
//
// This reduces the need to write separate functions
// for every animal type.

func demonstrateAnimals() {

	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	makeSound(dog)
	makeSound(cat)

	// OUTPUT:
	// Woof!
	// Meow!

	// The same makeSound function handles both values.
	//
	// Dog provides its own Sound implementation.
	// Cat provides its own Sound implementation.
	//
	// This is one of the main benefits of interfaces.
}

// ============================================
// 10. STRUCT VS INTERFACE
// ============================================

// STRUCT:
//
// type User struct {
//     Name string
//     Age  int
// }
//
// A struct represents data.
//
// We can create concrete values such as:
//
//     User{Name: "Niyaf", Age: 24}
//
// INTERFACE:
//
// type Animal interface {
//     Sound() string
// }
//
// An interface represents required behavior.
//
// We can use it as a parameter:
//
//     func makeSound(a Animal)
//
// IMPORTANT:
//
// A struct and an interface are not replacements
// for each other.
//
// A struct can provide data and methods.
//
// An interface defines which methods a value must
// have to be accepted in a particular context.
//
// A struct can satisfy many interfaces if it provides
// all the methods required by each interface.

// ============================================
// 11. REAL BACKEND CASE STUDY: PAYMENTS
// ============================================

// Imagine we are building an e-commerce backend.
//
// Our application supports:
//   Card payment
//   UPI payment
//   Cash payment
//
// Every payment method processes a payment, but
// each method may use a different implementation.
//
// We want one checkout function that works with
// every supported payment method.
//
// We can use an interface to describe the behavior.

// ============================================
// 12. PAYMENT INTERFACE
// ============================================

// PaymentProcessor requires a Pay method.
//
// Pay accepts an amount of type int.
// Pay returns a bool indicating success or failure.
//
// Any type with this exact method signature satisfies
// PaymentProcessor.

type PaymentProcessor interface {
	Pay(amount int) bool
}

// ============================================
// 13. PAYMENT STRUCTS
// ============================================

// These structs represent concrete payment methods.
//
// CardPayment stores card-related information.
// UPIPayment stores UPI-related information.
//
// The fields are simplified for this example.

type CardPayment struct {
	CardNumber string
}

type UPIPayment struct {
	UPIID string
}

// ============================================
// 14. IMPLEMENTING PAYMENT METHODS
// ============================================

// CardPayment implements Pay(int) bool.
//
// This example only checks whether the amount
// is positive.
//
// A real application would call a secure payment
// provider and handle its response.

func (c CardPayment) Pay(amount int) bool {
	fmt.Println("Processing card payment:", amount)

	return amount > 0
}

// UPIPayment implements the same required method.
//
// The implementation is different, but the method
// signature is identical.

func (u UPIPayment) Pay(amount int) bool {
	fmt.Println("Processing UPI payment:", amount)

	return amount > 0
}

// ============================================
// 15. CHECKOUT USING AN INTERFACE
// ============================================

// Notice the parameter type:
//
//     PaymentProcessor
//
// It is not CardPayment.
// It is not UPIPayment.
//
// The function accepts any compatible payment method.
//
// This means checkout does not need to know the
// concrete payment type.

func checkout(payment PaymentProcessor, amount int) {

	success := payment.Pay(amount)

	if success {
		fmt.Println("Payment successful")
	} else {
		fmt.Println("Payment failed")
	}
}

// ============================================
// 16. USING THE CHECKOUT FUNCTION
// ============================================

func demonstratePayments() {

	card := CardPayment{
		CardNumber: "1234",
	}

	upi := UPIPayment{
		UPIID: "niyaf@upi",
	}

	// Both concrete types can be passed to the
	// same checkout function.

	checkout(card, 500)

	fmt.Println("----------")

	checkout(upi, 1000)

	// OUTPUT:
	//
	// Processing card payment: 500
	// Payment successful
	// ----------
	// Processing UPI payment: 1000
	// Payment successful
}

// ============================================
// 17. WHAT IF WE ADD A NEW PAYMENT METHOD?
// ============================================

// Imagine we now support cash payments.
//
// We define another struct.

type CashPayment struct{}

// CashPayment implements the required Pay method.
//
// Because its signature matches the interface,
// it satisfies PaymentProcessor automatically.

func (c CashPayment) Pay(amount int) bool {
	fmt.Println("Processing cash payment:", amount)

	return amount > 0
}

// We can use the existing checkout function:
//
//     checkout(CashPayment{}, 200)
//
// We do not need to rewrite checkout.
//
// This makes the application easier to extend.

// ============================================
// 18. REAL BACKEND DESIGN: WHY AN INTERFACE
// ============================================

// Imagine a user registration API.
//
// The API needs to save users to a database.
//
// We could make the service depend directly on
// a specific PostgreSQL repository.
//
// However, this tightly couples the service to
// that particular implementation.
//
// Instead, we can define a UserRepository interface.
//
// The service depends on the required behavior,
// not on a specific database implementation.
//
// A PostgreSQL repository could implement this
// interface.
//
// A test repository could implement it too.

// ============================================
// 19. USER REPOSITORY INTERFACE
// ============================================

type UserRepository interface {
	Save(user User) error
}

// The interface requires one method:
// Save(user User) error
//
// A concrete type must implement the method with
// the exact matching signature to satisfy it.

// ============================================
// 20. IN-MEMORY REPOSITORY STRUCT
// ============================================

// MemoryRepository is a concrete implementation.
//
// It stores users in a slice instead of a database.
//
// This makes the example runnable without setting
// up PostgreSQL.

type MemoryRepository struct {
	Users []User
}

// Save implements UserRepository.Save.
//
// The receiver is a pointer because this method
// modifies the repository's Users slice.

func (r *MemoryRepository) Save(user User) error {
	r.Users = append(r.Users, user)

	fmt.Println("User saved:", user.Name)

	return nil
}

// ============================================
// 21. USER SERVICE STRUCT
// ============================================

// UserService contains business logic.
//
// Its repo field has the UserRepository interface type.
//
// IMPORTANT:
//
// UserService does not need to know whether its
// repository uses PostgreSQL, memory, or another
// storage implementation.
//
// It only needs the Save behavior.

type UserService struct {
	repo UserRepository
}

// Register creates a User and asks the repository
// to save it.

func (s UserService) Register(name, email string) error {

	if name == "" || email == "" {
		return fmt.Errorf("name and email are required")
	}

	user := User{
		Name:  name,
		Email: email,
		Age:   0,
	}

	return s.repo.Save(user)
}

// ============================================
// 22. CONNECTING THE STRUCTS AND INTERFACE
// ============================================

func demonstrateBackend() {

	// Create the concrete repository.

	repo := &MemoryRepository{
		Users: []User{},
	}

	// Pass the repository into the service.
	//
	// *MemoryRepository satisfies UserRepository
	// because it has the required Save method.

	service := UserService{
		repo: repo,
	}

	// Register a user.

	err := service.Register(
		"Niyaf",
		"niyaf@example.com",
	)

	if err != nil {
		fmt.Println("Error:", err)
	}

	// OUTPUT:
	// User saved: Niyaf

	// The service uses the interface.
	// The actual implementation is MemoryRepository.
}

// ============================================
// 23. WHY THIS DESIGN HELPS TESTING
// ============================================

// We can replace MemoryRepository with another
// type that implements UserRepository.
//
// For example, a test repository can store users
// in memory or deliberately return an error.
//
// That allows us to test the service without
// requiring a real database.
//
// The important point is:
//
// UserService depends on UserRepository behavior,
// not directly on MemoryRepository.
//
// This is useful when an application has multiple
// implementations or when we need isolated tests.

// ============================================
// 24. IMPORTANT: DO NOT CREATE INTERFACES
// JUST FOR THE SAKE OF CREATING THEM
// ============================================

// Interfaces are useful when they provide a real
// benefit, such as:
//
//   1. Supporting multiple implementations.
//   2. Separating business logic from infrastructure.
//   3. Making dependencies easier to replace in tests.
//   4. Allowing one function to accept different types.
//
// If a concrete struct is enough, use the struct.
//
// Not every struct needs its own interface.

// ============================================
// 25. MAIN FUNCTION
// ============================================

func main() {

	// Example 1: Structs group related data.

	fmt.Println("===== STRUCT EXAMPLE =====")
	createUsers()

	// OUTPUT:
	// ===== STRUCT EXAMPLE =====
	// Niyaf
	// Rahul

	fmt.Println()

	// Example 2: Interfaces allow one function
	// to work with different concrete types.

	fmt.Println("===== INTERFACE EXAMPLE =====")
	demonstrateAnimals()

	// OUTPUT:
	// ===== INTERFACE EXAMPLE =====
	// Woof!
	// Meow!

	fmt.Println()

	// Example 3: Payment implementations share
	// the same interface.

	fmt.Println("===== PAYMENT EXAMPLE =====")
	demonstratePayments()

	// OUTPUT:
	// ===== PAYMENT EXAMPLE =====
	// Processing card payment: 500
	// Payment successful
	// ----------
	// Processing UPI payment: 1000
	// Payment successful

	fmt.Println()

	// Example 4: A backend service depends on
	// a repository interface.

	fmt.Println("===== BACKEND EXAMPLE =====")
	demonstrateBackend()

	// OUTPUT:
	// ===== BACKEND EXAMPLE =====
	// User saved: Niyaf
}

// ============================================
// FINAL SUMMARY
// ============================================
//
// STRUCT:
//
// Represents data and can have methods.
//
// Examples:
//   User
//   Dog
//   CardPayment
//   MemoryRepository
//   UserService
//
// INTERFACE:
//
// Defines required behavior.
//
// Examples:
//   Animal
//   PaymentProcessor
//   UserRepository
//
// WHY USE BOTH?
//
// Structs represent the concrete data and implementations.
//
// Interfaces allow code to depend on behavior instead
// of a particular concrete type.
//
// REMEMBER:
//
// STRUCT = WHAT DATA DOES IT HOLD?
//
// INTERFACE = WHAT BEHAVIOR MUST IT PROVIDE?
//
// Start with concrete structs and introduce interfaces
// when they make your code more flexible or testable.
