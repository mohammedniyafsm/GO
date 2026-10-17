package main

import "fmt"

func main() {

	// ==================================================
	// 1. WHAT IS A MAP?
	// ==================================================

	// A map stores data in KEY-VALUE pairs.
	//
	// Example:
	// Name  -> Niyaf
	// Age   -> 22
	//
	// Syntax:
	// map[keyType]valueType
	//
	// map[string]int means:
	// string = key type
	// int    = value type
	//
	// Unlike arrays and slices, maps use keys to access
	// values instead of numeric indexes.

	ages := map[string]int{
		"Niyaf": 22,
		"Rahul": 25,
		"Aisha": 23,
	}

	fmt.Println(ages)

	// Output (order may vary):
	// map[Aisha:23 Niyaf:22 Rahul:25]

	// ==================================================
	// 2. DECLARE A MAP
	// ==================================================

	// Method 1: Declare a nil map.
	// It is not initialized, so we cannot add entries yet.

	var users map[string]int

	fmt.Println(users)

	// Output:
	// map[]

	// Method 2: Create an initialized empty map using make().
	// We can immediately add entries to this map.

	users = make(map[string]int)
	users["Niyaf"] = 22

	fmt.Println(users["Niyaf"])

	// Output:
	// 22

	// Method 3: Create and initialize a map with values.

	prices := map[string]int{
		"Laptop": 50000,
		"Phone":  20000,
	}

	fmt.Println(prices["Laptop"])

	// Output:
	// 50000

	// Method 4: Short declaration with make().

	scores := make(map[string]int)
	scores["Math"] = 90

	fmt.Println(scores["Math"])

	// Output:
	// 90

	// ==================================================
	// 3. MAP SYNTAX
	// ==================================================

	// map[string]int
	// Key   = string
	// Value = int
	//
	// map[int]string
	// Key   = int
	// Value = string
	//
	// map[string]bool
	// Key   = string
	// Value = bool
	//
	// map[string]float64
	// Key   = string
	// Value = float64
	//
	// Keys must be comparable types.
	// Strings, integers and booleans can be keys.
	// Slices, maps and functions cannot be map keys.

	// ==================================================
	// 4. ACCESS A VALUE
	// ==================================================

	// Use the key inside square brackets.

	person := map[string]int{
		"age":    22,
		"salary": 30000,
	}

	fmt.Println(person["age"])
	fmt.Println(person["salary"])

	// Output:
	// 22
	// 30000

	// ==================================================
	// 5. ADD A NEW KEY-VALUE PAIR
	// ==================================================

	// Assign a value to a new key to add an entry.

	person["height"] = 175

	fmt.Println(person["height"])

	// Output:
	// 175

	// ==================================================
	// 6. UPDATE AN EXISTING VALUE
	// ==================================================

	// If the key already exists, the new value replaces
	// the previous value. Duplicate keys are not created.

	person["age"] = 23

	fmt.Println(person["age"])

	// Output:
	// 23

	// ==================================================
	// 7. DELETE A KEY
	// ==================================================

	// Built-in function:
	// delete(mapName, key)
	//
	// It removes the specified entry.
	// If the key does not exist, nothing happens.

	delete(person, "height")

	fmt.Println(person)

	// Output (order may vary):
	// map[age:23 salary:30000]

	// ==================================================
	// 8. CHECK WHETHER A KEY EXISTS
	// ==================================================

	// IMPORTANT:
	// Reading a missing key returns the zero value
	// of the map's value type.
	//
	// For int, the zero value is 0.
	//
	// Use the comma-ok pattern to check key existence.

	age, exists := person["age"]

	fmt.Println(age)
	fmt.Println(exists)

	// Output:
	// 23
	// true
	//
	// age = the value stored in the map.
	// exists = true if the key exists.

	// Check a key that does not exist.

	value, exists := person["country"]

	fmt.Println(value)
	fmt.Println(exists)

	// Output:
	// 0
	// false
	//
	// The value is 0 because the key is missing.
	// exists is false, so we know the key is absent.

	// ==================================================
	// 9. CHECK A KEY USING if
	// ==================================================

	// The comma-ok pattern is commonly used with if.

	if value, ok := person["salary"]; ok {
		fmt.Println("Salary:", value)
	} else {
		fmt.Println("Salary not found")
	}

	// Output:
	// Salary: 30000

	// ==================================================
	// 10. BUILT-IN FUNCTION: len()
	// ==================================================

	// len(mapName) returns the number of key-value pairs.
	// It does not return the sum of the values.

	subjects := map[string]int{
		"Math":    90,
		"Science": 85,
		"English": 80,
	}

	fmt.Println(len(subjects))

	// Output:
	// 3
	//
	// There are 3 entries in the map.

	// ==================================================
	// 11. BUILT-IN FUNCTION: delete()
	// ==================================================

	// delete() removes an entry using its key.
	// It does not return a value.

	delete(subjects, "Science")

	fmt.Println(len(subjects))

	// Output:
	// 2

	// ==================================================
	// 12. BUILT-IN FUNCTION: clear()
	// ==================================================

	// clear() removes all entries from a map.
	// It was introduced in Go 1.21.
	//
	// It leaves the map initialized and empty.

	clear(subjects)

	fmt.Println(subjects)
	fmt.Println(len(subjects))

	// Output:
	// map[]
	// 0

	// ==================================================
	// 13. LOOP THROUGH A MAP USING range
	// ==================================================

	// range gives each key and its corresponding value.
	//
	// Syntax:
	// for key, value := range mapName {
	//     // code
	// }
	//
	// IMPORTANT:
	// Map iteration order is not guaranteed.

	capitals := map[string]string{
		"India":  "New Delhi",
		"Japan":  "Tokyo",
		"France": "Paris",
	}

	for country, capital := range capitals {
		fmt.Println(country, ":", capital)
	}

	// Possible output:
	// India : New Delhi
	// Japan : Tokyo
	// France : Paris
	//
	// Your output may appear in a different order.

	// ==================================================
	// 14. LOOP THROUGH ONLY KEYS
	// ==================================================

	// Use one variable in range to get only the keys.

	for country := range capitals {
		fmt.Println(country)
	}

	// Possible output:
	// India
	// Japan
	// France

	// ==================================================
	// 15. LOOP THROUGH ONLY VALUES
	// ==================================================

	// Go does not provide a direct values-only iteration
	// syntax for maps.
	//
	// Use the blank identifier _ to ignore each key.

	for _, capital := range capitals {
		fmt.Println(capital)
	}

	// Possible output:
	// New Delhi
	// Tokyo
	// Paris

	// ==================================================
	// 16. NESTED MAPS
	// ==================================================

	// A map can contain another map as its value.
	//
	// Here:
	// Outer key   = username
	// Inner map   = details about that user

	usersInfo := map[string]map[string]string{
		"Niyaf": {
			"city": "Mumbai",
			"role": "Developer",
		},
	}

	fmt.Println(usersInfo["Niyaf"]["city"])
	fmt.Println(usersInfo["Niyaf"]["role"])

	// Output:
	// Mumbai
	// Developer

	// ==================================================
	// 17. MAP WITH BOOL VALUES
	// ==================================================

	// Maps can store true or false values.
	// This is useful for tracking whether something exists
	// or whether a particular condition has been met.

	isActive := map[string]bool{
		"Niyaf": true,
		"Rahul": false,
	}

	fmt.Println(isActive["Niyaf"])
	fmt.Println(isActive["Rahul"])

	// Output:
	// true
	// false

	// ==================================================
	// 18. COUNT OCCURRENCES USING A MAP
	// ==================================================

	// Maps are useful for counting how often each value
	// appears in a collection.

	fruits := []string{
		"apple",
		"banana",
		"apple",
		"mango",
		"banana",
		"apple",
	}

	count := make(map[string]int)

	for _, fruit := range fruits {
		count[fruit]++
	}

	fmt.Println(count["apple"])
	fmt.Println(count["banana"])
	fmt.Println(count["mango"])

	// Output:
	// 3
	// 2
	// 1
	//
	// Explanation:
	// count[fruit]++ increases the count by 1.
	// If the key is absent, its initial int value is 0.

	// ==================================================
	// 19. MAP WITH STRUCT VALUES
	// ==================================================

	// A map can store structs as values.
	// Structs let us group related fields together.

	type User struct {
		Name string
		Age  int
	}

	userMap := map[int]User{
		1: {Name: "Niyaf", Age: 22},
		2: {Name: "Rahul", Age: 25},
	}

	fmt.Println(userMap[1].Name)
	fmt.Println(userMap[1].Age)

	// Output:
	// Niyaf
	// 22

	// ==================================================
	// 20. MAP ASSIGNMENT
	// ==================================================

	// Assigning a map to another variable does not make
	// an independent copy of all its entries.
	// Both variables refer to the same underlying map.

	first := map[string]int{
		"apple": 10,
	}

	second := first
	second["apple"] = 50

	fmt.Println(first["apple"])
	fmt.Println(second["apple"])

	// Output:
	// 50
	// 50
	//
	// Updating an entry through second is visible through
	// first because both refer to the same map.

	// ==================================================
	// 21. NIL MAP VS EMPTY MAP
	// ==================================================

	// A nil map has not been initialized.
	// An empty map has been initialized but has no entries.

	var nilMap map[string]int
	emptyMap := make(map[string]int)

	fmt.Println(nilMap == nil)
	fmt.Println(emptyMap == nil)

	// Output:
	// true
	// false
	//
	// Both maps have length 0.
	// However, only the initialized map can accept writes.

	fmt.Println(len(nilMap))
	fmt.Println(len(emptyMap))

	// Output:
	// 0
	// 0

	// ==================================================
	// 22. WHY WRITING TO A NIL MAP CAUSES A PANIC
	// ==================================================

	// This would cause a runtime panic:
	//
	// nilMap["age"] = 22
	//
	// Why?
	// Because nilMap has not been initialized.
	//
	// Correct approach:
	//
	// nilMap = make(map[string]int)
	// nilMap["age"] = 22

	// ==================================================
	// 23. MAP KEYS MUST BE UNIQUE
	// ==================================================

	// Each key can appear only once in a map.
	// Assigning to an existing key replaces its value.

	stock := map[string]int{
		"laptop": 5,
	}

	stock["laptop"] = 10

	fmt.Println(stock["laptop"])

	// Output:
	// 10
	//
	// The value is 10, not two separate entries.

	// ==================================================
	// 24. MAPS ARE NOT ORDERED
	// ==================================================

	// A Go map does not guarantee iteration order.
	// Do not depend on entries appearing in insertion order.
	//
	// If you need sorted output:
	// 1. Collect the keys into a slice.
	// 2. Sort the slice using the sort package.
	// 3. Loop through the sorted keys to access values.

	// ==================================================
	// 25. COMPARING MAPS
	// ==================================================

	// You cannot compare two maps using ==.
	//
	// This is invalid:
	//
	// map1 == map2
	//
	// The only valid map comparison with == is against nil.
	//
	// For example:
	//
	// map1 == nil
	//
	// To compare two maps' contents, use an explicit
	// comparison or the maps package from the standard
	// library (available since Go 1.21).

	// ==================================================
	// 26. MAPS AND CONCURRENCY
	// ==================================================

	// Ordinary Go maps are not safe for concurrent
	// reads and writes without synchronization.
	//
	// If multiple goroutines access a map and at least
	// one writes to it, use a synchronization strategy
	// such as sync.Mutex or consider sync.Map.
	//
	// We will learn goroutines and synchronization later.

	// ==================================================
	// 27. IMPORTANT BUILT-IN OPERATIONS REVISION
	// ==================================================

	// make(map[string]int)
	// Creates an initialized map.
	//
	// len(myMap)
	// Returns the number of entries.
	//
	// delete(myMap, key)
	// Removes the entry with the specified key.
	//
	// clear(myMap)
	// Removes all entries (Go 1.21+).
	//
	// myMap[key]
	// Reads the value associated with a key.
	//
	// myMap[key] = value
	// Adds a new entry or updates an existing entry.
	//
	// value, ok := myMap[key]
	// Checks whether a key exists.
	//
	// range myMap
	// Iterates over map entries.

	// ==================================================
	// 28. QUICK REVISION
	// ==================================================

	// Map declaration:
	// var m map[string]int
	//
	// Initialize an empty map:
	// m = make(map[string]int)
	//
	// Initialize with values:
	// m = map[string]int{"age": 22}
	//
	// Add or update:
	// m["age"] = 23
	//
	// Read:
	// age := m["age"]
	//
	// Check existence:
	// age, ok := m["age"]
	//
	// Delete:
	// delete(m, "age")
	//
	// Length:
	// len(m)
	//
	// Clear:
	// clear(m)
	//
	// Loop:
	// for key, value := range m {
	//     fmt.Println(key, value)
	// }

}
