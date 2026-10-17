package main

import "fmt"

func main() {

	// ============================================
	// 1. WHAT IS A SLICE?
	// ============================================

	// A slice is used to store multiple values of the same type.
	// It is similar to an array, but its length can change.
	//
	// Array:  [3]int  -> fixed length of 3
	// Slice:  []int   -> length can grow or shrink
	//
	// Example:
	nums := []int{10, 20, 30, 40}

	fmt.Println(nums)

	// Output:
	// [10 20 30 40]

	// ============================================
	// 2. ACCESS SLICE ELEMENTS
	// ============================================

	// Index starts from 0.
	// Index 0 -> 10
	// Index 1 -> 20
	// Index 2 -> 30
	// Index 3 -> 40

	fmt.Println(nums[0]) // Output: 10
	fmt.Println(nums[1]) // Output: 20
	fmt.Println(nums[2]) // Output: 30

	// Important:
	// Accessing an index that does not exist causes a runtime panic.
	// Example: nums[4] would cause a panic.

	// ============================================
	// 3. MODIFY SLICE ELEMENTS
	// ============================================

	// We can change an element using its index.

	nums[1] = 99

	fmt.Println(nums)

	// Output:
	// [10 99 30 40]
	//
	// Explanation:
	// The value at index 1 changed from 20 to 99.

	// ============================================
	// 4. FIND THE LENGTH USING len()
	// ============================================

	// len() returns the number of elements currently in a slice.

	fmt.Println(len(nums))

	// Output:
	// 4
	//
	// Even after modifying an element, the length stays the same.

	// ============================================
	// 5. CREATE AN EMPTY SLICE
	// ============================================

	// We can create a slice without putting values in it.

	var empty []int

	fmt.Println(empty)
	fmt.Println(len(empty))

	// Output:
	// []
	// 0
	//
	// Explanation:
	// This slice has no elements.
	// A nil slice has length 0.

	// ============================================
	// 6. ADD ELEMENTS USING append()
	// ============================================

	// append() adds elements to the end of a slice.
	// It returns the updated slice, so assign the result back.

	var numbers []int

	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	numbers = append(numbers, 30)

	fmt.Println(numbers)

	// Output:
	// [10 20 30]
	//
	// Step by step:
	// First append  -> [10]
	// Second append -> [10 20]
	// Third append  -> [10 20 30]

	// ============================================
	// 7. APPEND MULTIPLE ELEMENTS
	// ============================================

	// We can add multiple values in one append() call.

	numbers = append(numbers, 40, 50, 60)

	fmt.Println(numbers)

	// Output:
	// [10 20 30 40 50 60]

	// ============================================
	// 8. SLICE LENGTH vs CAPACITY
	// ============================================

	// len()  -> number of elements currently in the slice.
	// cap()  -> number of elements the slice can hold
	//           in its current underlying array before
	//           append needs to allocate a larger array.
	//
	// Capacity is not necessarily the maximum size a slice
	// can ever reach. A slice can grow beyond its capacity.

	values := make([]int, 3, 5)

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(cap(values))

	// Output:
	// [0 0 0]
	// 3
	// 5
	//
	// Explanation:
	// make([]int, 3, 5)
	// 3 -> initial length
	// 5 -> initial capacity
	//
	// All three elements start with the default int value 0.

	// ============================================
	// 9. CREATE A SLICE USING make()
	// ============================================

	// make() creates a slice with a specified length
	// and optionally a specified capacity.

	data := make([]int, 3)

	fmt.Println(data)

	// Output:
	// [0 0 0]
	//
	// Explanation:
	// The slice starts with 3 elements.
	// Since the capacity was not specified, its initial
	// capacity is also 3.

	// ============================================
	// 10. SLICE AN EXISTING SLICE
	// ============================================

	// Syntax:
	// slice[start:end]
	//
	// start -> included
	// end   -> excluded
	//
	// Example:
	// [10 20 30 40 50]
	//  index: 0  1  2  3  4

	items := []int{10, 20, 30, 40, 50}

	fmt.Println(items[1:4])

	// Output:
	// [20 30 40]
	//
	// Explanation:
	// Start at index 1 -> 20
	// Stop before index 4 -> 50 is excluded.

	// ============================================
	// 11. SLICE FROM THE BEGINNING
	// ============================================

	fmt.Println(items[:3])

	// Output:
	// [10 20 30]
	//
	// Explanation:
	// When start is omitted, it defaults to 0.
	// items[:3] means items[0:3].

	// ============================================
	// 12. SLICE UNTIL THE END
	// ============================================

	fmt.Println(items[2:])

	// Output:
	// [30 40 50]
	//
	// Explanation:
	// When end is omitted, it defaults to len(items).
	// items[2:] means items[2:5].

	// ============================================
	// 13. COPY A SLICE USING copy()
	// ============================================

	// copy() copies elements into another slice.
	// The destination slice must already have enough length.

	source := []int{1, 2, 3}
	destination := make([]int, len(source))

	copied := copy(destination, source)

	fmt.Println(destination)
	fmt.Println(copied)

	// Output:
	// [1 2 3]
	// 3
	//
	// Explanation:
	// copy(destination, source)
	// Copies the elements and returns the number copied.
	//
	// Changing destination[0] does not change source[0]
	// because their elements are stored in separate arrays.

	// ============================================
	// 14. LOOP THROUGH A SLICE
	// ============================================

	fruits := []string{"apple", "banana", "mango"}

	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}

	// Output:
	// 0 apple
	// 1 banana
	// 2 mango
	//
	// Explanation:
	// range gives the index and value of each element.
	//
	// If you only need the value:
	//
	// for _, fruit := range fruits {
	//     fmt.Println(fruit)
	// }
	//
	// The underscore _ ignores the index.

	// ============================================
	// 15. REMOVE AN ELEMENT FROM A SLICE
	// ============================================

	// Go does not have a built-in remove() method for slices.
	// We can combine slicing and append() to remove an element.
	//
	// Here, we remove the element at index 2.

	numList := []int{10, 20, 30, 40, 50}

	indexToRemove := 2

	numList = append(
		numList[:indexToRemove],
		numList[indexToRemove+1:]...,
	)

	fmt.Println(numList)

	// Output:
	// [10 20 40 50]
	//
	// Explanation:
	// numList[:2]       -> [10 20]
	// numList[3:]       -> [40 50]
	// The ... passes the second slice's elements
	// as individual arguments to append().

	// ============================================
	// 16. NIL SLICE vs EMPTY SLICE
	// ============================================

	var nilSlice []int
	emptySlice := []int{}

	fmt.Println(nilSlice == nil)
	fmt.Println(emptySlice == nil)
	fmt.Println(len(nilSlice))
	fmt.Println(len(emptySlice))

	// Output:
	// true
	// false
	// 0
	// 0
	//
	// Explanation:
	// nilSlice has no underlying array associated with it.
	// emptySlice is initialized as an empty, non-nil slice.
	//
	// Both have length 0.
	// Both can be used with append().

	// ============================================
	// 17. IMPORTANT: SLICES SHARE UNDERLYING ARRAYS
	// ============================================

	// A slice is a view of an underlying array.
	// Slicing an existing slice can create another slice
	// that shares the same underlying array.

	original := []int{10, 20, 30, 40}

	part := original[1:3]
	part[0] = 99

	fmt.Println(original)
	fmt.Println(part)

	// Output:
	// [10 99 30 40]
	// [99 30]
	//
	// Explanation:
	// part[0] refers to the same underlying array element
	// as original[1].
	// Therefore, changing part also changes original.

	// ============================================
	// 18. QUICK REVISION
	// ============================================

	// []int           -> slice of integers
	// []string        -> slice of strings
	// len(slice)      -> current number of elements
	// cap(slice)      -> current capacity
	// append(slice, x)-> append an element
	// make([]int, 3)  -> create a slice of length 3
	// slice[1]        -> access an element at index 1
	// slice[1:4]      -> index 1 included, index 4 excluded
	// copy(dst, src)  -> copy elements into destination
	// range slice     -> loop through slice elements

}
