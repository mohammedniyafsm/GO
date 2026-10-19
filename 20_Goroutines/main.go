# Goroutines and WaitGroup in Go

```go
package main

import (
	"fmt"
	"sync"
)

// ============================================
// 1. WHAT IS A GOROUTINE?
// ============================================
//
// A goroutine is a lightweight unit of execution
// managed by the Go runtime.
//
// We use the "go" keyword to start a function
// as a goroutine.
//
// Normal function call:
//     sayHello()
//
// Starts a new goroutine:
//     go sayHello()
//
// IMPORTANT:
//
// 1. A goroutine is NOT an operating-system thread.
// 2. Go's runtime schedules goroutines onto OS threads.
// 3. Multiple goroutines can run concurrently.
// 4. Their execution order is NOT guaranteed.
// 5. A goroutine does not automatically make the
//    main function wait for it to finish.
//
// JavaScript comparison:
//
// JavaScript normally executes one callback at a time
// on its main thread.
//
// Node.js can handle concurrent asynchronous operations
// using its event loop and other runtime mechanisms.
//
// Go uses goroutines and its runtime scheduler to
// manage concurrent tasks.
//
// These are different concurrency models.

// ============================================
// 2. NORMAL FUNCTION CALL
// ============================================

func sayHello() {
	fmt.Println("Hello from function")
}

// ============================================
// 3. FUNCTION EXECUTION WITH A GOROUTINE
// ============================================

func sayHi() {
	fmt.Println("Hello from goroutine")
}

// ============================================
// 4. WHAT IS A WAITGROUP?
// ============================================
//
// WaitGroup belongs to Go's "sync" package.
//
// It allows us to wait until a group of goroutines
// has finished its work.
//
// Think of it as a counter that tracks unfinished tasks.
//
// Three important methods:
//
// 1. wg.Add(1)
//
//    Increases the counter by 1.
//    We call it before starting the goroutine.
//
// 2. wg.Done()
//
//    Decreases the counter by 1.
//    We call it when a task finishes.
//
// 3. wg.Wait()
//
//    Blocks the calling goroutine until the counter
//    reaches zero.
//
// Example:
//
// Add(3)  -> Counter = 3
// Done()  -> Counter = 2
// Done()  -> Counter = 1
// Done()  -> Counter = 0
// Wait()  -> Can now return
//
// IMPORTANT:
//
// WaitGroup does not control the execution order.
// It only allows us to wait for task completion.

// ============================================
// 5. UNDERSTANDING POINTERS AND DEFER
// ============================================
//
// *sync.WaitGroup means a pointer to a WaitGroup.
//
// &wg means the address of the WaitGroup variable.
//
// We pass &wg so that all goroutines use the SAME
// WaitGroup instead of receiving separate copies.
//
// defer wg.Done()
//
// "defer" schedules Done() to run when the function
// is about to return.
//
// Example:
//
// func doTask(wg *sync.WaitGroup) {
//     defer wg.Done()
//
//     fmt.Println("Doing task")
// }
//
// Execution:
// 1. defer registers the Done() call.
// 2. The function performs its work.
// 3. The function returns.
// 4. Done() executes and decreases the counter.
//
// Forgetting Done() after Add() can make Wait()
// block forever.

// ============================================
// 6. FUNCTION THAT USES A WAITGROUP
// ============================================

func doTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Task started:", id)

	// Imagine some work happens here.

	fmt.Println("Task completed:", id)
}

// ============================================
// 7. WORKER FUNCTION
// ============================================

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker", id, "is working")
}

// ============================================
// 8. MAIN FUNCTION
// ============================================

func main() {

	// ========================================
	// A. NORMAL FUNCTION CALL
	// ========================================
	//
	// A normal function call runs in the
	// current goroutine.
	//
	// The caller waits for the function to return.

	fmt.Println("===== A. Normal Function =====")

	sayHello()

	fmt.Println("Main continues")

	// OUTPUT:
	//
	// ===== A. Normal Function =====
	// Hello from function
	// Main continues
	//
	// WHY?
	//
	// sayHello() finishes before main continues
	// to the next statement.

	// ========================================
	// B. START A GOROUTINE
	// ========================================
	//
	// The "go" keyword starts a new goroutine.
	//
	// However, main does NOT automatically wait
	// for that goroutine to finish.

	fmt.Println("\n===== B. Goroutine =====")

	var wg1 sync.WaitGroup

	wg1.Add(1)

	go func() {
		defer wg1.Done()
		sayHi()
	}()

	fmt.Println("Main is still running")

	wg1.Wait()

	fmt.Println("Goroutine finished")

	// POSSIBLE OUTPUT:
	//
	// ===== B. Goroutine =====
	// Main is still running
	// Hello from goroutine
	// Goroutine finished
	//
	// The first two messages can appear
	// in either order.
	//
	// WHY?
	//
	// main and the new goroutine can make
	// progress independently.
	//
	// wg1.Wait() makes main wait until the
	// goroutine finishes.
	//
	// Therefore, "Goroutine finished" always
	// appears after "Hello from goroutine".

	// ========================================
	// C. WAITGROUP WITH ONE TASK
	// ========================================
	//
	// We use a WaitGroup to track one task.

	fmt.Println("\n===== C. One Task =====")

	var wg2 sync.WaitGroup

	wg2.Add(1)

	go doTask(1, &wg2)

	wg2.Wait()

	fmt.Println("All tasks completed")

	// OUTPUT:
	//
	// ===== C. One Task =====
	// Task started: 1
	// Task completed: 1
	// All tasks completed
	//
	// WHY?
	//
	// 1. Add(1) registers one unfinished task.
	// 2. go doTask() starts the goroutine.
	// 3. defer wg.Done() runs when doTask returns.
	// 4. Wait() returns when the counter reaches zero.
	// 5. main prints "All tasks completed".

	// ========================================
	// D. WAITGROUP WITH THREE GOROUTINES
	// ========================================
	//
	// We want to start three workers and wait
	// until all three have finished.

	fmt.Println("\n===== D. Multiple Goroutines =====")

	var wg3 sync.WaitGroup

	// Register three unfinished tasks.
	wg3.Add(3)

	// Start three goroutines.
	go worker(1, &wg3)
	go worker(2, &wg3)
	go worker(3, &wg3)

	// Wait until all three tasks finish.
	wg3.Wait()

	fmt.Println("All workers finished")

	// POSSIBLE OUTPUT:
	//
	// ===== D. Multiple Goroutines =====
	// Worker 3 is working
	// Worker 1 is working
	// Worker 2 is working
	// All workers finished
	//
	// The worker order can change on different runs.
	//
	// WHY?
	//
	// Each goroutine is scheduled independently.
	//
	// wg3.Wait() does NOT force workers to run
	// in order 1, 2, 3.
	//
	// It only ensures that all three finish before
	// "All workers finished" is printed.

	// ========================================
	// E. UNDERSTANDING THE WAITGROUP COUNTER
	// ========================================
	//
	// Imagine three workers:
	//
	// wg.Add(3)
	//
	// Counter = 3
	//
	// Worker 1 finishes:
	// wg.Done()
	// Counter = 2
	//
	// Worker 2 finishes:
	// wg.Done()
	// Counter = 1
	//
	// Worker 3 finishes:
	// wg.Done()
	// Counter = 0
	//
	// Now wg.Wait() can return.
	//
	// NOTE:
	// Add(3) registers three tasks at once.
	// Each task must call Done() exactly once.

	// ========================================
	// F. WHY NOT USE time.Sleep()?
	// ========================================
	//
	// Another way to delay main is:
	//
	// time.Sleep(2 * time.Second)
	//
	// But Sleep waits for a fixed duration.
	//
	// It does NOT know whether the goroutines
	// have actually finished.
	//
	// If the tasks finish in 10 milliseconds,
	// we still wait for the remaining time.
	//
	// If the tasks need more than 2 seconds,
	// main might return before they finish.
	//
	// A WaitGroup waits for actual task completion.
}

// ============================================
// 9. IMPORTANT RULES TO REMEMBER
// ============================================
//
// 1. go functionName()
//
//    Starts a new goroutine.
//
// 2. wg.Add(1)
//
//    Registers an unfinished task.
//    Normally call it before starting the goroutine.
//
// 3. defer wg.Done()
//
//    Decreases the counter when the function returns.
//
// 4. wg.Wait()
//
//    Blocks until all registered tasks finish.
//
// 5. Goroutines do not automatically mean
//    separate OS threads.
//
// 6. Goroutine execution order is not guaranteed.
//
// 7. WaitGroup does not return results from goroutines.
//    Use channels or another suitable mechanism
//    when you need to communicate results.
//
// 8. A WaitGroup must not be copied after first use.
//
// 9. Every Add(1) must be balanced by a Done().
//
// 10. main returning ends the Go program, even if
//     other goroutines are still running.
//
// ============================================
// 10. FINAL MENTAL MODEL
// ============================================
//
// Normal function:
//     Call function -> Wait for it to return.
//
// Goroutine:
//     Start concurrent execution -> Continue.
//
// WaitGroup:
//     Register tasks -> Start goroutines
//     -> Wait until all tasks finish.
//
// JavaScript:
//     async functions return Promises.
//     await waits for a Promise to settle.
//     Promise.all() waits for all supplied
//     Promises to fulfill, or rejects when one
//     rejects.
//
// Go:
//     goroutines execute concurrently.
//     WaitGroup waits for completion.
//     Channels can communicate values between
//     goroutines.
//
// CONCURRENCY:
//     Multiple tasks make progress during
//     overlapping periods.
//
// PARALLELISM:
//     Multiple tasks execute at the same time.
//
// Both Node.js and Go support concurrency,
// but their execution and scheduling models differ.
```