package main

import "fmt"

// ============================================
// 1. WHAT IS A CLOSURE?
// ============================================

// A closure is a function that can access variables
// from the surrounding function where it was created.
//
// Even after the surrounding function has finished
// executing, the closure can still access those variables.
//
// In simple words:
//
// A closure = a function + the variables it remembers.
//
// Closures are useful when we want a function to remember
// some data between multiple function calls.
//
// In Go, anonymous functions can create closures.

func main() {

	// ============================================
	// 2. BASIC CLOSURE
	// ============================================

	// Here, message is a variable inside main.
	//
	// The anonymous function below can access message
	// even though message was not passed as a parameter.
	//
	// This function is a closure because it uses a
	// variable from its surrounding scope.

	message := "Hello from outside the function"

	sayMessage := func() {
		fmt.Println(message)
	}

	sayMessage()

	// Output:
	// Hello from outside the function
	//
	// Explanation:
	// sayMessage is an anonymous function stored
	// in a variable.
	//
	// It can access message because message belongs
	// to the surrounding scope.

	// ============================================
	// 3. CLOSURE THAT MODIFIES A VARIABLE
	// ============================================

	// A closure can not only read an outside variable.
	// It can also modify that variable.
	//
	// Here, count is created inside main.
	// The closure increments count every time it runs.

	count := 0

	increment := func() {
		count++
		fmt.Println(count)
	}

	increment()
	increment()
	increment()

	// Output:
	// 1
	// 2
	// 3
	//
	// Explanation:
	// First call:  count becomes 1.
	// Second call: count becomes 2.
	// Third call:  count becomes 3.
	//
	// The closure accesses and modifies the same count
	// variable each time it is called.

	// ============================================
	// 4. CLOSURE THAT REMEMBERS ITS OWN STATE
	// ============================================

	// This is one of the most important uses of closures.
	//
	// createCounter returns a function.
	//
	// Inside createCounter, count is initialized to 0.
	// The returned function increments count.
	//
	// Even after createCounter finishes, the returned
	// function can still access count.

	counter := createCounter()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// Output:
	// 1
	// 2
	// 3
	//
	// Explanation:
	// createCounter() runs once and initializes count = 0.
	//
	// counter stores the function returned by createCounter.
	//
	// First call:  count becomes 1.
	// Second call: count becomes 2.
	// Third call:  count becomes 3.
	//
	// Normally, local variables are associated with
	// a function call.
	//
	// But the closure still needs count, so Go keeps
	// the variable alive for as long as it is needed.
	//
	// The closure remembers the variable, not just
	// the value that the variable had when created.

	// ============================================
	// 5. EACH CLOSURE CAN HAVE ITS OWN STATE
	// ============================================

	// Let's create two independent counters.
	//
	// Each call to createCounter creates a new count
	// variable for that particular closure.

	counterA := createCounter()
	counterB := createCounter()

	fmt.Println(counterA())
	fmt.Println(counterA())
	fmt.Println(counterB())
	fmt.Println(counterA())
	fmt.Println(counterB())

	// Output:
	// 1
	// 2
	// 1
	// 3
	// 2
	//
	// Explanation:
	// counterA has its own count variable.
	// counterB has a different count variable.
	//
	// counterA:
	// 1 -> 2 -> 3
	//
	// counterB:
	// 1 -> 2
	//
	// Calling counterA does not change counterB's count.

	// ============================================
	// 6. CLOSURE WITH A PARAMETER
	// ============================================

	// A closure can also use regular parameters.
	//
	// Here, createMultiplier receives a number.
	// It returns a function that multiplies another
	// number by the remembered multiplier.

	double := createMultiplier(2)
	triple := createMultiplier(3)

	fmt.Println(double(5))
	fmt.Println(double(10))
	fmt.Println(triple(5))

	// Output:
	// 10
	// 20
	// 15
	//
	// Explanation:
	// double remembers multiplier = 2.
	// triple remembers multiplier = 3.
	//
	// double(5) means 5 * 2 = 10.
	// double(10) means 10 * 2 = 20.
	// triple(5) means 5 * 3 = 15.
	//
	// The returned functions keep access to the
	// multiplier variable from createMultiplier.

	// ============================================
	// 7. CLOSURE VS FUNCTION PARAMETERS
	// ============================================

	// Compare these two approaches.
	//
	// Normal function:
	// We pass both values whenever we call it.
	//
	// Closure:
	// We can remember one value and pass only the
	// changing value.

	fmt.Println(multiply(2, 5))
	fmt.Println(multiply(2, 10))

	doubleNumber := createMultiplier(2)

	fmt.Println(doubleNumber(5))
	fmt.Println(doubleNumber(10))

	// Output:
	// 10
	// 20
	// 10
	// 20
	//
	// Explanation:
	// multiply needs both arguments every time.
	//
	// doubleNumber already remembers 2.
	// Therefore, we only pass the number to multiply.
	//
	// This is useful when the same configuration or
	// value needs to be reused repeatedly.

	// ============================================
	// 8. CLOSURE WITH A STRING
	// ============================================

	// Closures can remember any kind of variable,
	// including strings.
	//
	// Here, createGreeter remembers a person's name.

	greetNiyaf := createGreeter("Niyaf")
	greetRahul := createGreeter("Rahul")

	greetNiyaf()
	greetRahul()

	// Output:
	// Hello, Niyaf
	// Hello, Rahul
	//
	// Explanation:
	// greetNiyaf remembers "Niyaf".
	// greetRahul remembers "Rahul".
	//
	// Each returned function uses its own remembered name.

	// ============================================
	// 9. CLOSURE WITH A SLICE
	// ============================================

	// A closure can also access and modify a slice.
	//
	// Here, createAppender remembers the items slice.
	// Each call appends a new item to the same slice.

	appendItem := createAppender()

	appendItem("Go")
	appendItem("JavaScript")
	appendItem("Rust")

	// Output:
	// [Go]
	// [Go JavaScript]
	// [Go JavaScript Rust]
	//
	// Explanation:
	// The items slice is created inside createAppender.
	//
	// The returned closure continues using that slice
	// across multiple calls.

	// ============================================
	// 10. CLOSURES AND LOOPS
	// ============================================

	// Be careful when creating closures inside loops.
	//
	// In Go 1.22 and later, variables declared by
	// a for loop using := are new for each iteration.
	//
	// This helps prevent a common closure bug.
	//
	// Here, each function remembers its own value.

	var functions []func()

	for i := 1; i <= 3; i++ {
		functions = append(functions, func() {
			fmt.Println(i)
		})
	}

	for _, function := range functions {
		function()
	}

	// Output in Go 1.22+:
	// 1
	// 2
	// 3
	//
	// Note:
	// This example relies on the Go 1.22+ loop-variable
	// semantics, which apply when the module's Go version
	// is 1.22 or later.
	//
	// In older language versions, these closures can
	// all observe the same loop variable after the loop.

	// ============================================
	// 11. PRACTICAL EXAMPLE: DISCOUNT CALCULATOR
	// ============================================

	// Imagine an application that needs to calculate
	// prices using a particular discount percentage.
	//
	// Instead of passing the discount every time,
	// we can create a closure that remembers it.

	discount10 := createDiscount(10)
	discount20 := createDiscount(20)

	fmt.Println(discount10(1000))
	fmt.Println(discount20(1000))

	// Output:
	// 900
	// 800
	//
	// Explanation:
	// discount10 remembers a 10% discount.
	// discount20 remembers a 20% discount.
	//
	// Price 1000 with 10% discount = 900.
	// Price 1000 with 20% discount = 800.
}

