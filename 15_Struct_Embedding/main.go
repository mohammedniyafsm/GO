package main

import "fmt"

// ============================================
// 25. WHAT IS STRUCT EMBEDDING?
// ============================================

// Struct embedding allows us to include one struct
// type inside another struct without explicitly
// naming a field for the embedded type.
//
// It is a way to reuse fields and methods from
// another type.
//
// Example:
//
// A Person has:
//   Name
//   Age
//
// An Employee also has:
//   EmployeeID
//
// Instead of declaring Name and Age again inside
// Employee, we can embed Person.
//
// IMPORTANT:
// Go embedding is not the same as inheritance
// in languages like Java or C++.
//
// Go does not have traditional class inheritance.
// Struct embedding promotes fields and methods,
// allowing convenient access through the outer value.

// ============================================
// 26. DECLARING STRUCTS FOR EMBEDDING
// ============================================

// This is the struct we want to reuse.
//
// Person has two fields:
// Name -> string
// Age  -> int

type Person struct {
	Name string
	Age  int
}

// Employee embeds Person.
//
// Notice that Person appears as a type name on its own:
//
// type Employee struct {
//     Person
//     EmployeeID int
// }
//
// We did not write:
//
// Person Person
//
// Because Person is embedded, its fields and methods
// can be accessed conveniently through Employee.

type Employee struct {
	Person
	EmployeeID int
}

// ============================================
// 27. CREATING A STRUCT WITH EMBEDDING
// ============================================

