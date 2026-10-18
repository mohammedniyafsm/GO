package main

import "fmt"

// ============================================
// 1. WHAT ARE GENERICS?
// ============================================

// Generics allow us to write functions and data structures
// that work with different types without repeating code.
//
// Example without generics:
//
// func PrintInt(value int) {
//     fmt.Println(value)
// }
//
// func PrintString(value string) {
//     fmt.Println(value)
// }
//
// We need separate functions for different types.
//
// With generics, we can write ONE function that works
// with int, string, and other allowed types.
//
// Similar idea in TypeScript:
// function printValue<T>(value: T): void {
//     console.log(value)
// }

// ============================================
// 2. GENERIC FUNCTION
// ============================================

// [T any] means:
// T = a type parameter; it represents a type.
// any = T can be any type.
//
// The actual type is determined when we call the function.
//
// printValue(100)    -> T is int
// printValue("Go")   -> T is string
// printValue(true)   -> T is bool

func printValue[T any](value T) {
	fmt.Println(value)
}

// ============================================
// 3. GENERIC FUNCTION WITH DIFFERENT TYPES
// ============================================

func display[T any](value T) {
	fmt.Printf("Value: %v | Type: %T\n", value, value)
}

// %v prints the value.
// %T prints the actual type.
//
// Example:
// Value: 100 | Type: int
// Value: Go | Type: string

// ============================================
// 4. GENERIC FUNCTION WITH TWO PARAMETERS
// ============================================

// Both parameters use the same type T.
//
// Therefore, both arguments must have the same type
// when T is inferred.
//
// add(10, 20) works.
// add("Hello ", "Go") works.
//
// add(10, "Go") does NOT work because the types differ.

func add[T any](a T, b T) {
	fmt.Println(a, b)
}

// ============================================
// 5. GENERICS WITH A TYPE CONSTRAINT
// ============================================

// Sometimes we want to allow only specific types.
//
// This constraint allows int, int64, float32, and float64.
//
// The | symbol means OR between allowed types.
//
// ~ means the type and any custom type whose underlying
// type is that type.
//
// Example:
// type Age int
//
// Age can satisfy ~int because its underlying type is int.

type Number interface {
	~int | ~int64 | ~float32 | ~float64
}

// Now T must be one of the types allowed by Number.
//
// We can use + because all types in Number support addition.

func sum[T Number](a T, b T) T {
	return a + b
}

// ============================================
// 6. GENERIC FUNCTION RETURNING A VALUE
// ============================================

// Generics can return values too.
//
// Here, the function returns whichever value is greater.
// This constraint allows only numeric types.

func maximum[T Number](a T, b T) T {
	if a > b {
		return a
	}

	return b
}

// ============================================
// 7. GENERICS WITH SLICES
// ============================================

// We can write one function that works with slices
// containing different element types.
//
// []T means a slice containing values of type T.
//
// This function returns the first element.
//
// IMPORTANT:
// We must handle an empty slice because it has no first element.

func first[T any](items []T) (T, bool) {
	if len(items) == 0 {
		var zero T
		return zero, false
	}

	return items[0], true
}

// var zero T creates the zero value of T.
// For int, the zero value is 0.
// For string, the zero value is "".
// For bool, the zero value is false.
//
// The bool result tells us whether an element was found.

// ============================================
// 8. GENERIC STRUCT
// ============================================

// A generic struct can store a value of different types.
//
// Box[int] stores an int.
// Box[string] stores a string.
// Box[bool] stores a bool.

type Box[T any] struct {
	Value T
}

// ============================================
// 9. GENERICS WITH CUSTOM TYPES
// ============================================

// Generics also work with custom types when the type
// satisfies the constraint.
//
// Age is a custom type whose underlying type is int.

type Age int

func printNumber[T Number](value T) {
	fmt.Println("Number:", value)
}

// ============================================
// 10. MAIN FUNCTION
// ============================================

func main() {

	// ----------------------------------------
	// A. BASIC GENERIC FUNCTION
	// ----------------------------------------

	fmt.Println("===== Basic Generics =====")

	printValue(100)
	printValue("Hello Go")
	printValue(true)

	// Output:
	// 100
	// Hello Go
	// true

	// ----------------------------------------
	// B. DISPLAY VALUE AND TYPE
	// ----------------------------------------

	fmt.Println("\n===== Value and Type =====")

	display(25)
	display("Golang")
	display(3.14)

	// Output:
	// Value: 25 | Type: int
	// Value: Golang | Type: string
	// Value: 3.14 | Type: float64

	// ----------------------------------------
	// C. GENERIC FUNCTION WITH TWO PARAMETERS
	// ----------------------------------------

	fmt.Println("\n===== Two Parameters =====")

	add(10, 20)
	add("Hello ", "Go")

	// Output:
	// 10 20
	// Hello  Go

	// ----------------------------------------
	// D. GENERIC SUM
	// ----------------------------------------

	fmt.Println("\n===== Sum =====")

	fmt.Println(sum(10, 20))
	fmt.Println(sum(2.5, 3.5))
	fmt.Println(sum(int64(100), int64(200)))

	// Output:
	// 30
	// 6
	// 300

	// ----------------------------------------
	// E. FIND MAXIMUM
	// ----------------------------------------

	fmt.Println("\n===== Maximum =====")

	fmt.Println(maximum(10, 50))
	fmt.Println(maximum(4.5, 2.5))

	// Output:
	// 50
	// 4.5

	// ----------------------------------------
	// F. GENERICS WITH SLICES
	// ----------------------------------------

	fmt.Println("\n===== First Element =====")

	numbers := []int{10, 20, 30}
	firstNumber, found := first(numbers)

	fmt.Println(firstNumber, found)

	names := []string{"Niyaf", "Rahul", "Aman"}
	firstName, nameFound := first(names)

	fmt.Println(firstName, nameFound)

	empty := []int{}
	emptyNumber, emptyFound := first(empty)

	fmt.Println(emptyNumber, emptyFound)

	// Output:
	// 10 true
	// Niyaf true
	// 0 false

	// ----------------------------------------
	// G. GENERIC STRUCT
	// ----------------------------------------

	fmt.Println("\n===== Generic Struct =====")

	intBox := Box[int]{Value: 100}
	stringBox := Box[string]{Value: "Hello Go"}
	boolBox := Box[bool]{Value: true}

	fmt.Println(intBox.Value)
	fmt.Println(stringBox.Value)
	fmt.Println(boolBox.Value)

	// Output:
	// 100
	// Hello Go
	// true

	// ----------------------------------------
	// H. CUSTOM TYPE WITH CONSTRAINT
	// ----------------------------------------

	fmt.Println("\n===== Custom Type =====")

	var age Age = 25

	printNumber(age)

	// Output:
	// Number: 25
}
