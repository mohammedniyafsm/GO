package main

import "fmt"

// ============================================
// 1. WHAT IS A FUNCTION?
// ============================================

// A function is a reusable block of code that performs
// a particular task.
//
// Why use functions?
// 1. Avoid repeating the same code.
// 2. Make code easier to understand.
// 3. Divide a large program into smaller parts.
//
// Syntax:
//
// func functionName() {
//     // code
// }
//
// func is the keyword used to declare a function.

func sayHello() {
	fmt.Println("Hello, Go!")
}

func main() {

	// ============================================
	// 2. CALL A FUNCTION
	// ============================================

	// Defining a function does not execute it.
	// We must call it using its name and parentheses.

	sayHello()

	// Output:
	// Hello, Go!

	// ============================================
	// 3. FUNCTION WITH PARAMETERS
	// ============================================

	// Parameters allow us to pass data into a function.
	//
	// Here:
	// name is a parameter of type string.

	greet("Niyaf")

	// Output:
	// Hello, Niyaf

	// ============================================
	// 4. FUNCTION WITH MULTIPLE PARAMETERS
	// ============================================

	// A function can accept multiple parameters.
	// Each parameter has a type.

	addAndPrint(10, 20)

	// Output:
	// 30

	// ============================================
	// 5. FUNCTION WITH A RETURN VALUE
	// ============================================

	// A function can return a result to the caller.
	// The return type is written after the parentheses.

	result := add(10, 20)

	fmt.Println(result)

	// Output:
	// 30
	//
	// Explanation:
	// add(10, 20) calculates 10 + 20.
	// The returned value 30 is stored in result.

	// ============================================
	// 6. FUNCTION WITH MULTIPLE RETURN VALUES
	// ============================================

	// Go allows a function to return multiple values.
	// Here, divide returns both the quotient and remainder.

	quotient, remainder := divide(10, 3)

	fmt.Println(quotient)
	fmt.Println(remainder)

	// Output:
	// 3
	// 1
	//
	// Explanation:
	// 10 / 3 gives quotient 3 and remainder 1.

	// ============================================
	// 7. IGNORE A RETURN VALUE USING _
	// ============================================

	// If we do not need a returned value, use _.
	// This is called the blank identifier.

	quotientOnly, _ := divide(10, 3)

	fmt.Println(quotientOnly)

	// Output:
	// 3

	// ============================================
	// 8. NAMED RETURN VALUES
	// ============================================

	// Go lets us name the return values in the signature.
	// These names act like local variables in the function.

	sum, product := calculate(4, 5)

	fmt.Println(sum)
	fmt.Println(product)

	// Output:
	// 9
	// 20

	// ============================================
	// 9. FUNCTION WITH NO PARAMETERS AND NO RETURN
	// ============================================

	// Some functions simply perform an action.
	// They do not accept data or return a value.

	showMessage()

	// Output:
	// Learning Go functions

	// ============================================
	// 10. VARIADIC FUNCTIONS
	// ============================================

	// A variadic function accepts zero or more arguments
	// of the same type.
	//
	// The ... syntax is used before the parameter type.
	// Inside the function, numbers behaves like a slice.

	fmt.Println(sumAll(10, 20, 30))
	fmt.Println(sumAll(1, 2, 3, 4, 5))

	// Output:
	// 60
	// 15

	// ============================================
	// 11. PASS A SLICE TO A VARIADIC FUNCTION
	// ============================================

	// If you already have a slice, use ... to pass
	// its elements as variadic arguments.

	nums := []int{10, 20, 30}

	fmt.Println(sumAll(nums...))

	// Output:
	// 60

	// ============================================
	// 12. ANONYMOUS FUNCTIONS
	// ============================================

	// An anonymous function is a function without a name.
	// We can assign it to a variable and call it.

	multiply := func(a int, b int) int {
		return a * b
	}

	fmt.Println(multiply(4, 5))

	// Output:
	// 20

	// ============================================
	// 13. IMMEDIATELY INVOKED FUNCTION
	// ============================================

	// We can define and execute a function immediately.
	// The final () calls the function.

	func() {
		fmt.Println("Function executed immediately")
	}()

	// Output:
	// Function executed immediately

	// ============================================
	// 14. FUNCTION AS A VALUE
	// ============================================

	// In Go, functions can be stored in variables
	// and passed around like other values.

	operation := add

	fmt.Println(operation(5, 7))

	// Output:
	// 12
	//
	// operation refers to the add function.
	// operation(5, 7) calls that function.

	// ============================================
	// 15. PASS A FUNCTION TO ANOTHER FUNCTION
	// ============================================

	// A function can accept another function as a parameter.
	// This is useful for reusable operations and callbacks.

	answer := applyOperation(10, 5, subtract)

	fmt.Println(answer)

	// Output:
	// 5

	// ============================================
	// 16. RECURSION
	// ============================================

	// Recursion means a function calls itself.
	// Every recursive function needs a stopping condition.
	//
	// Here, factorial(5) = 5 * 4 * 3 * 2 * 1.

	fmt.Println(factorial(5))

	// Output:
	// 120

	// ============================================
	// 17. DEFER
	// ============================================

	// defer postpones a function call until the
	// surrounding function is about to return.
	//
	// Deferred calls execute in last-in, first-out order.

	defer fmt.Println("Third")
	defer fmt.Println("Second")

	fmt.Println("First")

	// Output at the end of main:
	// First
	// Second
	// Third
	//
	// Explanation:
	// The deferred calls execute in reverse order.

	// ============================================
	// 18. MULTIPLE DEFER CALLS
	// ============================================

	// Deferred calls are often used for cleanup,
	// such as closing files or releasing resources.
	//
	// This example demonstrates execution order.

	defer fmt.Println("defer A")
	defer fmt.Println("defer B")
	defer fmt.Println("defer C")

	fmt.Println("Normal statement")

	// Output at the end of main:
	// Normal statement
	// defer C
	// defer B
	// defer A
	//
	// Note:
	// These deferred calls run when main returns.
	// Because they were declared later in this example,
	// they execute before the earlier defers above.

	// ============================================
	// 19. FUNCTIONS AND ERROR RETURNS
	// ============================================

	// Go commonly returns an error as an additional value.
	// Callers should check the error before using a result.

	value, err := safeDivide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(value)
	}

	// Output:
	// 5

	// Example with an invalid operation:

	_, err = safeDivide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	}

	// Output:
	// Error: cannot divide by zero

}

// ============================================
// FUNCTION DEFINITIONS
// ============================================

// Function with no parameters and no return value.

func greet(name string) {
	fmt.Println("Hello,", name)
}

// When consecutive parameters have the same type,
// Go lets us write the type only once.

func addAndPrint(a, b int) {
	fmt.Println(a + b)
}

// Function that returns one integer.

func add(a, b int) int {
	return a + b
}

// Function returning two values.

func divide(a, b int) (int, int) {
	return a / b, a % b
}

// Function with named return values.

func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b

	return
	// This is called a naked return.
	// The named variables are returned automatically.
	// In larger functions, explicit return values may
	// be easier to read.
}

// Function with no parameters and no return value.

func showMessage() {
	fmt.Println("Learning Go functions")
}

// Variadic function.

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// Function used as an operation.

func subtract(a, b int) int {
	return a - b
}

// A function that accepts another function.

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Recursive function.

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

// Function returning a result and an error.

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}
