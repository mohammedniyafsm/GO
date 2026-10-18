package main

import "fmt"

// ============================================
// 1. WHAT IS A VARIADIC FUNCTION?
// ============================================

// A variadic function is a function that can accept
// zero, one, two, or many arguments.
//
// Normally, a function expects a specific number
// of arguments.
//
// Example:
//
// func add(a, b int) int {
//     return a + b
// }
//
// This function accepts exactly two integers.
//
// But what if we want to pass 2, 3, 4, or 10 numbers?
// Creating separate functions for each case is unnecessary.
//
// Go provides variadic functions for this purpose.
//
// We use ... before the parameter type.
//
// Syntax:
//
// func functionName(parameterName ...type) {
//     // code
// }

func main() {

	// ============================================
	// 2. BASIC VARIADIC FUNCTION
	// ============================================

	// Here, printNumbers can accept any number of
	// integer arguments.
	//
	// We can pass one number, multiple numbers,
	// or even zero numbers.

	printNumbers(10)
	printNumbers(10, 20)
	printNumbers(10, 20, 30, 40)
	printNumbers()

	// Output:
	// [10]
	// [10 20]
	// [10 20 30 40]
	// []
	//
	// Explanation:
	// The ...int syntax allows a variable number
	// of integer arguments.
	//
	// Inside the function, numbers behaves like
	// a slice of integers ([]int).
	//
	// When we call printNumbers(10, 20, 30),
	// Go collects those arguments into a slice:
	//
	// numbers = []int{10, 20, 30}

	// ============================================
	// 3. VARIADIC FUNCTION TO CALCULATE A SUM
	// ============================================

	// We can use a variadic function to perform
	// an operation on any number of arguments.
	//
	// Here, sumAll adds all the numbers and
	// returns the final sum.

	fmt.Println(sumAll(10, 20))
	fmt.Println(sumAll(10, 20, 30))
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	fmt.Println(sumAll())

	// Output:
	// 30
	// 60
	// 15
	// 0
	//
	// Explanation:
	// sumAll(10, 20, 30) receives three numbers.
	//
	// The function starts with total = 0.
	//
	// First iteration:
	// total = 0 + 10 = 10
	//
	// Second iteration:
	// total = 10 + 20 = 30
	//
	// Third iteration:
	// total = 30 + 30 = 60
	//
	// Finally, return total returns 60.
	//
	// If no arguments are passed, the loop does
	// not execute and the function returns 0.

	// ============================================
	// 4. HOW ...int WORKS INSIDE A FUNCTION
	// ============================================

	// A variadic parameter behaves like a slice
	// inside the function.
	//
	// That means we can use len, range, and indexing.

	showNumbers(10, 20, 30)

	// Output:
	// Numbers: [10 20 30]
	// Length: 3
	// First number: 10
	//
	// Explanation:
	// len(numbers) gives the number of arguments.
	//
	// numbers[0] accesses the first element.
	// Slice indexes start at 0.
	//
	// Since numbers contains three elements,
	// its length is 3.

	// ============================================
	// 5. PASSING A SLICE TO A VARIADIC FUNCTION
	// ============================================

	// Sometimes, we already have a slice containing
	// the numbers we want to pass.
	//
	// Example:
	//
	// nums := []int{10, 20, 30}
	//
	// We cannot pass the slice directly like this:
	//
	// sumAll(nums)
	//
	// This produces a compile-time error because
	// sumAll expects individual int arguments,
	// but nums is a []int slice.
	//
	// To pass a slice's elements as variadic arguments,
	// we write ... after the slice name.

	nums := []int{10, 20, 30}

	fmt.Println(sumAll(nums...))

	// Output:
	// 60
	//
	// Explanation:
	// nums... passes the slice elements as arguments.
	//
	// It is equivalent to:
	//
	// sumAll(10, 20, 30)
	//
	// Remember:
	// numbers ...int  -> collects arguments into a slice.
	// nums...         -> passes slice elements as arguments.

	// ============================================
	// 6. VARIADIC FUNCTION WITH REGULAR PARAMETERS
	// ============================================

	// A function can have regular parameters
	// together with a variadic parameter.
	//
	// Here:
	// name is a regular string parameter.
	// scores ...int accepts multiple integer arguments.
	//
	// Important:
	// The variadic parameter must always be
	// the LAST parameter in the function definition.

	printStudentScores("Niyaf", 80, 90, 95)
	printStudentScores("Rahul", 75)

	// Output:
	// Student: Niyaf
	// Scores: [80 90 95]
	// Student: Rahul
	// Scores: [75]
	//
	// Explanation:
	// "Niyaf" is passed to the regular parameter name.
	//
	// 80, 90, and 95 are collected into scores.
	//
	// In the second call, only one score is passed.
	// That is valid because variadic parameters can
	// accept any number of arguments, including one.

	// ============================================
	// 7. VARIADIC FUNCTION WITH STRINGS
	// ============================================

	// Variadic functions are not limited to integers.
	// We can use any parameter type.
	//
	// Here, ...string accepts multiple strings.

	printNames("Niyaf", "Rahul", "Aisha")

	// Output:
	// Names: [Niyaf Rahul Aisha]

	// ============================================
	// 8. AN EMPTY VARIADIC ARGUMENT
	// ============================================

	// A variadic function can be called without
	// passing any arguments for its variadic parameter.
	//
	// Inside the function, the variadic slice will be nil
	// when no arguments are supplied directly.
	//
	// We can check this using len and comparing with nil.

	checkNumbers()
	checkNumbers(10, 20)

	// Output:
	// Numbers: []
	// Length: 0
	// Is nil: true
	//
	// Numbers: [10 20]
	// Length: 2
	// Is nil: false
	//
	// Explanation:
	// With no variadic arguments, the parameter is nil.
	//
	// With one or more arguments, Go creates a slice
	// containing those values.
	//
	// Note:
	// An empty but non-nil slice is also possible
	// when passing an existing slice using ... .

	// ============================================
	// 9. PASSING AN EMPTY SLICE
	// ============================================

	// An empty slice and a nil slice both have length 0.
	// However, they are not necessarily the same.
	//
	// Here, we create an empty, non-nil slice.

	emptyNums := []int{}

	fmt.Println(sumAll(emptyNums...))
	fmt.Println(len(emptyNums))
	fmt.Println(emptyNums == nil)

	// Output:
	// 0
	// 0
	// false
	//
	// Explanation:
	// emptyNums contains no elements, so the sum is 0.
	//
	// len(emptyNums) is 0 because there are no elements.
	//
	// emptyNums == nil is false because the slice
	// was created as an empty slice, not a nil slice.

	// ============================================
	// 10. A VARIADIC PARAMETER MUST BE LAST
	// ============================================

	// This is VALID:
	//
	// func example(name string, numbers ...int)
	//
	// This is INVALID:
	//
	// func example(numbers ...int, name string)
	//
	// Explanation:
	// Go needs to know which arguments belong to
	// the regular parameters and which belong to
	// the variadic parameter.
	//
	// Therefore, the variadic parameter must come last.

	// ============================================
	// 11. VARIADIC FUNCTIONS VS NORMAL FUNCTIONS
	// ============================================

	// Normal function:
	//
	// func add(a, b int) int
	//
	// It accepts exactly two integers.
	//
	// Variadic function:
	//
	// func sumAll(numbers ...int) int
	//
	// It accepts zero or more integers.

	fmt.Println(addTwo(10, 20))
	fmt.Println(sumAll(10, 20))
	fmt.Println(sumAll(10, 20, 30))

	// Output:
	// 30
	// 30
	// 60
	//
	// Explanation:
	// addTwo accepts exactly two arguments.
	//
	// sumAll can accept two, three, or any number
	// of integer arguments.
	//
	// Use a normal function when the number of
	// arguments is fixed.
	//
	// Use a variadic function when the number
	// of arguments can change.
}

