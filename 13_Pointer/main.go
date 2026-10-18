package main

import "fmt"

// ============================================
// 1. WHAT IS A POINTER?
// ============================================

// A pointer is a variable that stores the memory address
// of another variable.
//
// Normally, when we use a variable, we access its value.
//
// A pointer allows us to access the memory address where
// that value is stored and indirectly access or modify it.
//
// Example:
//
// age := 25
// ptr := &age
//
// age stores the value 25.
// ptr stores the memory address of age.
//
// Two important operators:
//
// &  -> Gets the memory address of a variable.
// *  -> Accesses the value stored at the address.
//
// Remember:
//
// &age means "address of age".
// *ptr means "value at the address stored in ptr".

func main() {

	// ============================================
	// 2. NORMAL VARIABLE
	// ============================================

	// A normal variable stores a value.

	age := 25

	fmt.Println(age)

	// Output:
	// 25

	// ============================================
	// 3. GET THE MEMORY ADDRESS USING &
	// ============================================

	// The & operator gives us the memory address
	// of a variable.

	fmt.Println(&age)

	// Example output:
	// 0xc000012090
	//
	// IMPORTANT:
	// Your actual address will probably be different.
	//
	// Explanation:
	// &age gives the memory address where age is stored.
	// The address is displayed in hexadecimal format.

	// ============================================
	// 4. CREATE A POINTER
	// ============================================

	// We can store a variable's address inside a pointer.
	//
	// The &age expression gives the address of age.
	//
	// ptr now points to age.

	ptr := &age

	fmt.Println(ptr)

	// Output:
	// A memory address, for example:
	// 0xc000012090
	//
	// Explanation:
	// ptr stores the address of age.
	//
	// Therefore:
	// ptr == &age
	//
	// The actual address is determined at runtime.

	// ============================================
	// 5. GET THE VALUE USING *
	// ============================================

	// The * operator can dereference a pointer.
	//
	// Dereferencing means accessing the value
	// stored at the address the pointer points to.

	fmt.Println(*ptr)

	// Output:
	// 25
	//
	// Explanation:
	// ptr contains the address of age.
	//
	// *ptr accesses the value at that address.
	//
	// Since age is 25, *ptr is also 25.

	// ============================================
	// 6. MODIFY A VARIABLE USING A POINTER
	// ============================================

	// A pointer can be used to modify the original
	// variable without directly using its name.
	//
	// Here, ptr points to age.
	//
	// Assigning a new value to *ptr changes age itself.

	*ptr = 30

	fmt.Println(age)
	fmt.Println(*ptr)

	// Output:
	// 30
	// 30
	//
	// Explanation:
	// *ptr = 30 means:
	// "Store 30 at the memory address ptr points to."
	//
	// Since ptr points to age, age becomes 30.
	//
	// Both age and *ptr now show 30.

	// ============================================
	// 7. POINTER TYPE
	// ============================================

	// A pointer has a type based on the value it points to.
	//
	// For example:
	//
	// *int    -> pointer to an int
	// *string -> pointer to a string
	// *bool   -> pointer to a bool
	//
	// Here, &age has the type *int because age is an int.

	fmt.Printf("%T\n", &age)
	fmt.Printf("%T\n", ptr)

	// Output:
	// *int
	// *int
	//
	// Explanation:
	// %T prints the type of a value.
	//
	// Both &age and ptr are pointers to integers.

	// ============================================
	// 8. POINTER TO A STRING
	// ============================================

	// Pointers work with strings too.

	name := "Niyaf"

	namePtr := &name

	fmt.Println(name)
	fmt.Println(*namePtr)

	*namePtr = "Rahul"

	fmt.Println(name)

	// Output:
	// Niyaf
	// Niyaf
	// Rahul
	//
	// Explanation:
	// namePtr stores the address of name.
	//
	// *namePtr accesses the string at that address.
	//
	// Updating *namePtr changes the original name variable.

	// ============================================
	// 9. NIL POINTER
	// ============================================

	// A pointer can contain no address.
	//
	// Go represents this using nil.
	//
	// The zero value of a pointer is nil.

	var numberPtr *int

	fmt.Println(numberPtr)

	// Output:
	// <nil>
	//
	// Explanation:
	// numberPtr does not point to an integer variable.
	//
	// It is valid to declare a nil pointer.
	//
	// But we must not dereference it before it points
	// to a valid value.

	if numberPtr == nil {
		fmt.Println("Pointer is nil")
	}

	// Output:
	// Pointer is nil
	//
	// IMPORTANT:
	// Uncommenting the next line causes a runtime panic:
	//
	// fmt.Println(*numberPtr)
	//
	// Why?
	// Because there is no valid value to access.

	// ============================================
	// 10. POINTERS AND FUNCTIONS
	// ============================================

	// Go passes arguments to functions by value.
	//
	// If we pass an ordinary integer, the function
	// receives a copy of that integer.
	//
	// Changes to the copy do not change the original.

	number := 10

	changeValue(number)

	fmt.Println(number)

	// Output:
	// 10
	//
	// Explanation:
	// changeValue receives a copy of number.
	//
	// Changing that copy does not modify the original.

	// ============================================
	// 11. MODIFY A VARIABLE THROUGH A POINTER PARAMETER
	// ============================================

	// To modify the caller's variable, we can pass
	// its address to a function.
	//
	// &number gives the address of number.
	//
	// changeValueUsingPointer receives that address.

	changeValueUsingPointer(&number)

	fmt.Println(number)

	// Output:
	// 100
	//
	// Explanation:
	// The function receives a pointer to number.
	//
	// Inside the function, *ptr = 100 changes the
	// original variable through that pointer.
	//
	// Go still passes the pointer itself by value,
	// but the copied pointer refers to the same variable.

	// ============================================
	// 12. POINTER PARAMETER SYNTAX
	// ============================================

	// In this function:
	//
	// func changeValueUsingPointer(ptr *int)
	//
	// ptr *int means ptr is a pointer to an integer.
	//
	// The expression:
	//
	// *ptr = 100
	//
	// changes the integer at the address ptr stores.
	//
	// Remember the difference:
	//
	// ptr  -> the memory address
	// *ptr -> the value at that address

	// ============================================
	// 13. SWAP TWO VALUES USING POINTERS
	// ============================================

	// We can use pointers to modify two variables
	// from inside a function.

	a := 10
	b := 20

	fmt.Println("Before swap:", a, b)

	swap(&a, &b)

	fmt.Println("After swap:", a, b)

	// Output:
	// Before swap: 10 20
	// After swap: 20 10
	//
	// Explanation:
	// We pass the addresses of a and b.
	//
	// Inside swap, we exchange the values
	// at those addresses.
	//
	// Therefore, the original variables are changed.

	// ============================================
	// 14. POINTERS AND ARRAYS
	// ============================================

	// An array has a fixed length.
	//
	// We can take a pointer to the entire array.
	//
	// The type *[3]int means a pointer to an array
	// containing exactly three integers.

	nums := [3]int{10, 20, 30}

	arrayPtr := &nums

	fmt.Println(nums)
	fmt.Println(*arrayPtr)

	// Output:
	// [10 20 30]
	// [10 20 30]
	//
	// Explanation:
	// arrayPtr points to the entire array nums.
	//
	// *arrayPtr accesses the array through that pointer.

	// ============================================
	// 15. MODIFY AN ARRAY THROUGH A POINTER
	// ============================================

	// Because arrayPtr points to nums, we can modify
	// an array element through the pointer.

	arrayPtr[0] = 100

	fmt.Println(nums)

	// Output:
	// [100 20 30]
	//
	// Explanation:
	// arrayPtr[0] modifies the first element of the
	// array that arrayPtr points to.
	//
	// Go allows this convenient syntax.
	// Writing (*arrayPtr)[0] would also access
	// the first element explicitly.

	// ============================================
	// 16. POINTERS AND SLICES
	// ============================================

	// A slice already refers to an underlying array.
	//
	// Copying a slice copies its slice descriptor,
	// not all of its underlying elements.
	//
	// Therefore, two slices can refer to the same
	// underlying array.

	numsSlice := []int{10, 20, 30}

	anotherSlice := numsSlice

	anotherSlice[0] = 999

	fmt.Println(numsSlice)
	fmt.Println(anotherSlice)

	// Output:
	// [999 20 30]
	// [999 20 30]
	//
	// Explanation:
	// Both slices refer to the same underlying array.
	//
	// We do not need a pointer to a slice to modify
	// an existing element of that underlying array.
	//
	// A pointer to a slice can still be useful when
	// a function needs to replace the caller's slice
	// descriptor, such as by assigning a new slice.

	// ============================================
	// 17. POINTER TO A STRUCT
	// ============================================

	// Pointers are commonly used with structs.
	//
	// A struct groups related fields into one type.
	//
	// We can take the address of a struct variable
	// and use the pointer to access its fields.

	type Person struct {
		Name string
		Age  int
	}

	person := Person{
		Name: "Niyaf",
		Age:  22,
	}

	personPtr := &person

	fmt.Println(personPtr.Name)
	fmt.Println(personPtr.Age)

	// Output:
	// Niyaf
	// 22
	//
	// Explanation:
	// Go automatically allows personPtr.Name as
	// shorthand for (*personPtr).Name.
	//
	// Both expressions access the Name field of
	// the struct that personPtr points to.

	// ============================================
	// 18. MODIFY A STRUCT THROUGH A POINTER
	// ============================================

	// We can change the fields of the original
	// struct using its pointer.

	personPtr.Age = 23

	fmt.Println(person.Age)

	// Output:
	// 23
	//
	// Explanation:
	// personPtr points to person.
	//
	// Updating personPtr.Age changes person.Age.

	// ============================================
	// 19. POINTERS AND NEW()
	// ============================================

	// The built-in new function allocates zero-valued
	// storage for a value of the requested type
	// and returns a pointer to it.
	//
	// Example:
	//
	// new(int) returns *int.
	//
	// The integer initially has the value 0.

	ptrNumber := new(int)

	fmt.Println(*ptrNumber)

	// Output:
	// 0
	//
	// Explanation:
	// ptrNumber is a pointer to an integer.
	//
	// The pointed-to integer starts at its zero value,
	// which is 0.

	*ptrNumber = 500

	fmt.Println(*ptrNumber)

	// Output:
	// 500
	//
	// Explanation:
	// We assign 500 to the value at that address.

	// ============================================
	// 20. POINTERS DO NOT SUPPORT ARITHMETIC
	// ============================================

	// In Go, you cannot perform ordinary pointer
	// arithmetic like adding 1 to a pointer.
	//
	// For example, this is invalid:
	//
	// ptr := &number
	// ptr = ptr + 1
	//
	// Go does not allow that operation for ordinary
	// pointers.
	//
	// This restriction helps make pointer usage safer.
	//
	// To move through a collection, use indexes,
	// loops, arrays, or slices instead.

	// ============================================
	// 21. IMPORTANT POINTER RULES
	// ============================================

	// 1. &variable gives the address of a variable.
	//
	// 2. *pointer accesses the value at that address.
	//
	// 3. A pointer's zero value is nil.
	//
	// 4. Dereferencing a nil pointer causes a panic.
	//
	// 5. Go passes arguments by value, including pointers.
	//
	// 6. Passing a pointer lets a function modify
	//    the value that the pointer refers to.
	//
	// 7. Arrays and slices are different:
	//    arrays are copied when passed by value;
	//    slices contain a descriptor referring to
	//    an underlying array.
	//
	// 8. Go manages memory automatically.
	//    You generally do not manually free memory.
}

// ============================================
// FUNCTION DEFINITIONS
// ============================================

// ============================================
// 1. NORMAL FUNCTION
// ============================================

// This function receives an int by value.
// The parameter is a copy of the caller's value.

func changeValue(value int) {
	value = 100
}

// ============================================
// 2. FUNCTION THAT ACCEPTS A POINTER
// ============================================

// ptr *int means ptr is a pointer to an integer.
//
// Dereferencing ptr lets us modify the integer
// at the address stored in ptr.

func changeValueUsingPointer(ptr *int) {
	*ptr = 100
}

// ============================================
// 3. SWAP FUNCTION
// ============================================

// This function receives pointers to two integers.
//
// It exchanges the values at those addresses.

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
