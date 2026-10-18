package main

import "fmt"

// ============================================
// 1. WHAT IS A STRUCT?
// ============================================

// A struct is a custom data type that groups related
// fields together.
//
// Each field can have its own data type.
//
// For example, a user can have:
//   Name   -> string
//   Age    -> int
//   Email  -> string
//   Active -> bool
//
// Instead of creating separate variables for every
// user property, we can group them into one struct.
//
// JavaScript comparison:
//
// const user = {
//     name: "Niyaf",
//     age: 22,
//     email: "niyaf@example.com"
// }
//
// Go uses a struct definition to declare the fields
// that values of that type can contain.

// ============================================
// 2. DECLARING A STRUCT
// ============================================

// The type keyword declares a new type.
//
// User is the name of our custom type.
//
// struct defines a collection of fields.
//
// Name string means:
// The field Name stores a string.
//
// Age int means:
// The field Age stores an integer.
//
// Email string means:
// The field Email stores a string.
//
// Active bool means:
// The field Active stores true or false.
//
// IMPORTANT:
// Defining a struct does not automatically create
// a user value. It only defines the type.

type User struct {
	Name   string
	Age    int
	Email  string
	Active bool
}

// ============================================
// 3. CREATING A STRUCT VALUE
// ============================================

func main() {

	// We create a variable named user1.
	//
	// User{...} creates a value of type User.
	//
	// Each field is assigned a value using
	// the field name followed by a colon.

	user1 := User{
		Name:   "Niyaf",
		Age:    22,
		Email:  "niyaf@example.com",
		Active: true,
	}

	fmt.Println(user1)

	// Output:
	// {Niyaf 22 niyaf@example.com true}
	//
	// Explanation:
	// fmt.Println prints the struct's field values
	// in the order the fields are declared.
	//
	// The output does not show the field names.
	// It shows the values contained in the struct.

	// ============================================
	// 4. ACCESSING STRUCT FIELDS
	// ============================================

	// We use the dot operator to access a field.
	//
	// user1.Name means:
	// Access the Name field inside user1.
	//
	// This is similar to accessing an object property
	// in JavaScript using user1.Name.

	fmt.Println(user1.Name)
	fmt.Println(user1.Age)
	fmt.Println(user1.Email)
	fmt.Println(user1.Active)

	// Output:
	// Niyaf
	// 22
	// niyaf@example.com
	// true

	// ============================================
	// 5. MODIFYING STRUCT FIELDS
	// ============================================

	// We can update a field after creating a struct.
	//
	// Use = to assign a new value to an existing field.
	//
	// We use := when declaring a new variable.
	// We do not use := to update an existing field.

	user1.Name = "Mohammed"
	user1.Age = 23
	user1.Active = false

	fmt.Println(user1.Name)
	fmt.Println(user1.Age)
	fmt.Println(user1.Active)

	// Output:
	// Mohammed
	// 23
	// false

	// Explanation:
	// Only the selected fields are changed.
	// The other fields keep their previous values.

	// ============================================
	// 6. CREATING MULTIPLE STRUCT VALUES
	// ============================================

	// One struct type can be used to create
	// many independent values.
	//
	// Each value gets its own set of fields.

	user2 := User{
		Name:   "Rahul",
		Age:    25,
		Email:  "rahul@example.com",
		Active: true,
	}

	user3 := User{
		Name:   "Aisha",
		Age:    24,
		Email:  "aisha@example.com",
		Active: false,
	}

	fmt.Println(user2.Name)
	fmt.Println(user3.Name)

	// Output:
	// Rahul
	// Aisha
	//
	// Explanation:
	// user2 and user3 have the same type: User.
	//
	// But each contains its own field values.
	//
	// Updating user2.Name does not change user3.Name.

	// ============================================
	// 7. ZERO VALUES OF A STRUCT
	// ============================================

	// We can declare a struct without providing
	// any field values.
	//
	// Go automatically initializes every field
	// to that field's zero value.

	var user4 User

	fmt.Println(user4.Name)
	fmt.Println(user4.Age)
	fmt.Println(user4.Email)
	fmt.Println(user4.Active)

	// Output:
	//
	// 0
	//
	// false
	//
	// Explanation:
	// Name   -> ""
	// Age    -> 0
	// Email  -> ""
	// Active -> false
	//
	// The first and third output lines are empty
	// strings, so they look like blank lines.
	//
	// Go does not automatically set these fields
	// to undefined, as you might expect from JavaScript.

	// ============================================
	// 8. STRUCT LITERAL WITHOUT FIELD NAMES
	// ============================================

	// Go also allows values to be provided in
	// declaration order without field names.
	//
	// This is called a positional struct literal.
	//
	// IMPORTANT:
	// Every field must be supplied in the correct order.
	// You cannot omit a field in this form.
	//
	// Named fields are generally easier to read.

	user5 := User{
		"Arun",
		28,
		"arun@example.com",
		true,
	}

	fmt.Println(user5.Name)
	fmt.Println(user5.Age)

	// Output:
	// Arun
	// 28

	// ============================================
	// 9. PARTIAL STRUCT INITIALIZATION
	// ============================================

	// With a named-field struct literal, we can
	// provide only the fields we need.
	//
	// Any omitted fields receive their zero values.

	user6 := User{
		Name: "Sara",
		Age:  21,
	}

	fmt.Println(user6.Name)
	fmt.Println(user6.Age)
	fmt.Println(user6.Email)
	fmt.Println(user6.Active)

	// Output:
	// Sara
	// 21
	//
	// false
	//
	// Explanation:
	// Email was omitted, so it is "".
	// Active was omitted, so it is false.

	// ============================================
	// 10. PASSING A STRUCT TO A FUNCTION
	// ============================================

	// Go passes arguments by value.
	//
	// Passing a struct to a function copies
	// the struct's field values.
	//
	// If the function modifies its copy,
	// the original struct is unchanged.

	product1 := Product{
		Name:  "Laptop",
		Price: 50000,
	}

	changeProductPrice(product1)

	fmt.Println(product1.Price)

	// Output:
	// 50000
	//
	// Explanation:
	// changeProductPrice receives a copy of product1.
	//
	// Changing the copy's Price does not change
	// the original product1.Price.

	// ============================================
	// 11. MODIFYING A STRUCT USING A POINTER
	// ============================================

	// If a function must modify the original struct,
	// we can pass a pointer to that struct.
	//
	// &product1 gets the address of product1.
	//
	// The function receives *Product, meaning
	// a pointer to a Product value.

	changeProductPriceUsingPointer(&product1)

	fmt.Println(product1.Price)

	// Output:
	// 60000
	//
	// Explanation:
	// The function accesses the original struct
	// through its pointer.
	//
	// Therefore, the original Price is updated.

	// ============================================
	// 12. STRUCTS INSIDE STRUCTS
	// ============================================

	// A struct can contain another struct as a field.
	//
	// This is useful when an entity contains
	// another related group of information.
	//
	// Example:
	// A Person has an Address.
	//
	// Person contains:
	//   Name    -> string
	//   Address -> Address struct

	person := Person{
		Name: "Niyaf",
		Address: Address{
			City:    "Sullia",
			Country: "India",
		},
	}

	fmt.Println(person.Name)
	fmt.Println(person.Address.City)
	fmt.Println(person.Address.Country)

	// Output:
	// Niyaf
	// Sullia
	// India
	//
	// Explanation:
	// person.Address accesses the Address field.
	//
	// person.Address.City accesses the City field
	// inside the nested Address struct.

	// ============================================
	// 13. ANONYMOUS STRUCTS
	// ============================================

	// An anonymous struct is a struct value created
	// without declaring a reusable named type.
	//
	// It is useful for temporary data that does not
	// need its own named type.
	//
	// The fields are declared directly in the value.

	employee := struct {
		Name       string
		Department string
	}{
		Name:       "Niyaf",
		Department: "Engineering",
	}

	fmt.Println(employee.Name)
	fmt.Println(employee.Department)

	// Output:
	// Niyaf
	// Engineering
	//
	// Explanation:
	// We define the struct's fields and initialize
	// the value in the same expression.
	//
	// Unlike User, this does not declare a named
	// Employee type that can be reused by name.

	// ============================================
	// 14. COMPARING STRUCTS
	// ============================================

	// A struct can be compared using == when all
	// its fields are comparable.
	//
	// Strings, integers and booleans are comparable.
	//
	// Structs containing slices or maps cannot
	// be compared using ==.

	first := Product{
		Name:  "Mouse",
		Price: 1000,
	}

	second := Product{
		Name:  "Mouse",
		Price: 1000,
	}

	third := Product{
		Name:  "Keyboard",
		Price: 2000,
	}

	fmt.Println(first == second)
	fmt.Println(first == third)

	// Output:
	// true
	// false
	//
	// Explanation:
	// first and second have equal field values.
	//
	// first and third differ in their field values.
	//
	// Go compares each field for equality.

	// ============================================
	// 15. STRUCTS WITH SLICES
	// ============================================

	// A struct can contain a slice.
	//
	// However, structs containing slices cannot
	// be compared directly using ==.
	//
	// Slices are not comparable except against nil.

	student := Student{
		Name:  "Niyaf",
		Marks: []int{80, 90, 95},
	}

	fmt.Println(student.Name)
	fmt.Println(student.Marks)

	// Output:
	// Niyaf
	// [80 90 95]

	// The following comparison is invalid:
	//
	// student1 := Student{Name: "A", Marks: []int{1}}
	// student2 := Student{Name: "A", Marks: []int{1}}
	// fmt.Println(student1 == student2)
	//
	// It causes a compilation error because
	// Student contains a slice field.

	// ============================================
	// 16. EXPORTED AND UNEXPORTED FIELDS
	// ============================================

	// In Go, identifiers starting with an uppercase
	// letter are exported.
	//
	// Name, Age and Email are exported fields.
	//
	// Identifiers starting with a lowercase letter
	// are unexported.
	//
	// An unexported field is accessible only within
	// the package where it is declared.
	//
	// This matters when designing reusable packages.
	//
	// Example:
	//
	// type Account struct {
	//     Owner   string
	//     balance int
	// }
	//
	// Owner can be accessed from another package.
	// balance cannot be accessed directly from another
	// package.
	//
	// IMPORTANT:
	// Uppercase/lowercase controls package visibility,
	// not whether a field is private to a single object.

	account := Account{
		Owner:   "Niyaf",
		balance: 1000,
	}

	fmt.Println(account.Owner)
	fmt.Println(account.balance)

	// Output:
	// Niyaf
	// 1000
	//
	// Both fields can be accessed here because main
	// and Account belong to the same package.

	// ============================================
	// 17. STRUCT METHODS
	// ============================================

	// Go allows methods to be associated with a type.
	//
	// A method is similar to a function, but it has
	// a receiver.
	//
	// For example:
	//
	// func (p Product) Display() { ... }
	//
	// Here:
	// p is the receiver variable.
	// Product is the receiver type.
	// Display is the method name.
	//
	// We can call the method using:
	// product.Display()

	methodProduct := Product{
		Name:  "Monitor",
		Price: 15000,
	}

	methodProduct.Display()

	// Output:
	// Product: Monitor Price: 15000

	// ============================================
	// 18. VALUE RECEIVER VS POINTER RECEIVER
	// ============================================

	// A value receiver receives a copy of the struct.
	//
	// A pointer receiver can modify the original struct.
	//
	// Value receiver:
	// func (p Product) ApplyDiscount() { ... }
	//
	// Pointer receiver:
	// func (p *Product) ApplyDiscount() { ... }
	//
	// Use a pointer receiver when the method needs
	// to modify the original struct or when copying
	// the value would be undesirable.

	discountProduct := Product{
		Name:  "Phone",
		Price: 20000,
	}

	discountProduct.ApplyDiscount()

	fmt.Println(discountProduct.Price)

	// Output:
	// 18000
	//
	// Explanation:
	// ApplyDiscount uses a pointer receiver.
	//
	// The original Price changes from 20000 to 18000.

	// ============================================
	// 19. STRUCTS IN SLICES
	// ============================================

	// We can store multiple struct values in a slice.
	//
	// This is common when processing a list of users,
	// products, orders or database records.

	products := []Product{
		{Name: "Laptop", Price: 50000},
		{Name: "Mouse", Price: 1000},
		{Name: "Keyboard", Price: 2000},
	}

	for _, product := range products {
		fmt.Println(product.Name, product.Price)
	}

	// Output:
	// Laptop 50000
	// Mouse 1000
	// Keyboard 2000
	//
	// Explanation:
	// range iterates over the slice.
	//
	// The first value is the index, which we ignore
	// using _.
	//
	// The second value is a copy of each Product.
	//
	// Modifying product.Price inside this loop would
	// not update the struct stored in the slice.
	//
	// To update an element, use its index:
	//
	// products[0].Price = 45000

	// ============================================
	// 20. POINTERS TO STRUCTS IN A SLICE
	// ============================================

	// We can also store pointers to structs in a slice.
	//
	// This allows multiple parts of a program to
	// refer to the same struct values.

	p1 := &Product{Name: "Tablet", Price: 25000}
	p2 := &Product{Name: "Camera", Price: 40000}

	productPointers := []*Product{p1, p2}

	productPointers[0].Price = 23000

	fmt.Println(p1.Price)

	// Output:
	// 23000
	//
	// Explanation:
	// productPointers[0] points to p1's Product.
	//
	// Updating its Price changes the original Product
	// that p1 points to.

	// ============================================
	// 21. STRUCTS AND FUNCTIONS RETURNING VALUES
	// ============================================

	// A function can return a struct.
	//
	// This is useful when constructing a value from
	// input or returning a result from a calculation.

	newProduct := createProduct("Headphones", 3000)

	fmt.Println(newProduct.Name)
	fmt.Println(newProduct.Price)

	// Output:
	// Headphones
	// 3000

	// ============================================
	// 22. STRUCTS AND ZERO-VALUE INITIALIZATION
	// ============================================

	// Struct fields can themselves be structs.
	//
	// If we omit a nested struct field, that nested
	// struct's fields receive their own zero values.

	var emptyPerson Person

	fmt.Println(emptyPerson.Name)
	fmt.Println(emptyPerson.Address.City)

	// Output:
	//
	//
	// Explanation:
	// Name is an empty string.
	// Address.City is also an empty string.
	//
	// The Address struct itself is not nil.
	// It exists as a zero-valued field inside Person.

	// ============================================
	// 23. STRUCTS AND MEMORY
	// ============================================

	// A struct value contains its fields.
	//
	// Passing a struct by value copies its fields.
	//
	// When a field is itself a slice, map or pointer,
	// copying the struct copies that field's descriptor
	// or pointer, not the entire referenced data.
	//
	// For example, copying a struct containing a slice
	// can produce two struct values whose slices refer
	// to the same underlying array.
	//
	// This is important:
	// A struct is copied by value, but referenced data
	// inside its fields may still be shared.

	original := Container{
		Values: []int{10, 20, 30},
	}

	copied := original

	copied.Values[0] = 999

	fmt.Println(original.Values)
	fmt.Println(copied.Values)

	// Output:
	// [999 20 30]
	// [999 20 30]
	//
	// Explanation:
	// original and copied are different struct values.
	//
	// But their Values slices refer to the same
	// underlying array.
	//
	// Changing an element through copied.Values
	// is visible through original.Values.

	// ============================================
	// 24. WHEN SHOULD WE USE STRUCTS?
	// ============================================

	// Structs are useful for representing related data.
	//
	// Examples:
	//
	// User:
	//   Name, Email, Age
	//
	// Product:
	//   Name, Price, Stock
	//
	// Order:
	//   OrderID, UserID, Total
	//
	// Employee:
	//   Name, Department, Salary
	//
	// Address:
	//   City, State, Country
	//
	// In backend development, structs are commonly
	// used for API request bodies, API responses,
	// application models and configuration values.
	//
	// Structs themselves do not automatically validate
	// data or communicate with a database.
	//
	// Those behaviours must be implemented separately.
}

