
package main

import "fmt"

func main() {

	// ==================================================
	// 1. BASIC IF STATEMENT
	// ==================================================

	// An if statement executes its block only when the
	// condition evaluates to true.

	age := 22

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// JavaScript equivalent:
	// if (age >= 18) {
	//     console.log("You are an adult");
	// }

	// IMPORTANT:
	// Go does not require parentheses around conditions.
	// Curly braces {} are mandatory.


	// ==================================================
	// 2. IF - ELSE
	// ==================================================

	age = 16

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You cannot vote yet")
	}

	// If the condition is true, the if block executes.
	// Otherwise, the else block executes.
	//
	// Exactly one of these two blocks executes.


	// ==================================================
	// 3. IF - ELSE IF - ELSE
	// ==================================================

	marks := 75

	if marks >= 90 {
		fmt.Println("Grade A")
	} else if marks >= 75 {
		fmt.Println("Grade B")
	} else if marks >= 50 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Fail")
	}

	// Go checks conditions from top to bottom.
	// The FIRST matching block executes.
	// Once a matching block executes, the remaining
	// else-if and else blocks are skipped.
	//
	// Here, marks = 75, so Grade B is printed.
	//
	// IMPORTANT:
	// Check the highest threshold first.
	// If you check marks >= 50 first, marks = 75
	// will match it, and the later conditions won't run.


	// ==================================================
	// 4. IF WITH A LOGICAL AND (&&)
	// ==================================================

	age = 25
	hasID := true

	if age >= 18 && hasID {
		fmt.Println("Entry allowed")
	}

	// && means AND.
	// BOTH conditions must be true.
	//
	// JavaScript: age >= 18 && hasID
	// Go:         age >= 18 && hasID


	// ==================================================
	// 5. IF WITH A LOGICAL OR (||)
	// ==================================================

	day := "Saturday"

	if day == "Saturday" || day == "Sunday" {
		fmt.Println("It's the weekend")
	}

	// || means OR.
	// At least ONE condition must be true.


	// ==================================================
	// 6. IF WITH LOGICAL NOT (!)
	// ==================================================

	isLoggedIn := false

	if !isLoggedIn {
		fmt.Println("Please log in")
	}

	// ! reverses a boolean value:
	// !true  becomes false
	// !false becomes true
	//
	// JavaScript also uses ! for logical NOT.


	// ==================================================
	// 7. COMBINING CONDITIONS
	// ==================================================

	userAge := 22
	isVerified := true
	hasBlockedAccount := false

	if userAge >= 18 && isVerified && !hasBlockedAccount {
		fmt.Println("User can continue")
	}

	// Read this as:
	// age is at least 18
	// AND user is verified
	// AND user is NOT blocked.


	// ==================================================
	// 8. IF WITH A SHORT STATEMENT
	// ==================================================

	// Go allows a short statement before the condition.
	// The variable declared here is limited to the
	// scope of the if statement and its else blocks.

	if score := 85; score >= 50 {
		fmt.Println("Passed with score:", score)
	} else {
		fmt.Println("Failed with score:", score)
	}

	// Structure:
	// if initialization; condition {
	//     ...
	// }
	//
	// score is available inside the if and else blocks.
	// score is NOT available after the entire if statement.
	//
	// JavaScript does not have this exact if syntax.


	// ==================================================
	// 9. NESTED IF
	// ==================================================

	age = 25
	isMember := true

	if age >= 18 {
		if isMember {
			fmt.Println("Adult member: discount available")
		} else {
			fmt.Println("Adult non-member")
		}
	} else {
		fmt.Println("Not an adult")
	}

	// A nested if is an if statement inside another if.
	// The inner condition is checked only if execution
	// reaches the inner if statement.
	//
	// Avoid unnecessary nesting when a simpler condition
	// using && would be easier to understand.


	// ==================================================
	// 10. COMPARISON OPERATORS
	// ==================================================

	a := 10
	b := 20

	fmt.Println(a == b) // false: equal to
	fmt.Println(a != b) // true: not equal to
	fmt.Println(a < b)  // true: less than
	fmt.Println(a > b)  // false: greater than
	fmt.Println(a <= b) // true: less than or equal
	fmt.Println(a >= b) // false: greater than or equal

	// These comparisons produce a bool: true or false.
	//
	// IMPORTANT:
	// =  is assignment.
	// == is comparison.
	//
	// Go does not allow assigning a value inside an
	// if condition the way some languages might.


	// ==================================================
	// 11. SWITCH STATEMENT
	// ==================================================

	// Use switch when you want to compare one value
	// against several possible cases.

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("Apple selected")
	case "banana":
		fmt.Println("Banana selected")
	case "mango":
		fmt.Println("Mango selected")
	default:
		fmt.Println("Unknown fruit")
	}

	// Only the matching case executes.
	// default is optional and runs when no case matches.
	//
	// Unlike JavaScript switch, Go does not automatically
	// fall through to the next case.
	// You generally do not need break in each Go case.


	// ==================================================
	// 12. SWITCH WITH MULTIPLE VALUES
	// ==================================================

	day = "Sunday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	default:
		fmt.Println("Invalid day")
	}

	// Multiple values can share the same case.
	// This avoids writing multiple else-if conditions.


	// ==================================================
	// 13. SWITCH WITHOUT AN EXPRESSION
	// ==================================================

	// Go also allows switch without a value.
	// Each case can be a boolean condition.

	marks = 82

	switch {
	case marks >= 90:
		fmt.Println("Grade A")
	case marks >= 75:
		fmt.Println("Grade B")
	case marks >= 50:
		fmt.Println("Grade C")
	default:
		fmt.Println("Fail")
	}

	// This is similar to an else-if chain.
	// Cases are checked from top to bottom.
	// The first true case executes.


	// ==================================================
	// 14. SWITCH WITH INITIALIZATION
	// ==================================================

	switch number := 2; number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Another number")
	}

	// Like if, switch can have an initialization
	// statement before the value being checked.
	// number is scoped to this switch statement.


	// ==================================================
	// 15. TYPE SWITCH
	// ==================================================

	// A type switch checks the actual dynamic type stored
	// inside an interface value.

	var value any = "Hello Go"

	switch v := value.(type) {
	case string:
		fmt.Println("String value:", v)
	case int:
		fmt.Println("Integer value:", v)
	case bool:
		fmt.Println("Boolean value:", v)
	default:
		fmt.Printf("Other type: %T\n", v)
	}

	// any is an alias for interface{}.
	// A type switch is useful when a value can hold
	// different types and you need to handle them.
	//
	// value.(type) is special syntax valid in a type switch.


	// ==================================================
	// 16. EARLY RETURN WITH IF
	// ==================================================

	// In functions, you can handle invalid input early
	// and return before doing the main work.

	checkAge := 15

	if checkAge < 18 {
		fmt.Println("Access denied")
	} else {
		fmt.Println("Access granted")
	}

	// In a real function, the first branch might instead
	// return an error. See the example below.
}

// ==================================================
// 17. IF WITH ERROR HANDLING (COMMON GO PATTERN)
// ==================================================

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}
