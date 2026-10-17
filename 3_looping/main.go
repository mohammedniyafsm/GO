
package main

import "fmt"

func main() {

    // ============================================================
    // 1. BASIC FOR LOOP
    // ============================================================

    // Syntax:
    //
    // for initialization; condition; update {
    //     // Code to execute
    // }
    //
    // initialization -> runs once at the beginning.
    // condition      -> checked before each iteration.
    // update         -> runs after each iteration.
    //
    // The loop stops when the condition becomes false.

    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // Output:
    // 1
    // 2
    // 3
    // 4
    // 5


    // ============================================================
    // 2. HOW A FOR LOOP EXECUTES
    // ============================================================

    // Example:
    //
    // for i := 1; i <= 3; i++ {
    //     fmt.Println(i)
    // }
    //
    // Execution order:
    //
    // 1. i := 1       -> initialize i.
    // 2. i <= 3       -> check condition.
    // 3. Print i      -> execute loop body.
    // 4. i++          -> increment i.
    // 5. Check again  -> repeat while condition is true.
    //
    // When i becomes 4, i <= 3 is false, so the loop stops.


    // ============================================================
    // 3. FOR LOOP WITH A DIFFERENT UPDATE
    // ============================================================

    // You can increment or decrement by different amounts.

    for i := 0; i <= 10; i += 2 {
        fmt.Println(i)
    }

    // Output: 0, 2, 4, 6, 8, 10


    // ============================================================
    // 4. COUNTDOWN LOOP
    // ============================================================

    // Decrement the variable using i--.

    for i := 5; i >= 1; i-- {
        fmt.Println(i)
    }

    // Output: 5, 4, 3, 2, 1


    // ============================================================
    // 5. FOR LOOP WITH ONLY A CONDITION
    // ============================================================

    // Go can use for as a while-style loop.
    //
    // Syntax:
    // for condition {
    //     // Code
    // }
    //
    // There is no initialization or update section.
    // You must update the condition's variables yourself.

    count := 1

    for count <= 5 {
        fmt.Println(count)
        count++
    }

    // This behaves similarly to:
    // while (count <= 5) { ... } in JavaScript.


    // ============================================================
    // 6. INFINITE LOOP
    // ============================================================

    // A for loop without a condition runs indefinitely
    // unless you exit it using break, return, or another control flow.

    // Uncomment to test:
    //
    // for {
    //     fmt.Println("Running...")
    // }

    // Infinite loops are useful for workers, servers, and event
    // processing, but they need a proper exit or cancellation strategy.


    // ============================================================
    // 7. BREAK STATEMENT
    // ============================================================

    // break immediately exits the nearest enclosing loop.

    for i := 1; i <= 10; i++ {
        if i == 5 {
            break
        }

        fmt.Println(i)
    }

    // Output: 1, 2, 3, 4
    // When i becomes 5, the loop terminates.


    // ============================================================
    // 8. CONTINUE STATEMENT
    // ============================================================

    // continue skips the remaining code in the current iteration
    // and proceeds to the next iteration.
    //
    // In a three-part for loop, the post statement still runs
    // after continue.

    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue
        }

        fmt.Println(i)
    }

    // Output: 1, 2, 4, 5


    // ============================================================
    // 9. RANGE LOOP WITH A SLICE
    // ============================================================

    // range iterates over elements of collections.
    //
    // Syntax:
    // for index, value := range collection {
    //     // Code
    // }
    //
    // index -> position of the element.
    // value -> a copy of the element's value.

    numbers := []int{10, 20, 30, 40, 50}

    for index, value := range numbers {
        fmt.Println("Index:", index, "Value:", value)
    }

    // Output:
    // Index: 0 Value: 10
    // Index: 1 Value: 20
    // ...
    //
    // Unlike a traditional indexed loop, range provides the index
    // and value directly.


    // ============================================================
    // 10. RANGE LOOP WITH ONLY THE VALUE
    // ============================================================

    // Use _ to ignore a value you do not need.
    //
    // The blank identifier _ discards a value.

    for _, value := range numbers {
        fmt.Println(value)
    }

    // If you only need the index, use:
    //
    // for index := range numbers {
    //     fmt.Println(index)
    // }


    // ============================================================
    // 11. MODIFYING SLICE ELEMENTS
    // ============================================================

    // The value received by range is a copy of the element.
    // Changing that local variable does not modify the slice.

    scores := []int{10, 20, 30}

    for _, value := range scores {
        value = value * 2
    }

    fmt.Println(scores) // [10 20 30]

    // To modify the original elements, use their indexes.

    for index := range scores {
        scores[index] *= 2
    }

    fmt.Println(scores) // [20 40 60]


    // ============================================================
    // 12. RANGE LOOP WITH AN ARRAY
    // ============================================================

    // range can iterate over arrays too.

    fruits := [3]string{"Apple", "Banana", "Mango"}

    for index, fruit := range fruits {
        fmt.Println(index, fruit)
    }


    // ============================================================
    // 13. RANGE LOOP WITH A MAP
    // ============================================================

    // A map contains key-value pairs.
    //
    // For maps:
    // key   -> map key.
    // value -> corresponding value.
    //
    // IMPORTANT:
    // Go does not guarantee map iteration order.

    userAges := map[string]int{
        "Niyaf": 22,
        "Alex":  25,
        "Sam":   30,
    }

    for name, age := range userAges {
        fmt.Println(name, age)
    }

    // The output order can differ between executions.
    // Do not rely on map iteration order.


    // ============================================================
    // 14. RANGE LOOP WITH A STRING
    // ============================================================

    // range over a string iterates over Unicode code points.
    //
    // The index is the byte offset at which each UTF-8 encoded
    // code point begins, not necessarily its character number.

    text := "Go😊"

    for index, character := range text {
        fmt.Println("Byte index:", index, "Character:", string(character))
    }

    // "G" begins at byte offset 0.
    // "o" begins at byte offset 1.
    // "😊" begins at byte offset 2 because it occupies multiple bytes.


    // ============================================================
    // 15. RANGE LOOP WITH AN INTEGER (GO 1.22+)
    // ============================================================

    // In modern Go versions, range can iterate over an integer.
    // This loop runs from 0 through 4.

    for i := range 5 {
        fmt.Println(i)
    }

    // Output: 0, 1, 2, 3, 4
    //
    // This feature requires Go 1.22 or later.


    // ============================================================
    // 16. NESTED LOOPS
    // ============================================================

    // A loop inside another loop is called a nested loop.
    //
    // The inner loop completes its iterations for every
    // iteration of the outer loop.

    for i := 1; i <= 3; i++ {
        for j := 1; j <= 2; j++ {
            fmt.Println("i:", i, "j:", j)
        }
    }

    // The inner loop runs 2 times for every outer iteration.
    // Total executions of the inner body: 3 * 2 = 6.


    // ============================================================
    // 17. PRINTING A PATTERN USING NESTED LOOPS
    // ============================================================

    // Example output:
    // *
    // **
    // ***
    // ****

    for row := 1; row <= 4; row++ {
        for column := 1; column <= row; column++ {
            fmt.Print("*")
        }

        fmt.Println()
    }

    // The outer loop controls the number of rows.
    // The inner loop controls the number of stars per row.


    // ============================================================
    // 18. BREAK IN NESTED LOOPS
    // ============================================================

    // An ordinary break exits only the nearest enclosing loop.

    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if j == 2 {
                break
            }

            fmt.Println("i:", i, "j:", j)
        }
    }

    // The inner loop stops when j == 2.
    // The outer loop continues with the next i.


    // ============================================================
    // 19. LABELED BREAK
    // ============================================================

    // A label lets break exit a specific enclosing loop.

