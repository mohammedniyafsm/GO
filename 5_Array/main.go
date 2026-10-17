package main

import "fmt"

func main() {
	// =========================================================
	// 1. WHAT IS AN ARRAY?
	// =========================================================
	// An array stores a fixed number of values of the SAME type.
	// Its length is fixed and is part of its type.
	//
	// JavaScript:
	// const numbers = [10, 20, 30];
	//
	// Go:
	var numbers [3]int
	fmt.Println("Zero-value array:", numbers) // [0 0 0]
	// [3]int means "an array of exactly 3 integers".
	// Since int's zero value is 0, every element starts as 0.

	// =========================================================
	// 2. DECLARING AND INITIALIZING ARRAYS
	// =========================================================
	var scores [3]int
	scores[0] = 80
	scores[1] = 90
	scores[2] = 100
	fmt.Println("Assigned scores:", scores)

	// Declare and initialize with an array literal.
	prices := [3]float64{10.5, 20.0, 30.25}
	fmt.Println("Prices:", prices)

	// Let Go count the elements using ...
	fruits := [...]string{"apple", "banana", "mango"}
	fmt.Println("Fruits:", fruits)
	fmt.Println("Fruits length:", len(fruits)) // 3
	// [...] is allowed in an array literal. Go counts elements
	// and infers the fixed length. The resulting type is [3]string.

	// =========================================================
	// 3. ARRAY LENGTH IS PART OF THE TYPE
	// =========================================================
	var threeItems [3]int
	var fourItems [4]int
	fmt.Println("Different array lengths:", threeItems, fourItems)
	// [3]int and [4]int are DIFFERENT types.
	// You cannot directly assign one to the other:
	// threeItems = fourItems // compile error: different types
	//
	// This is a major difference from JavaScript arrays, which
	// can grow or shrink dynamically.

	// =========================================================
	// 4. ACCESSING ELEMENTS WITH INDEXES
	// =========================================================
	letters := [4]string{"A", "B", "C", "D"}
	fmt.Println("First:", letters[0]) // A
	fmt.Println("Last:", letters[3])  // D
	// Indexes start at 0.
	// For an array of length N, valid indexes are 0 through N-1.
	// Accessing letters[4] causes a runtime panic.

	// =========================================================
	// 5. CHANGING AN ELEMENT
	// =========================================================
	letters[1] = "Bee"
	fmt.Println("After changing index 1:", letters)
	// Arrays are mutable: individual elements can be changed.
	// The array length, however, cannot be changed.

	// =========================================================
	// 6. len() AND THE LAST INDEX
	// =========================================================
	values := [5]int{11, 22, 33, 44, 55}
	fmt.Println("Length:", len(values))                 // 5
	fmt.Println("Last index:", len(values)-1)           // 4
	fmt.Println("Last element:", values[len(values)-1]) // 55
	// len(array) returns the number of elements, not the last index.
	// The last index is len(array) - 1.

	// =========================================================
	// 7. LOOP THROUGH AN ARRAY WITH A CLASSIC FOR LOOP
	// =========================================================
	for i := 0; i < len(values); i++ {
		fmt.Println("Index and value:", i, values[i])
	}
	// i starts at 0. The loop continues while i < len(values).
	// It stops before i reaches 5, so it never reads values[5].

	// =========================================================
	// 8. LOOP THROUGH AN ARRAY WITH range
	// =========================================================
	for index, value := range values {
		fmt.Println("range:", index, value)
	}
	// range gives the index and a COPY of the element's value.
	// Use _ if you do not need one of the values:
	for _, value := range values {
		fmt.Println("Value only:", value)
	}
	// This is similar to JavaScript's:
	// for (const value of values) { console.log(value); }
	// Go's range over an array iterates over each array element.

	// =========================================================
	// 9. IMPORTANT: range VALUES ARE COPIES
	// =========================================================
	original := [3]int{1, 2, 3}
	for _, value := range original {
		value = value * 10
	}
	fmt.Println("Array after changing range variable:", original)
	// Output remains [1 2 3].
	// The variable value is a copy, so changing it doesn't change
	// the element stored in the array.

	// To update elements, use indexes:
	for i := range original {
		original[i] = original[i] * 10
	}
	fmt.Println("Array after changing by index:", original)
	// Output: [10 20 30].
	// for i := range original gives indexes in modern Go versions.

	// =========================================================
	// 10. ARRAY LITERAL WITH SELECTED INDEXES
	// =========================================================
	// You can initialize specific indexes. Unspecified elements
	// receive their zero value.
	selected := [5]int{0: 100, 3: 400}
	fmt.Println("Selected indexes:", selected)
	// Output: [100 0 0 400 0]

	// =========================================================
	// 11. ARRAYS WITH DIFFERENT TYPES
	// =========================================================
	var integers [3]int
	var names [2]string
	var flags [2]bool
	fmt.Println("Typed zero values:", integers, names, flags)
	// Each array has one element type:
	// [3]int, [2]string, [2]bool.
	// One array cannot directly mix ints, strings, and bools.

	// =========================================================
	// 12. ARRAYS ARE VALUE TYPES: ASSIGNMENT COPIES THE ARRAY
	// =========================================================
	first := [3]int{10, 20, 30}
	second := first
	second[0] = 999
	fmt.Println("first:", first)   // [10 20 30]
	fmt.Println("second:", second) // [999 20 30]
	// Assigning an array copies all its elements.
	// Changing second does not change first.
	// This differs from JavaScript arrays, where assigning
	// another variable copies the reference to the same array.

	// =========================================================
	// 13. PASSING ARRAYS TO FUNCTIONS ALSO COPIES THEM
	// =========================================================
	// See printArrayCopy below. The function receives a copy.
	printArrayCopy(first)
	fmt.Println("first after function call:", first)
	// first remains unchanged because the function parameter
	// is a separate array value.

	// If a function needs to modify the original array, pass a
	// pointer to it, as shown in modifyArray below.
	modifyArray(&first)
	fmt.Println("first after pointer function:", first)

	// =========================================================
	// 14. COMPARING ARRAYS
	// =========================================================
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2, 4}
	fmt.Println("a == b:", a == b) // true
	fmt.Println("a == c:", a == c) // false
	// Arrays of the same type can be compared with == and !=
	// when their element type is comparable.
	// Comparison checks corresponding elements.
	// Arrays containing slices, maps, or functions cannot be
	// compared with == because those element types are not comparable.

	// =========================================================
	// 15. ARRAYS AS FUNCTION PARAMETERS
	// =========================================================
	fmt.Println("Sum:", sumThree([3]int{2, 4, 6}))
	// The parameter type [3]int requires exactly 3 integers.
	// An array with a different length cannot be passed to it.

	// =========================================================
	// 16. POINTER TO AN ARRAY
	// =========================================================
	pointerArray := [3]int{7, 8, 9}
	ptr := &pointerArray
	ptr[0] = 70
	fmt.Println("Changed through pointer:", pointerArray)
	// &pointerArray means "address of the whole array".
	// Go allows ptr[0] as shorthand for (*ptr)[0].
	// This lets code modify the original array without copying it.

	// =========================================================
	// 17. MULTIDIMENSIONAL ARRAYS
	// =========================================================
	// An array can contain other arrays.
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix:", matrix)
	fmt.Println("Row 0:", matrix[0])                  // [1 2 3]
	fmt.Println("Element row 1 col 2:", matrix[1][2]) // 6

	// Loop over rows and columns:
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Print(matrix[row][col], " ")
		}
		fmt.Println()
	}

	// =========================================================
	// 18. ARRAY LENGTH IN FUNCTION TYPES
	// =========================================================
	// These are different types:
	var small [2]int
	var large [3]int
	fmt.Printf("small type: %T, large type: %T\n", small, large)
	// Output: [2]int and [3]int.

	// =========================================================
	// 19. ARRAY ZERO VALUES
	// =========================================================
	var zeroInts [3]int
	var zeroStrings [2]string
	var zeroBools [2]bool
	fmt.Println("Zero int array:", zeroInts)       // [0 0 0]
	fmt.Println("Zero string array:", zeroStrings) // [ ]
	fmt.Println("Zero bool array:", zeroBools)     // [false false]
	// Each element gets the zero value of its element type:
	// int -> 0, string -> "", bool -> false, pointers -> nil.

	// =========================================================
	// 20. ARRAY LENGTH CANNOT BE CHANGED
	// =========================================================
	fixed := [3]int{1, 2, 3}
	fmt.Println("Fixed array:", fixed)
	// fixed = append(fixed, 4) // compile error: append expects a slice
	// Go's append works with slices, not arrays.
	// If you need a collection that can grow, use a slice.
	// Slices are covered separately because they are extremely
	// important in everyday Go code.

	// =========================================================
	// 21. ARRAYS VS SLICES VS JAVASCRIPT ARRAYS
	// =========================================================
	// Go array:   [3]int{1, 2, 3}  -> fixed length, value type.
	// Go slice:   []int{1, 2, 3}   -> flexible length, references
	//                                an underlying array.
	// JavaScript: [1, 2, 3]        -> dynamic array-like object.
	//
	// In real Go applications, slices are usually used more often
	// than arrays for lists of items. Arrays are useful when a
	// fixed size is part of the requirement.

	// =========================================================
	// 22. COMMON MISTAKES
	// =========================================================
	// var x [3]int = [4]int{1, 2, 3, 4}
	// ERROR: array lengths differ, so the types differ.
	//
	// x[3] = 100
	// ERROR AT RUNTIME: valid indexes for [3]int are 0, 1, 2.
	//
	// x = append(x, 4)
	// COMPILE ERROR: append expects a slice, not an array.
	//
	// for _, value := range x {
	//     value = 100
	// }
	// This does not change x; value is a copy.
	//
	// Use x[index] to change an element.
}

func printArrayCopy(arr [3]int) {
	// arr is a copy of the caller's array.
	arr[0] = -1
	fmt.Println("Inside printArrayCopy:", arr)
}

func modifyArray(arr *[3]int) {
	// arr points to the caller's original array.
	arr[0] = 500
}

func sumThree(numbers [3]int) int {
	// This function accepts exactly [3]int.
	return numbers[0] + numbers[1] + numbers[2]
}