// ============================================
// FUNCTION DEFINITIONS
// ============================================

// ============================================
// 1. BASIC VARIADIC FUNCTION
// ============================================

// numbers ...int means this function accepts
// zero or more integer arguments.
//
// Inside the function, numbers behaves like []int.

func printNumbers(numbers ...int) {
	fmt.Println(numbers)
}

// ============================================
// 2. VARIADIC FUNCTION THAT RETURNS A SUM
// ============================================

// This function accepts multiple integers
// and returns their sum.

func sumAll(numbers ...int) int {
	total := 0

	// range loops through every number in the slice.
	//
	// The underscore _ ignores the index because
	// we only need the number itself.

	for _, number := range numbers {
		total += number
	}

	// Return the final calculated total.
	return total
}

// ============================================
// 3. ACCESSING VARIADIC ARGUMENTS
// ============================================

func showNumbers(numbers ...int) {
	fmt.Println("Numbers:", numbers)

	fmt.Println("Length:", len(numbers))

	// Check the length before accessing numbers[0].
	// Otherwise, an empty argument list would cause
	// a runtime panic.

	if len(numbers) > 0 {
		fmt.Println("First number:", numbers[0])
	}
}

// ============================================
// 4. REGULAR PARAMETER + VARIADIC PARAMETER
// ============================================

func printStudentScores(name string, scores ...int) {
	fmt.Println("Student:", name)
	fmt.Println("Scores:", scores)
}

// ============================================
// 5. VARIADIC FUNCTION WITH STRINGS
// ============================================

func printNames(names ...string) {
	fmt.Println("Names:", names)
}

// ============================================
// 6. CHECKING NIL VARIADIC SLICES
// ============================================

func checkNumbers(numbers ...int) {
	fmt.Println("Numbers:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Is nil:", numbers == nil)
	fmt.Println()
}

// ============================================
// 7. NORMAL FUNCTION FOR COMPARISON
// ============================================

// This function accepts exactly two integers.
// It cannot accept just one or three arguments.

func addTwo(a, b int) int {
	return a + b
}