outer:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if i == 2 && j == 2 {
                break outer
            }

            fmt.Println(i, j)
        }
    }

    // When i == 2 and j == 2, break outer exits both loops.


    // ============================================================
    // 20. LABELED CONTINUE
    // ============================================================

    // A labeled continue proceeds to the next iteration of
    // the specifically labeled enclosing loop.

outerLoop:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if j == 2 {
                continue outerLoop
            }

            fmt.Println(i, j)
        }
    }

    // When j == 2, the current outer-loop iteration ends,
    // and the next outer iteration begins.


    // ============================================================
    // 21. LOOPING THROUGH A SLICE USING AN INDEX
    // ============================================================

    // You can use len() to get a slice's length.
    // This is useful when you need precise control over indexes.

    items := []string{"Keyboard", "Mouse", "Monitor"}

    for i := 0; i < len(items); i++ {
        fmt.Println(items[i])
    }


    // ============================================================
    // 22. SUM OF NUMBERS USING A LOOP
    // ============================================================

    values := []int{10, 20, 30, 40}

    sum := 0

    for _, value := range values {
        sum += value
    }

    fmt.Println("Sum:", sum) // 100


    // ============================================================
    // 23. FINDING A VALUE USING A LOOP
    // ============================================================

    nums := []int{5, 10, 15, 20}
    target := 15
    found := false

    for _, value := range nums {
        if value == target {
            found = true
            break
        }
    }

    fmt.Println("Found:", found) // true


    // ============================================================
    // 24. SKIPPING EVEN NUMBERS
    // ============================================================

    // The modulo operator % gives the remainder.
    // A number is even when number % 2 == 0.

    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }

        fmt.Println(i)
    }

    // Output: 1, 3, 5, 7, 9


    // ============================================================
    // 25. FOR LOOP WITH MULTIPLE CONDITIONS
    // ============================================================

    // You can use logical operators in the condition.
    // && means AND.
    // || means OR.

    for i := 1; i <= 10 && i%2 != 0; i += 2 {
        fmt.Println(i)
    }

    // Stops when either part of the condition becomes false.


    // ============================================================
    // 26. COMMON LOOP MISTAKES
    // ============================================================

    // Mistake 1: Infinite loop due to missing update.
    //
    // i := 1
    // for i <= 5 {
    //     fmt.Println(i)
    //     // Missing i++ means the condition stays true.
    // }

    // Mistake 2: Off-by-one error.
    //
    // for i := 0; i < 5; i++ {
    //     // Runs 5 times: i = 0, 1, 2, 3, 4.
    // }

    // Mistake 3: Ignoring unused variables.
    //
    // for index, value := range numbers {
    //     fmt.Println(value)
    // }
    //
    // This fails to compile if index is unused.
    // Use: for _, value := range numbers { ... }

    // Mistake 4: Expecting map iteration to be ordered.
    //
    // Go maps do not guarantee iteration order.

    // Mistake 5: Modifying a range value instead of the slice element.
    //
    // for _, value := range scores {
    //     value = 0
    // }
    //
    // This does not modify scores. Use scores[index] instead.
}
