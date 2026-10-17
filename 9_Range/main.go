package main

import "fmt"

func main() {

	// ============================================
	// 1. WHAT IS range?
	// ============================================

	// range lets us go through elements one by one.
	//
	// Syntax:
	//
	// for index, value := range collection {
	//     // code
	// }
	//
	// index -> position of the element (for arrays/slices)
	// value -> the element stored at that position
	//
	// For maps, range gives keys and values instead.

	// ============================================
	// 2. range WITH A SLICE
	// ============================================

	// A slice stores multiple values.
	// range gives us each index and its corresponding value.

	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// Output:
	// 0 10
	// 1 20
	// 2 30
	// 3 40
	//
	// Explanation:
	// First iteration: index = 0, value = 10
	// Second iteration: index = 1, value = 20
	// Third iteration: index = 2, value = 30
	// Fourth iteration: index = 3, value = 40

	// ============================================
	// 3. range WITH ONLY VALUES
	// ============================================

	// If we do not need the index, use _ to ignore it.
	// The underscore is called the blank identifier.

	fruits := []string{"apple", "banana", "mango"}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Output:
	// apple
	// banana
	// mango
	//
	// Explanation:
	// _ ignores the index.
	// fruit receives each element.

	// ============================================
	// 4. range WITH ONLY INDEXES
	// ============================================

	// If we declare only one variable in a range loop
	// over a slice or array, it receives the index.

	for index := range numbers {
		fmt.Println(index)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	//
	// Explanation:
	// Only the index is returned because we declared
	// only one variable.

	// ============================================
	// 5. range WITH AN ARRAY
	// ============================================

	// range works with arrays too.
	// An array has a fixed length.

	arr := [3]int{100, 200, 300}

	for index, value := range arr {
		fmt.Println(index, value)
	}

	// Output:
	// 0 100
	// 1 200
	// 2 300

	// ============================================
	// 6. range WITH A MAP
	// ============================================

	// A map stores key-value pairs.
	// In a map loop:
	// First variable  -> key
	// Second variable -> value

	ages := map[string]int{
		"Niyaf": 22,
		"Rahul": 25,
		"Aisha": 23,
	}

	for name, age := range ages {
		fmt.Println(name, age)
	}

	// Possible output:
	// Niyaf 22
	// Rahul 25
	// Aisha 23
	//
	// IMPORTANT:
	// Map iteration order is not guaranteed.
	// The output order may change when you run the code.

	// ============================================
	// 7. range WITH ONLY MAP KEYS
	// ============================================

	// Declare only one variable to receive the map key.

	for name := range ages {
		fmt.Println(name)
	}

	// Possible output:
	// Niyaf
	// Rahul
	// Aisha

	// ============================================
	// 8. range WITH A STRING
	// ============================================

	// When ranging over a string:
	// First variable  -> byte offset
	// Second variable -> Unicode code point (rune)
	//
	// A rune represents a Unicode code point.
	// For ASCII characters, the byte offset increases by 1.
	// For multi-byte characters, the offset can increase
	// by more than 1.

	word := "Go!"

	for index, character := range word {
		fmt.Println(index, string(character))
	}

	// Output:
	// 0 G
	// 1 o
	// 2 !
	//
	// string(character) converts the rune to a string
	// so that we can print the character itself.

	// ============================================
	// 9. range WITH A STRING CONTAINING UNICODE
	// ============================================

	// Some characters take more than one byte in UTF-8.
	// range over a string moves from one Unicode code point
	// to the next, rather than processing every byte.

	text := "Go😊"

	for index, character := range text {
		fmt.Println(index, string(character))
	}

	// Output:
	// 0 G
	// 1 o
	// 2 😊
	//
	// The emoji occupies multiple bytes, so the next
	// character's byte offset would jump by more than 1.

	// ============================================
	// 10. range WITH A COUNTER
	// ============================================

	// Sometimes we only want to repeat something a certain
	// number of times.
	// We can range over an integer in Go 1.22 and later.

	for i := range 5 {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	//
	// range 5 produces the values 0 through 4.
	// This requires Go 1.22 or later.
	//
	// It is similar to:
	// for i := 0; i < 5; i++ {
	//     fmt.Println(i)
	// }

	// ============================================
	// 11. MODIFY SLICE ELEMENTS USING THE INDEX
	// ============================================

	// IMPORTANT:
	// The value variable in range is a copy of each element.
	// Changing that variable alone does not change the slice.
	//
	// Use the index to update the actual element.

	nums := []int{1, 2, 3}

	for index := range nums {
		nums[index] = nums[index] * 2
	}

	fmt.Println(nums)

	// Output:
	// [2 4 6]
	//
	// Explanation:
	// nums[0] becomes 2
	// nums[1] becomes 4
	// nums[2] becomes 6

	// ============================================
	// 12. WHY CHANGING THE VALUE VARIABLE DOES NOT WORK
	// ============================================

	nums2 := []int{1, 2, 3}

	for _, value := range nums2 {
		value = value * 2
	}

	fmt.Println(nums2)

	// Output:
	// [1 2 3]
	//
	// Explanation:
	// value is a copy of each element.
	// Changing value does not modify nums2.
	//
	// Use nums2[index] if you want to update the slice.

	// ============================================
	// 13. range WITH A CONDITION
	// ============================================

	// We can use if inside range to filter values.

	values := []int{10, 15, 20, 25, 30}

	for _, value := range values {
		if value > 20 {
			fmt.Println(value)
		}
	}

	// Output:
	// 25
	// 30
	//
	// Explanation:
	// Only values greater than 20 are printed.

	// ============================================
	// 14. range WITH break
	// ============================================

	// break stops the loop immediately.

	for _, value := range numbers {
		if value == 30 {
			break
		}

		fmt.Println(value)
	}

	// Output:
	// 10
	// 20
	//
	// Explanation:
	// When value becomes 30, break stops the loop.

	// ============================================
	// 15. range WITH continue
	// ============================================

	// continue skips the current iteration
	// and moves to the next iteration.

	for _, value := range values {
		if value == 20 {
			continue
		}

		fmt.Println(value)
	}

	// Output:
	// 10
	// 15
	// 25
	// 30
	//
	// Explanation:
	// The value 20 is skipped.

	// ============================================
	// 16. range WITH A SLICE OF STRUCTS
	// ============================================

	// range can iterate through a slice containing structs.
	// Structs group related fields together.

	type User struct {
		Name string
		Age  int
	}

	users := []User{
		{Name: "Niyaf", Age: 22},
		{Name: "Rahul", Age: 25},
	}

	for _, user := range users {
		fmt.Println(user.Name, user.Age)
	}

	// Output:
	// Niyaf 22
	// Rahul 25

	// ============================================
	// 17. range AND THE NORMAL for LOOP
	// ============================================

	// Normal for loop:
	// You control the starting value, condition,
	// and increment.

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 1
	// 2
	//
	// range loop:
	// Convenient when you want to iterate through
	// the elements of a collection.

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// Output:
	// 0 10
	// 1 20
	// 2 30
	//
	// Use a normal for loop when you need custom
	// counting or more control over the loop.

	// ============================================
	// 18. QUICK REVISION
	// ============================================

	// for index, value := range slice
	// Gets both the index and the value.
	//
	// for _, value := range slice
	// Gets only the value.
	//
	// for index := range slice
	// Gets only the index.
	//
	// for key, value := range map
	// Gets the map key and value.
	//
	// for index, character := range string
	// Gets the byte offset and Unicode code point.
	//
	// for i := range 5
	// Produces 0, 1, 2, 3, 4 (Go 1.22+).
	//
	// break    -> stop the loop.
	// continue -> skip the current iteration.
	//
	// Important:
	// The second variable is generally a copy of the
	// collection element. Use the index to modify
	// actual slice or array elements.

}
