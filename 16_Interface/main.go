package main

import "fmt"

// ============================================
// 1. WHAT IS AN INTERFACE?
// ============================================

// An interface defines a set of method requirements.
//
// A type satisfies an interface if it implements
// all the methods required by that interface.
//
// IMPORTANT:
// Go uses IMPLICIT INTERFACE IMPLEMENTATION.
//
// This means we do not need to write:
//
//     implements Animal
//
// Unlike Java, Go automatically determines whether
// a type satisfies an interface based on its methods.
//
// WHY DO WE NEED INTERFACES?
//
// Imagine we have:
//   Dog
//   Cat
//   Cow
//
// All three animals can make a sound.
//
// Instead of writing separate functions for every
// animal, we can define one interface requiring
// a Sound() method.
//
// Then one function can work with any type that
// satisfies that interface.
//
// This helps us write flexible and reusable code.
//
// JavaScript comparison:
//
// In JavaScript, we might pass different objects
// that have a sound() method.
//
// Go makes this relationship explicit through
// an interface type.

// ============================================
// 2. DEFINING AN INTERFACE
// ============================================

// The type keyword declares a new type.
//
// Animal is the interface name.
//
// interface defines a set of required methods.
//
// Sound() string means:
// A type must have a method named Sound,
// which takes no arguments and returns a string.
//
// IMPORTANT:
// An interface describes behavior, not the data
// fields a struct must contain.

type Animal interface {
	Sound() string
}

// ============================================
// 3. CREATING STRUCTS
// ============================================

// Dog and Cat are ordinary struct types.
//
// They do not explicitly declare that they implement
// the Animal interface.
//
// They can satisfy Animal by providing the required
// Sound() string method.

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// ============================================
// 4. IMPLEMENTING AN INTERFACE
// ============================================

// We define a Sound method for Dog.
//
// The receiver (d Dog) means this method belongs
// to the Dog type.
//
// The method signature is:
//
//     Sound() string
//
// This exactly matches the method required by Animal.
//
// Therefore, Dog satisfies Animal.

func (d Dog) Sound() string {
	return "Woof!"
}

// Cat also implements the required method.

func (c Cat) Sound() string {
	return "Meow!"
}

// ============================================
// 5. USING AN INTERFACE AS A FUNCTION PARAMETER
// ============================================

// The parameter a has type Animal.
//
// Any value whose type satisfies Animal can be
// passed to this function.
//
// The function does not need to know whether the
// actual value is a Dog or a Cat.
//
// It only needs to know that Sound() is available.

func makeSound(a Animal) {
	fmt.Println(a.Sound())
}