// ============================================
// FUNCTION DEFINITIONS
// ============================================

// ============================================
// 1. CREATE A COUNTER CLOSURE
// ============================================

// This function returns another function.
//
// The returned function has type func() int.
//
// count belongs to createCounter, but the returned
// closure continues to access it after createCounter
// finishes executing.

func createCounter() func() int {
	count := 0

	return func() int {
		count++

		return count
	}
}

// ============================================
// 2. CREATE A MULTIPLIER CLOSURE
// ============================================

// This function remembers the multiplier value.
//
// It returns a function that accepts another integer
// and returns the product.

func createMultiplier(multiplier int) func(int) int {
	return func(number int) int {
		return number * multiplier
	}
}

// ============================================
// 3. NORMAL MULTIPLY FUNCTION
// ============================================

// This is a regular function.
// It does not remember any values between calls.

func multiply(a, b int) int {
	return a * b
}

// ============================================
// 4. CREATE A GREETER CLOSURE
// ============================================

// The returned function remembers the name parameter.

func createGreeter(name string) func() {
	return func() {
		fmt.Println("Hello,", name)
	}
}

// ============================================
// 5. CREATE AN APPENDER CLOSURE
// ============================================

// This function creates a slice and returns a function
// that appends values to it.
//
// The same items slice is reused by the returned closure.

func createAppender() func(string) {
	items := []string{}

	return func(item string) {
		items = append(items, item)

		fmt.Println(items)
	}
}

// ============================================
// 6. CREATE A DISCOUNT CLOSURE
// ============================================

// The returned function remembers the discount percent.
//
// Example:
// createDiscount(10) returns a function that applies
// a 10% discount to any price passed to it.

func createDiscount(percent float64) func(float64) float64 {
	return func(price float64) float64 {
		return price - (price * percent / 100)
	}
}
