
package main

import "fmt"

func main() {
    // ============================================================
    // 1. VARIABLE DECLARATION USING var
    // ============================================================

    // Syntax:
    // var variableName dataType = value

    var name string = "Niyaf"
    var age int = 22
    var price float64 = 99.5
    var isDeveloper bool = true

    // Here, we explicitly specify the data type.
    // name is string, age is int, price is float64,
    // and isDeveloper is bool.


    // ============================================================
    // 2. VARIABLE DECLARATION USING var WITH TYPE INFERENCE
    // ============================================================

    // Syntax:
    // var variableName = value

    var city = "Bengaluru"
    var count = 100
    var rating = 4.5
    var isActive = false

    // We do not specify the data type explicitly.
    // Go infers the type from the assigned value:
    // city     -> string
    // count    -> int
    // rating   -> float64
    // isActive -> bool


    // ============================================================
    // 3. SHORT VARIABLE DECLARATION USING :=
    // ============================================================

    // Syntax:
    // variableName := value

    country := "India"
    score := 95
    percentage := 89.5
    isStudent := true

    // := declares a new variable and infers its type.
    // This is commonly used for local variables inside functions.
    //
    // IMPORTANT:
    // := can only be used inside a function.
    // It declares at least one new variable on the left-hand side.
    //
    // Example:
    // score := 95
    // score = 100  // Reassign an existing variable.
    // score := 100 // Error if score is the only variable:
    //              // no new variables on the left-hand side.


    // ============================================================
    // 4. DECLARING VARIABLES WITHOUT INITIAL VALUES
    // ============================================================

    // Syntax:
    // var variableName dataType

    var username string
    var total int
    var balance float64
    var loggedIn bool

    // Go automatically assigns a zero value when a variable
    // is declared without an initial value.
    //
    // string  -> ""
    // int     -> 0
    // float64 -> 0
    // bool    -> false
    //
    // You can assign values later:
    username = "niyaf_dev"
    total = 50
    balance = 1500.75
    loggedIn = true


    // ============================================================
    // 5. MULTIPLE VARIABLE DECLARATIONS
    // ============================================================

    // Declare multiple variables with the same data type.

    var firstName, lastName string = "Niyaf", "M"

    // Declare multiple variables with inferred types.

    var x, y, z = 10, "Go", true

    // Short declaration can also declare multiple variables.

    width, height := 100, 200

    // Each variable can have a different type when Go can infer it.


    // ============================================================
    // 6. GROUPED VARIABLE DECLARATION
    // ============================================================

    // When declaring many variables, you can group them
    // inside parentheses to keep the code organized.

    var (
        appName string = "Go Learning"
        version int    = 1
        enabled bool   = true
    )

    // This is useful when you have several related variables.


    // ============================================================
    // 7. CONSTANT DECLARATION USING const
    // ============================================================

    // Syntax:
    // const constantName = value
    // const constantName dataType = value

    const state string = "Karnataka"
    const maxUsers = 500
    const pi = 3.14159

    // Constants store values that cannot be reassigned.
    //
    // A constant must be initialized when declared.
    // Constants can be declared inside functions or at package level.
    // Constants can hold Boolean, string, and numeric values.
    // Constants cannot be declared using :=.
    //
    // maxUsers = 1000 // Error: cannot assign to a constant.
    //
    // Unlike variables, Go constants can also be untyped.
    // For example, maxUsers and pi above are untyped constants
    // until a type is required by their use.


    // ============================================================
    // 8. VARIABLE REASSIGNMENT
    // ============================================================

    // Use = to change the value of an existing variable.
    // Reassignment does not declare a new variable.

    age = 25
    city = "Mumbai"
    isDeveloper = false

    // A variable's type does not change when reassigned.
    //
    // age = "twenty-five" // Error: string cannot be assigned to int.
    //
    // Remember:
    // := declares a variable.
    // = assigns or reassigns a value.


    // ============================================================
    // 9. BASIC DATA TYPES
    // ============================================================

    var textValue string = "Hello Go"
    var integerValue int = 100
    var floatValue float64 = 99.99
    var booleanValue bool = true

    // string  -> text
    // int     -> whole numbers (size depends on architecture)
    // float32 -> 32-bit floating-point numbers
    // float64 -> 64-bit floating-point numbers
    // bool    -> true or false
    //
    // Other integer types:
    // int8, int16, int32, int64
    // uint, uint8, uint16, uint32, uint64
    //
    // byte is an alias for uint8.
    // rune is an alias for int32, commonly used for Unicode code points.


    // ============================================================
    // 10. CHECKING VARIABLE TYPES
    // ============================================================

    // Use fmt.Printf with %T to print the type of a value.

    fmt.Printf("Type of name: %T\n", name)
    fmt.Printf("Type of age: %T\n", age)
    fmt.Printf("Type of price: %T\n", price)
    fmt.Printf("Type of isDeveloper: %T\n", isDeveloper)
    fmt.Printf("Type of score: %T\n", score)


    // ============================================================
    // 11. PRINTING VARIABLE VALUES
    // ============================================================

    // fmt.Println prints values with spaces between arguments.

    fmt.Println(name, age, price, isDeveloper)
    fmt.Println(city, count, rating, isActive)
    fmt.Println(country, score, percentage, isStudent)
    fmt.Println(username, total, balance, loggedIn)
    fmt.Println(firstName, lastName)
    fmt.Println(x, y, z)
    fmt.Println(width, height)
    fmt.Println(appName, version, enabled)
    fmt.Println(state, maxUsers, pi)
    fmt.Println(textValue, integerValue, floatValue, booleanValue)
}