func main() {

	// ============================================
	// 6. PASSING DIFFERENT TYPES TO ONE FUNCTION
	// ============================================

	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	makeSound(dog)
	makeSound(cat)

	// Output:
	// Woof!
	// Meow!
	//
	// Explanation:
	//
	// First, dog is passed to makeSound.
	// Dog has the required Sound() string method.
	//
	// Next, cat is passed to makeSound.
	// Cat also has the required Sound() string method.
	//
	// The same function works with both types.
	//
	// This is one of the main benefits of interfaces.

	// ============================================
	// 7. INTERFACE VARIABLES
	// ============================================

	// An interface variable can hold a value of
	// any concrete type that satisfies that interface.

	var animal Animal

	animal = dog

	fmt.Println(animal.Sound())

	animal = cat

	fmt.Println(animal.Sound())

	// Output:
	// Woof!
	// Meow!
	//
	// Explanation:
	//
	// Initially, animal holds a Dog value.
	//
	// Later, animal holds a Cat value.
	//
	// The interface variable provides access to
	// methods required by Animal.
	//
	// The actual method executed depends on the
	// concrete value stored in the interface.

	// ============================================
	// 8. INTERFACE IMPLEMENTATION IS IMPLICIT
	// ============================================

	// Notice that neither Dog nor Cat contains
	// any special declaration such as:
	//
	//     implements Animal
	//
	// Go checks the methods automatically.
	//
	// Because both types have:
	//
	//     Sound() string
	//
	// Both satisfy Animal.
	//
	// If a type is missing the required method,
	// it cannot be passed where Animal is required.

	// ============================================
	// 9. AN INTERFACE CAN REQUIRE MULTIPLE METHODS
	// ============================================

	// A type must implement EVERY method in the interface.
	//
	// If an interface requires:
	//   Sound() string
	//   Move() string
	//
	// A type that only implements Sound() does not
	// satisfy the complete interface.
	//
	// We define the interface and its implementations
	// outside main below.

	robot := Robot{Name: "R2"}

	describeCreature(dog)
	describeCreature(robot)

	// Output:
	// Buddy says Woof! and can move: Running
	// R2 says Beep! and can move: Rolling

	// Explanation:
	// Both Dog and Robot implement Creature.
	//
	// Dog has Sound() and Move().
	// Robot has Sound() and Move().
	//
	// Therefore, both can be passed to describeCreature.

	// ============================================
	// 10. INTERFACE WITH DIFFERENT RETURN TYPES
	// ============================================

	// Interfaces can require methods with parameters
	// and return values of different types.
	//
	// Here, PaymentMethod requires Pay(int) bool.
	//
	// The method must:
	//   Accept an int.
	//   Return a bool.
	//
	// A method with a different signature does not
	// satisfy this interface.

	var payment PaymentMethod

	payment = CardPayment{CardNumber: "1234"}

	fmt.Println(payment.Pay(500))

	payment = CashPayment{}

	fmt.Println(payment.Pay(500))

	// Output:
	// true
	// true
	//
	// Explanation:
	// Both CardPayment and CashPayment implement
	// the exact method Pay(int) bool.
	//
	// The payment variable can hold either type.

	// ============================================
	// 11. THE EMPTY INTERFACE: interface{}
	// ============================================

	// An empty interface has no required methods.
	//
	// Because it requires no methods, every type
	// satisfies it.
	//
	// In modern Go, any is an alias for interface{}.
	//
	// We can therefore use any when a value of
	// an unknown or varying type must be accepted.

	var value any

	value = 100
	fmt.Println(value)

	value = "Hello Go"
	fmt.Println(value)

	value = true
	fmt.Println(value)

	// Output:
	// 100
	// Hello Go
	// true
	//
	// Explanation:
	// The same interface variable holds values
	// of different concrete types at different times.
	//
	// IMPORTANT:
	// any does not mean Go has become dynamically typed.
	// The concrete values still have their own types.
	//
	// The compiler cannot directly use arbitrary
	// concrete-type-specific operations on an any value.

	// ============================================
	// 12. TYPE ASSERTION
	// ============================================

	// A type assertion extracts a concrete value
	// from an interface value.
	//
	// Syntax:
	//
	//     result := value.(string)
	//
	// This asserts that the dynamic value inside
	// the interface is a string.
	//
	// If the assertion is wrong and we use the
	// one-result form, the program panics at runtime.

	var data any = "Go Programming"

	textValue := data.(string)

	fmt.Println(textValue)

	// Output:
	// Go Programming

	// We can use the comma-ok form to check safely.

	numberValue, ok := data.(int)

	fmt.Println(numberValue)
	fmt.Println(ok)

	// Output:
	// 0
	// false
	//
	// Explanation:
	// data contains a string, not an int.
	//
	// The assertion fails safely.
	// numberValue receives the int zero value.
	// ok becomes false.
	//
	// No panic occurs.

	// ============================================
	// 13. TYPE SWITCH
	// ============================================

	// A type switch checks the concrete type stored
	// inside an interface value.
	//
	// It is useful when different types need
	// different handling.

	printType(100)
	printType("Hello")
	printType(true)
	printType(dog)

	// Output:
	// Integer: 100
	// String: Hello
	// Boolean: true
	// Other type: {Buddy}

	// Explanation:
	// The switch checks the concrete type stored
	// inside each any value.
	//
	// case int matches an integer.
	// case string matches a string.
	// case bool matches a boolean.
	// default handles everything else.

	// ============================================
	// 14. POINTER RECEIVERS AND INTERFACES
	// ============================================

	// Whether a type satisfies an interface depends
	// on its method set.
	//
	// If a method has a pointer receiver, such as:
	//
	//     func (p *Printer) Print()
	//
	// then *Printer has that method.
	//
	// Printer itself does not have that pointer-
	// receiver method in its method set.
	//
	// Therefore, if an interface requires Print(),
	// *Printer satisfies it, but Printer does not.

	printer := Printer{Name: "Office Printer"}

	// Pass a pointer because *Printer implements
	// the Printable interface.
	printDocument(&printer)

	// Output:
	// Printing from Office Printer

	// The following would not compile:
	//
	// printDocument(printer)
	//
	// Because Printer has no Print() method in
	// its method set. The method has a pointer receiver.

	// ============================================
	// 15. INTERFACES AND STRUCT EMBEDDING
	// ============================================

	// If a struct embeds another type, methods
	// may be promoted from the embedded type.
	//
	// Those promoted methods can help the outer type
	// satisfy an interface.
	//
	// However, the exact method set still matters,
	// especially when embedding pointers.

	manager := Manager{Name: "Niyaf"}

	var greeter Greeter = manager

	fmt.Println(greeter.Greet())

	// Output:
	// Hello, Niyaf
	//
	// Explanation:
	// Manager embeds PersonWithGreeting.
	//
	// PersonWithGreeting has Greet().
	//
	// The method is promoted to Manager, so Manager
	// satisfies the Greeter interface.

	// ============================================
	// 16. RETURNING AN INTERFACE FROM A FUNCTION
	// ============================================

	// Functions can return interface values.
	//
	// This allows the caller to work with behavior
	// without needing to depend on one concrete type.
	//
	// createAnimal returns an Animal interface.
	//
	// The returned concrete value is a Dog.

	animal2 := createAnimal()

	fmt.Println(animal2.Sound())

	// Output:
	// Woof!

	// ============================================
	// 17. INTERFACE COMPOSITION
	// ============================================

	// One interface can embed another interface.
	//
	// This allows us to combine method requirements.
	//
	// ReadWriter combines Reader and Writer.
	//
	// A type must implement both Read() and Write()
	// to satisfy ReadWriter.

	file := File{Name: "notes.txt"}

	useReadWriter(file)

	// Output:
	// Reading notes.txt
	// Writing notes.txt

	// ============================================
	// 18. A COMMON INTERFACE GOTCHA: NIL
	// ============================================

	// An interface value can be thought of as holding:
	//
	//   1. A dynamic concrete type.
	//   2. A dynamic concrete value.
	//
	// A truly nil interface has neither a dynamic
	// type nor a dynamic value.
	//
	// An interface holding a typed nil pointer is
	// NOT itself nil.

	var dogPointer *Dog = nil

	var animal3 Animal = dogPointer

	fmt.Println(dogPointer == nil)
	fmt.Println(animal3 == nil)

	// Output:
	// true
	// false
	//
	// Explanation:
	//
	// dogPointer is a nil *Dog pointer.
	//
	// But assigning it to Animal gives the interface
	// a dynamic type of *Dog.
	//
	// Therefore, animal3 is not a nil interface.
	//
	// IMPORTANT:
	// Calling a method on this interface may panic
	// if that method dereferences the nil receiver.
	//
	// We do not call animal3.Sound() here because
	// the current Sound implementation uses a value
	// receiver and this typed nil cannot be invoked
	// through it without a runtime panic.

	// ============================================
	// 19. WHY INTERFACES ARE USEFUL IN BACKEND CODE
	// ============================================

	// Imagine your backend can send notifications
	// using email or SMS.
	//
	// Instead of tightly coupling business logic
	// to one provider, define an interface:
	//
	// type Notifier interface {
	//     Send(message string) error
	// }
	//
	// EmailNotifier and SMSNotifier can implement it.
	//
	// Business logic can depend on Notifier instead
	// of depending directly on a particular provider.
	//
	// This makes code easier to test, replace and extend.
}

