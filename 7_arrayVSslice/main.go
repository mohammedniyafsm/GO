package main

import "fmt"

func main() {

	// ============================================
	// 1. ARRAY
	// ============================================

	// An array has a FIXED length.
	// Here, the array can store exactly 3 integers.

	var numbers [3]int

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Println(numbers)

	// Output:
	// [10 20 30]

	// We cannot add a fourth element to this array.
	// numbers[3] = 40 would cause a runtime panic
	// because the valid indexes are 0, 1 and 2.

	// ============================================
	// 2. SLICE
	// ============================================

	// A slice can grow when we append more elements.
	// We do not specify a fixed length in [].

	var values []int

	values = append(values, 10)
	values = append(values, 20)
	values = append(values, 30)

	fmt.Println(values)

	// Output:
	// [10 20 30]

	// Now we can append another element.

	values = append(values, 40)

	fmt.Println(values)

	// Output:
	// [10 20 30 40]

	// ============================================
	// 3. ARRAY LENGTH VS SLICE LENGTH
	// ============================================

	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}

	fmt.Println(len(arr))
	fmt.Println(len(slice))

	// Output:
	// 3
	// 3

	// Both currently have 3 elements.
	// But their lengths behave differently:
	//
	// Array: its length is fixed at 3.
	// Slice: its length can change using append().

	// ============================================
	// 4. ADD MORE ELEMENTS
	// ============================================

	// Array:
	// We cannot increase the length of arr.
	// [3]int will always have exactly 3 elements.

	// Slice:
	slice = append(slice, 4, 5)

	fmt.Println(slice)

	// Output:
	// [1 2 3 4 5]

	// ============================================
	// 5. DEFAULT VALUES
	// ============================================

	// Uninitialized integer arrays and slices have
	// zero-valued elements when those elements exist.

	var arr2 [3]int
	var slice2 []int

	fmt.Println(arr2)
	fmt.Println(slice2)

	// Output:
	// [0 0 0]
	// []

	// arr2 contains 3 zero-valued elements.
	// slice2 is a nil slice with no elements.

	// ============================================
	// 6. CHECK LENGTH AND CAPACITY
	// ============================================

	nums := make([]int, 3, 5)

	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	// Output:
	// 3
	// 5

	// len() gives the number of existing elements.
	// cap() gives the current capacity before the
	// slice needs a larger underlying array.
	//
	// Slices can grow beyond their current capacity.
	// Arrays do not have a separate capacity concept:
	// their length is fixed.

}