// ============================================
// STRUCT TYPES USED IN THIS FILE
// ============================================

type Product struct {
	Name  string
	Price int
}

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Address Address
}

type Student struct {
	Name  string
	Marks []int
}

type Account struct {
	Owner   string
	balance int
}

type Container struct {
	Values []int
}

// ============================================
// FUNCTION 1: CHANGE PRODUCT PRICE BY VALUE
// ============================================

func changeProductPrice(p Product) {
	// p is a copy of the original Product.
	//
	// This assignment changes only that copy.

	p.Price = 60000
}

// ============================================
// FUNCTION 2: CHANGE PRODUCT PRICE BY POINTER
// ============================================

func changeProductPriceUsingPointer(p *Product) {
	// p points to the original Product.
	//
	// Updating p.Price changes the original struct.
	//
	// Go automatically dereferences the pointer
	// when accessing a struct field.

	p.Price = 60000
}

// ============================================
// METHOD 1: DISPLAY A PRODUCT
// ============================================

func (p Product) Display() {
	// p is a value receiver.
	//
	// This method receives a copy of the Product.
	//
	// It can read the product's fields and print them.
	// Changes to p itself would not modify the original
	// struct's fields.

	fmt.Println("Product:", p.Name, "Price:", p.Price)
}

// ============================================
// METHOD 2: APPLY A DISCOUNT
// ============================================

func (p *Product) ApplyDiscount() {
	// p is a pointer receiver.
	//
	// This method changes the original Product.
	//
	// Here, we apply a 10% discount.
	//
	// For integer prices:
	// 20000 * 90 / 100 = 18000.

	p.Price = p.Price * 90 / 100
}

// ============================================
// FUNCTION 3: RETURN A NEW PRODUCT
// ============================================

func createProduct(name string, price int) Product {
	// This function constructs a Product and returns it.
	//
	// The caller receives a Product value.
	//
	// Go handles memory management automatically.

	return Product{
		Name:  name,
		Price: price,
	}
}