// ============================================
// INTERFACE 1: CREATURE
// ============================================

// Creature requires two methods.
//
// A type must implement both methods with exactly
// matching signatures to satisfy Creature.

type Creature interface {
	Sound() string
	Move() string
}

// ============================================
// STRUCT: ROBOT
// ============================================

type Robot struct {
	Name string
}

// Robot implements Sound() string.

func (r Robot) Sound() string {
	return "Beep!"
}

// Robot implements Move() string.

func (r Robot) Move() string {
	return "Rolling"
}

// ============================================
// METHOD ON DOG: MOVE
// ============================================

// Dog already has Sound() string.
// We add Move() string so Dog also satisfies Creature.

func (d Dog) Move() string {
	return "Running"
}

// ============================================
// FUNCTION ACCEPTING CREATURE
// ============================================

func describeCreature(c Creature) {
	// We can call the methods required by Creature.
	// We do not need to know the concrete type.

	name := ""

	switch value := c.(type) {
	case Dog:
		name = value.Name
	case Robot:
		name = value.Name
	}

	fmt.Println(name, "says", c.Sound(), "and can move:", c.Move())
}

// ============================================
// INTERFACE 2: PAYMENT METHOD
// ============================================

type PaymentMethod interface {
	Pay(amount int) bool
}

