
package main

import "fmt"

func main() {
    // ============================================================
    // 1. WHAT ARE DATA TYPES?
    // ============================================================

    // A data type defines what kind of value a variable can store.
    //
    // Go is a statically typed language.
    // Every variable has a type that is determined at compile time.
    //
    // Examples:
    // string  -> text
    // int     -> whole numbers
    // float64 -> decimal numbers
    // bool    -> true or false


    // ============================================================
    // 2. STRING TYPE
    // ============================================================

    // string stores a sequence of bytes, conventionally used for text.
    // String literals are written inside double quotes.

    var name string = "Niyaf"
    var city string = "Bengaluru"

    // Short declaration: Go infers the type as string.
    country := "India"

    fmt.Println(name, city, country)

    // Strings are immutable: their individual bytes cannot be changed.
    // To create modified text, create a new string.
    //
    // Example:
    message := "Hello"
    message = message + " Go"
    fmt.Println(message) // Hello Go

    // Single quotes are NOT used for Go strings.
    // 'A' is a rune literal, whereas "A" is a string literal.


    // ============================================================
    // 3. INTEGER TYPES
    // ============================================================

    // Integers store whole numbers without a fractional part.
    //
    // Signed integer types:
    // int8  -> 8 bits
    // int16 -> 16 bits
    // int32 -> 32 bits
    // int64 -> 64 bits
    // int   -> 32 or 64 bits, depending on architecture
    //
    // Unsigned integer types:
    // uint8  -> 8 bits
    // uint16 -> 16 bits
    // uint32 -> 32 bits
    // uint64 -> 64 bits
    // uint   -> architecture-dependent size
    //
    // Signed types support negative and positive values.
    // Unsigned types support zero and positive values only.

    var age int = 22
    var temperature int8 = -10
    var population int64 = 8000000
    var itemCount uint = 100
    var byteValue uint8 = 255

    fmt.Println(age, temperature, population, itemCount, byteValue)

    // Important:
    // int8 range:  -128 to 127
    // uint8 range: 0 to 255
    //
    // Assigning a constant outside a type's range causes a compile error.
    //
    // var invalid int8 = 200 // Compile-time error.


    // ============================================================
    // 4. BYTE TYPE
    // ============================================================

    // byte is an alias for uint8.
    // It is commonly used to represent individual bytes of data.

    var letter byte = 'A'

    fmt.Println(letter)         // 65: numeric byte value
    fmt.Printf("%c\n", letter)  // A: character formatting

    // Strings can be converted to byte slices when working with
    // their underlying UTF-8 bytes.
    bytes := []byte("Go")
    fmt.Println(bytes) // [71 111]


    // ============================================================
    // 5. RUNE TYPE
    // ============================================================

    // rune is an alias for int32.
    // It commonly represents a Unicode code point.
    //
    // Use rune when working with individual Unicode characters,
    // including characters outside basic ASCII.

    var character rune = 'A'
    var symbol rune = '₹'
    var emoji rune = '😊'

    fmt.Printf("%c\n", character)
    fmt.Printf("%c\n", symbol)
    fmt.Printf("%c\n", emoji)

    // Important:
    // A rune represents a Unicode code point, not necessarily
    // a complete user-perceived character (grapheme cluster).


    // ============================================================
    // 6. FLOATING-POINT TYPES
    // ============================================================

    // Floating-point types represent numbers with fractional parts.
    //
    // float32 -> 32-bit floating-point value
    // float64 -> 64-bit floating-point value
    //
    // Untyped decimal literals assigned using := are inferred as
    // float64 by default.

    var height float32 = 5.9
    var price float64 = 99.99

    rating := 4.5 // float64

    fmt.Println(height, price, rating)

    // Floating-point values have limited precision.
    // For exact monetary calculations, consider integer minor units
    // or a suitable decimal library instead of relying on float64.


    // ============================================================
    // 7. BOOLEAN TYPE
    // ============================================================

    // bool stores either true or false.
    // Boolean values are commonly used in conditions.

    var isLoggedIn bool = true
    var isAdmin bool = false

    hasAccess := isLoggedIn && isAdmin

    fmt.Println(isLoggedIn, isAdmin, hasAccess)

    // Go does not automatically treat numbers or strings as booleans.
    // Conditions must evaluate to a bool.
    //
    // Example:
    // if isLoggedIn { ... }


    // ============================================================
    // 8. COMPLEX NUMBER TYPES
    // ============================================================

    // Go supports complex numbers:
    // complex64  -> float32 real and imaginary components
    // complex128 -> float64 real and imaginary components
    //
    // The imaginary unit is written as i.

    var complexA complex64 = complex(1, 2)
    var complexB complex128 = 3 + 4i

    fmt.Println(complexA) // (1+2i)
    fmt.Println(complexB) // (3+4i)

    // Complex numbers are mainly useful in scientific,
    // mathematical, and engineering applications.


    // ============================================================
    // 9. UINTPTR TYPE
    // ============================================================

    // uintptr is an unsigned integer large enough to store the
    // bit pattern of a pointer.
    //
    // It is mainly used in low-level and unsafe programming.
    // It is not a general replacement for pointers.
    //
    // Example:
    var addressBits uintptr = 0
    fmt.Println(addressBits)


    // ============================================================
    // 10. ARRAY TYPE
    // ============================================================

    // An array stores a fixed number of elements of the same type.
    // Its length is part of its type.

    var numbers [3]int = [3]int{10, 20, 30}

    fmt.Println(numbers)
    fmt.Println(numbers[0]) // 10

    // Array length cannot change after declaration.
    //
    // [3]int and [4]int are different types.


    // ============================================================
    // 11. SLICE TYPE
    // ============================================================

    // A slice is a flexible-length view over an underlying array.
    // Slices are commonly used for collections in Go.

    scores := []int{80, 90, 100}

    scores = append(scores, 95)

    fmt.Println(scores) // [80 90 100 95]

    // Unlike arrays, slices can grow using append.
    // A slice's zero value is nil.
    //
    // var emptyScores []int


    // ============================================================
    // 12. MAP TYPE
    // ============================================================

    // A map stores key-value pairs.
    // All keys must have a comparable type.
    // All values in a map have the same declared value type.

    userAges := map[string]int{
        "Niyaf": 22,
        "Alex":  25,
    }

    userAges["Sam"] = 30

    fmt.Println(userAges)
    fmt.Println(userAges["Niyaf"]) // 22

    // A nil map can be read but cannot be used to insert entries.
    // Use make(...) or a map literal before adding entries.


    // ============================================================
    // 13. STRUCT TYPE
    // ============================================================

    // A struct groups related fields under one type.
    // Each field can have its own data type.

    type User struct {
        Name string
        Age  int
        Active bool
    }

    user := User{
        Name:   "Niyaf",
        Age:    22,
        Active: true,
    }

    fmt.Println(user.Name)
    fmt.Println(user.Age)
    fmt.Println(user.Active)


    // ============================================================
    // 14. POINTER TYPE
    // ============================================================

    // A pointer holds the address of a value.
    //
    // &variable -> obtains the address of a variable.
    // *pointer  -> accesses the value at that address.

    number := 100
    numberPointer := &number

    fmt.Println(numberPointer)  // Address (varies at runtime)
    fmt.Println(*numberPointer) // 100

    *numberPointer = 200

    fmt.Println(number) // 200

    // The zero value of a pointer is nil.


    // ============================================================
    // 15. INTERFACE TYPE
    // ============================================================

    // An interface type specifies a set of methods.
    // A value can satisfy an interface by implementing its methods.
    //
    // The empty interface, written as any, can hold a value of any type.
    // any is an alias for interface{}.

    var value any

    value = 100
    fmt.Println(value)

    value = "Hello"
    fmt.Println(value)

    value = true
    fmt.Println(value)

    // An interface holding a concrete value is not necessarily nil,
    // even if that concrete value is a nil pointer.


    // ============================================================
    // 16. CHANNEL TYPE
    // ============================================================

    // A channel is used for communication between goroutines.
    // Channels have a specific element type.

    numbersChannel := make(chan int, 1)

    numbersChannel <- 42

    receivedNumber := <-numbersChannel

    fmt.Println(receivedNumber) // 42

    // This buffered channel holds one integer.
    // The zero value of a channel is nil.


    // ============================================================
    // 17. ZERO VALUES
    // ============================================================

    // If a variable is declared without an initial value,
    // Go assigns the zero value for its type.
    //
    // Integer types -> 0
    // Floating-point types -> 0
    // bool -> false
    // string -> ""
    // Pointer -> nil
    // Slice -> nil
    // Map -> nil
    // Channel -> nil
    // Interface -> nil
    //
    // Arrays and structs receive zero values for each element or field.

    var defaultInt int
    var defaultFloat float64
    var defaultBool bool
    var defaultString string
    var defaultSlice []int
    var defaultMap map[string]int
    var defaultPointer *int

    fmt.Println(defaultInt)
    fmt.Println(defaultFloat)
    fmt.Println(defaultBool)
    fmt.Println(defaultString)
    fmt.Println(defaultSlice == nil)
    fmt.Println(defaultMap == nil)
    fmt.Println(defaultPointer == nil)


    // ============================================================
    // 18. TYPE CONVERSION
    // ============================================================

    // Go requires explicit conversions between different numeric types.
    // It does not automatically mix int and float64 in arithmetic.

    wholeNumber := 10
    decimalNumber := 2.5

    result := float64(wholeNumber) + decimalNumber

    fmt.Println(result) // 12.5

    // Converting a floating-point value to an integer truncates
    // the fractional part; it does not round to the nearest integer.

    decimal := 9.8
    converted := int(decimal)

    fmt.Println(converted) // 9

    // Conversions between strings and numbers require appropriate
    // conversions or standard-library functions such as strconv.


    // ============================================================
    // 19. CHECKING DATA TYPES
    // ============================================================

    // %T prints the type of a value.
    // %v prints the value in its default format.

    fmt.Printf("Type: %T, Value: %v\n", name, name)
    fmt.Printf("Type: %T, Value: %v\n", age, age)
    fmt.Printf("Type: %T, Value: %v\n", price, price)
    fmt.Printf("Type: %T, Value: %v\n", isDeveloper, isDeveloper)
}