func main() {

	// Create an Employee value.
	//
	// Person is initialized as a nested struct.
	//
	// EmployeeID belongs directly to Employee.

	employee := Employee{
		Person: Person{
			Name: "Niyaf",
			Age:  22,
		},
		EmployeeID: 101,
	}

	fmt.Println(employee)

	// Output:
	// {{Niyaf 22} 101}
	//
	// Explanation:
	//
	// The outer struct is Employee.
	//
	// Employee contains:
	//   Person
	//   EmployeeID
	//
	// Person itself contains:
	//   Name
	//   Age
	//
	// Therefore, the printed output shows the nested
	// Person value inside the Employee value.

	// ============================================
	// 28. ACCESSING EMBEDDED FIELDS
	// ============================================

	// We can access embedded fields directly using
	// the outer struct variable.
	//
	// employee.Name is shorthand for:
	// employee.Person.Name
	//
	// employee.Age is shorthand for:
	// employee.Person.Age
	//
	// This convenient access is called field promotion.

	fmt.Println(employee.Name)
	fmt.Println(employee.Age)
	fmt.Println(employee.EmployeeID)

	// Output:
	// Niyaf
	// 22
	// 101

	// We can also access the nested struct explicitly.

	fmt.Println(employee.Person.Name)
	fmt.Println(employee.Person.Age)

	// Output:
	// Niyaf
	// 22
	//
	// Both forms access the same embedded Person fields.

	// ============================================
	// 29. MODIFYING EMBEDDED FIELDS
	// ============================================

	// Because Name and Age are promoted fields,
	// we can update them directly through Employee.

	employee.Name = "Mohammed"
	employee.Age = 23

	fmt.Println(employee.Name)
	fmt.Println(employee.Person.Name)
	fmt.Println(employee.Age)

	// Output:
	// Mohammed
	// Mohammed
	// 23
	//
	// Explanation:
	// employee.Name updates the Name field inside
	// employee.Person.
	//
	// It does not create a separate Name field
	// in Employee.

	// ============================================
	// 30. EMBEDDING VS A NAMED NESTED FIELD
	// ============================================

	// These two declarations are different.
	//
	// A) Embedding:
	//
	// type Employee struct {
	//     Person
	//     EmployeeID int
	// }
	//
	// Access:
	// employee.Name
	// employee.Person.Name
	//
	// B) Named nested field:
	//
	// type EmployeeWithNamedPerson struct {
	//     Person Person
	//     EmployeeID int
	// }
	//
	// Access:
	// employee2.Person.Name
	//
	// With a named nested field called Person, you
	// cannot access employee2.Name directly because
	// Name is not a direct field of EmployeeWithNamedPerson.

	employee2 := EmployeeWithNamedPerson{
		Person: Person{
			Name: "Rahul",
			Age:  25,
		},
		EmployeeID: 102,
	}

	fmt.Println(employee2.Person.Name)
	fmt.Println(employee2.EmployeeID)

	// Output:
	// Rahul
	// 102
	//
	// The following would NOT compile:
	//
	// fmt.Println(employee2.Name)
	//
	// Why?
	// Person is a named field, not an embedded field.

	// ============================================
	// 31. STRUCT EMBEDDING WITH METHODS
	// ============================================

	// Go methods can have a receiver of a struct type.
	//
	// Let's define a method on Person.
	//
	// Because Employee embeds Person, Person's method
	// is promoted and can be called through Employee.

	employee.PrintDetails()

	// Output:
	// Name: Mohammed Age: 23

	// We can also call the method directly on Person.

	employee.Person.PrintDetails()

	// Output:
	// Name: Mohammed Age: 23
	//
	// Both calls invoke the method on the embedded
	// Person value.

	// ============================================
	// 32. EMBEDDING DOES NOT MEAN INHERITANCE
	// ============================================

	// In Go, Employee does not become a Person.
	//
	// Employee contains an embedded Person value.
	//
	// Embedding promotes selected fields and methods,
	// but it does not create a traditional parent-child
	// inheritance relationship.
	//
	// For example, Employee has its own EmployeeID,
	// while Name and Age belong to its embedded Person.

	fmt.Println(employee.EmployeeID)
	fmt.Println(employee.Name)

	// Output:
	// 101
	// Mohammed

	// ============================================
	// 33. EMBEDDING MULTIPLE STRUCTS
	// ============================================

	// A struct can embed more than one type.
	//
	// Here, Manager embeds both Person and JobInfo.
	//
	// Their promoted fields can be accessed directly
	// as long as there are no conflicting field names.

	manager := Manager{
		Person: Person{
			Name: "Aisha",
			Age:  30,
		},
		JobInfo: JobInfo{
			Department: "Engineering",
			Salary:     80000,
		},
		ManagerLevel: 2,
	}

	fmt.Println(manager.Name)
	fmt.Println(manager.Department)
	fmt.Println(manager.Salary)
	fmt.Println(manager.ManagerLevel)

	// Output:
	// Aisha
	// Engineering
	// 80000
	// 2

	// Explanation:
	// Name and Age come from Person.
	// Department and Salary come from JobInfo.
	// ManagerLevel belongs directly to Manager.

	// ============================================
	// 34. FIELD NAME CONFLICTS IN EMBEDDING
	// ============================================

	// What happens if two embedded structs have
	// a field with the same name?
	//
	// Example:
	//
	// type Contact struct {
	//     Name string
	// }
	//
	// type Company struct {
	//     Name string
	// }
	//
	// type Contractor struct {
	//     Contact
	//     Company
	// }
	//
	// contractor.Name is ambiguous because both
	// embedded fields provide a promoted Name field.
	//
	// Go reports an error if we try to access
	// contractor.Name directly.
	//
	// We can access each field explicitly:
	//
	// contractor.Contact.Name
	// contractor.Company.Name

	contractor := Contractor{
		Contact: Contact{
			Name: "Niyaf",
		},
		Company: Company{
			Name: "Example Technologies",
		},
	}

	fmt.Println(contractor.Contact.Name)
	fmt.Println(contractor.Company.Name)

	// Output:
	// Niyaf
	// Example Technologies

	// ============================================
	// 35. EMBEDDING POINTERS
	// ============================================

	// We can embed a pointer to a struct instead
	// of embedding the struct value itself.
	//
	// Example:
	//
	// type EmployeePointer struct {
	//     *Person
	//     EmployeeID int
	// }
	//
	// The Person field then stores a pointer.
	//
	// This can be useful when we want multiple values
	// to refer to the same Person.
	//
	// IMPORTANT:
	// A nil embedded pointer must be initialized
	// before accessing fields through it.

	pointerEmployee := EmployeePointer{
		Person: &Person{
			Name: "Sara",
			Age:  26,
		},
		EmployeeID: 103,
	}

	fmt.Println(pointerEmployee.Name)
	fmt.Println(pointerEmployee.Age)

	// Output:
	// Sara
	// 26
	//
	// Explanation:
	// Person is a pointer to a Person value.
	//
	// Go automatically follows the pointer when
	// accessing the promoted Name and Age fields.

	// ============================================
	// 36. IMPORTANT RULES OF STRUCT EMBEDDING
	// ============================================

	// 1. Embedding includes a type as an unnamed field.
	//
	// 2. Fields of the embedded type may be promoted.
	//
	// 3. Methods of the embedded type may also be promoted.
	//
	// 4. Embedding is not traditional class inheritance.
	//
	// 5. A named nested field is not the same as embedding.
	//
	// 6. Conflicting promoted fields may become ambiguous.
	//
	// 7. Embedded pointers must not be nil when you
	//    dereference them to access their underlying value.
	//
	// 8. You can embed more than one type.
	//
	// 9. Embedding is often used for composition:
	//    building a larger type from smaller types.
}

// ============================================
// METHOD DEFINED ON PERSON
// ============================================

func (p Person) PrintDetails() {
	fmt.Println("Name:", p.Name, "Age:", p.Age)
}

// ============================================
// OTHER STRUCT TYPES
// ============================================

type EmployeeWithNamedPerson struct {
	Person     Person
	EmployeeID int
}

type JobInfo struct {
	Department string
	Salary     int
}

type Manager struct {
	Person
	JobInfo
	ManagerLevel int
}

type Contact struct {
	Name string
}

type Company struct {
	Name string
}

type Contractor struct {
	Contact
	Company
}

type EmployeePointer struct {
	*Person
	EmployeeID int
}