// ============================================
// STRUCT: CARD PAYMENT
// ============================================

type CardPayment struct {
	CardNumber string
}

func (c CardPayment) Pay(amount int) bool {
	// This is a simplified example.
	// A real payment system would validate and process
	// a payment through a secure payment provider.

	return amount > 0
}

// ============================================
// STRUCT: CASH PAYMENT
// ============================================

type CashPayment struct{}

func (c CashPayment) Pay(amount int) bool {
	// This is also a simplified example.

	return amount > 0
}

// ============================================
// FUNCTION: TYPE SWITCH
// ============================================

func printType(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Other type:", v)
	}
}

// ============================================
// INTERFACE 3: PRINTABLE
// ============================================

type Printable interface {
	Print()
}

// ============================================
// STRUCT: PRINTER
// ============================================

type Printer struct {
	Name string
}

// The pointer receiver means *Printer implements
// Printable, not Printer itself.

func (p *Printer) Print() {
	fmt.Println("Printing from", p.Name)
}

func printDocument(p Printable) {
	p.Print()
}

// ============================================
// INTERFACE 4: GREETER
// ============================================

type Greeter interface {
	Greet() string
}

// ============================================
// STRUCTS: EMBEDDED GREETING
// ============================================

type PersonWithGreeting struct {
	Name string
}

func (p PersonWithGreeting) Greet() string {
	return "Hello, " + p.Name
}

type Manager struct {
	PersonWithGreeting
}

// Manager promotes Greet() from its embedded field.
// Therefore, Manager satisfies Greeter.

// ============================================
// FUNCTION RETURNING AN INTERFACE
// ============================================

func createAnimal() Animal {
	return Dog{Name: "Buddy"}
}

// ============================================
// INTERFACE COMPOSITION
// ============================================

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

// ReadWriter requires both Read() and Write().

type ReadWriter interface {
	Reader
	Writer
}

// ============================================
// STRUCT: FILE
// ============================================

type File struct {
	Name string
}

func (f File) Read() {
	fmt.Println("Reading", f.Name)
}

func (f File) Write() {
	fmt.Println("Writing", f.Name)
}

func useReadWriter(rw ReadWriter) {
	rw.Read()
	rw.Write()
}
