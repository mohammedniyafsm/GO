package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// 1. WHAT IS A CHANNEL?
// ============================================
//
// A channel is a Go feature that allows goroutines
// to communicate by sending and receiving values.
//
// Remember:
//
// Goroutines = Execute work concurrently.
// Channels   = Communicate data between goroutines.
// WaitGroup  = Wait until goroutines finish.
//
// Example:
//
// One goroutine calculates a price.
// Another goroutine receives that price.
//
// Channels allow them to exchange the result.
//
// JavaScript comparison:
//
// In JavaScript, we might return a value from an
// async function and receive it using await.
//
// In Go, a goroutine cannot directly return a value
// to the caller in the same way that a normal
// synchronous function does.
//
// A channel is one way to communicate its result.
//
// A channel is typed:
//
// chan int    -> carries integers
// chan string -> carries strings
// chan bool   -> carries booleans
//
// ============================================
// 2. CREATING A CHANNEL
// ============================================
//
// Syntax:
//
// ch := make(chan int)
//
// make(chan int) creates a channel that carries
// integer values.
//
// Important:
//
// Declaring a channel variable and creating a
// usable channel are different things.
//
// var ch chan int
//
// Here, ch is nil. It is not yet connected to
// an actual channel.
//
// ch = make(chan int)
//
// Now ch refers to a created channel.
//
// A nil channel blocks forever when sending or
// receiving. Do not send to or receive from a nil
// channel unless you specifically intend that behavior.
//
// ============================================
// 3. SENDING AND RECEIVING VALUES
// ============================================
//
// The <- operator is used for channel communication.
//
// Sending:
//
// ch <- 100
//
// This sends 100 into the channel.
//
// Receiving:
//
// value := <-ch
//
// This receives a value from the channel and stores
// it inside value.
//
// Remember:
//
// ch <- value  = Send a value.
// <-ch         = Receive a value.
//
// ============================================
// 4. YOUR FIRST CHANNEL EXAMPLE
// ============================================

func basicChannelExample() {
	fmt.Println("\n--- 4. Basic Channel ---")

	// Create a channel for strings.
	ch := make(chan string)

	go func() {
		// Send a string into the channel.
		ch <- "Hello from the goroutine!"
	}()

	// Receive the string.
	//
	// Because this is an unbuffered channel,
	// this receive waits until a sender is ready.
	message := <-ch

	fmt.Println(message)
}

// Expected output:
//
// --- 4. Basic Channel ---
// Hello from the goroutine!
//
// ============================================
// 5. HOW AN UNBUFFERED CHANNEL WORKS
// ============================================
//
// By default:
//
// ch := make(chan int)
//
// creates an UNBUFFERED channel.
//
// An unbuffered channel does not store values for
// later collection. The sender and receiver must
// synchronize with each other.
//
// Sending:
//
// ch <- 10
//
// The send cannot complete until a receiver is ready.
//
// Receiving:
//
// value := <-ch
//
// The receive waits until a sender provides a value.
//
// Example timeline:
//
// Sender: ch <- 10
//             |
//             | Waits for receiver
//             |
// Receiver: value := <-ch
//             |
//             | Value 10 is transferred
//             |
// Both operations can continue.
//
// Important:
//
// The channel does not automatically start a goroutine.
// The "go" keyword starts a goroutine.
//
// ============================================
// 6. WHY DO WE NEED CHANNELS?
// ============================================
//
// Imagine one goroutine performs a task and produces
// a result. Another goroutine needs that result.
//
// Without a communication mechanism, we need another
// way to share the result and coordinate execution.
//
// A channel can do both:
//
// 1. Transfer the result.
// 2. Synchronize the sender and receiver.
//
// Channels are useful for:
//
// - Passing results from goroutines.
// - Sending work to workers.
// - Receiving completed task results.
// - Communicating between concurrent tasks.
// - Coordinating the flow of a program.
//
// ============================================
// 7. CHANNELS VS WAITGROUP
// ============================================
//
// WaitGroup:
//
// wg.Add(1)
// go doTask(&wg)
// wg.Wait()
//
// A WaitGroup helps main wait until tasks finish.
//
// Channel:
//
// ch := make(chan int)
// go func() {
//     ch <- 100
// }()
// result := <-ch
//
// A channel can communicate the actual result.
//
// Use a WaitGroup when you need to wait for a group
// of goroutines to finish.
//
// Use a channel when goroutines need to send or
// receive values.
//
// They can also be used together.
//
// ============================================
// 8. RECEIVING A VALUE BLOCKS UNTIL IT IS AVAILABLE
// ============================================

func receiveWaitExample() {
	fmt.Println("\n--- 8. Receiving Waits ---")

	ch := make(chan int)

	go func() {
		time.Sleep(500 * time.Millisecond)

		// Send the value after a short delay.
		ch <- 42
	}()

	fmt.Println("Waiting for the result...")

	// This line waits until 42 is sent.
	result := <-ch

	fmt.Println("Received:", result)
}

// Expected output:
//
// --- 8. Receiving Waits ---
// Waiting for the result...
// Received: 42
//
// Note:
//
// time.Sleep is used here only to demonstrate waiting.
// It is not the correct way to synchronize goroutines
// in normal production code.
//
// ============================================
// 9. SENDING MULTIPLE VALUES
// ============================================
//
// A channel can carry multiple values over time.
//
// The receiver must receive each value that the sender
// sends through an unbuffered channel.

func multipleValuesExample() {
	fmt.Println("\n--- 9. Multiple Values ---")

	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
	}()

	first := <-ch
	second := <-ch
	third := <-ch

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}

// Expected output:
//
// --- 9. Multiple Values ---
// 10
// 20
// 30
//
// The sends and receives synchronize one by one.
//
// ============================================
// 10. BUFFERED CHANNELS
// ============================================
//
// An unbuffered channel requires a sender and receiver
// to synchronize for each transfer.
//
// A buffered channel can hold a limited number of
// values without a receiver immediately collecting them.
//
// Syntax:
//
// ch := make(chan int, 3)
//
// The number 3 is the buffer capacity.
//
// This channel can hold up to 3 integer values
// before a sender needs a receiver to make space.
//
// ============================================

func bufferedChannelExample() {
	fmt.Println("\n--- 10. Buffered Channel ---")

	// Create a channel with capacity 3.
	ch := make(chan int, 3)

	// These sends complete without a receiver,
	// because the buffer has enough free space.
	ch <- 100
	ch <- 200
	ch <- 300

	fmt.Println("Channel length:", len(ch))
	fmt.Println("Channel capacity:", cap(ch))

	// Receive the values.
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Expected output:
//
// --- 10. Buffered Channel ---
// Channel length: 3
// Channel capacity: 3
// 100
// 200
// 300
//
// len(ch) = Number of values currently waiting in the buffer.
// cap(ch) = Total buffer capacity.
//
// Important:
//
// If the buffer becomes full, another send blocks
// until space becomes available.
//
// If the buffer is empty, a receive blocks until
// a value becomes available.
//
// A buffer is not unlimited storage.
//
// ============================================
// 11. UNBUFFERED VS BUFFERED CHANNEL
// ============================================
//
// UNBUFFERED:
//
// ch := make(chan int)
//
// The sender and receiver synchronize directly.
//
// BUFFERED:
//
// ch := make(chan int, 5)
//
// The channel can hold up to 5 values without an
// immediate receiver.
//
// Use unbuffered channels when direct synchronization
// between sender and receiver is useful.
//
// Use buffered channels when limited queuing is useful.
//
// Buffering does not automatically make a program
// faster or guarantee that every goroutine finishes.
//
// ============================================
// 12. CLOSING A CHANNEL
// ============================================
//
// A channel can be closed to signal that no more
// values will be sent through it.
//
// Syntax:
//
// close(ch)
//
// Closing a channel does NOT mean:
//
// - Stop the receiving goroutine immediately.
// - Delete values already in the buffer.
// - Automatically stop every goroutine.
//
// A receiver can still receive values that were sent
// before the channel was closed.
//
// After buffered values are consumed, receiving from
// a closed channel returns the zero value immediately.
//
// Usually, the sender is responsible for closing
// the channel when no more values will be sent.
//
// Never close a channel just because one receiver
// has finished reading.
//
// ============================================

func closeChannelExample() {
	fmt.Println("\n--- 12. Closing a Channel ---")

	ch := make(chan int, 3)

	ch <- 10
	ch <- 20

	// Tell receivers that no more values will be sent.
	close(ch)

	// Values sent before closing are still available.
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// The channel is closed and empty.
	// Receiving returns the zero value for int: 0.
	value, ok := <-ch

	fmt.Println("Value:", value)
	fmt.Println("Received value:", ok)
}

// Expected output:
//
// --- 12. Closing a Channel ---
// 10
// 20
// Value: 0
// Received value: false
//
// ============================================
// 13. THE TWO-VALUE RECEIVE
// ============================================
//
// Syntax:
//
// value, ok := <-ch
//
// value = Received value.
// ok    = Whether the value was received because
//         the channel is still open.
//
// If ok is true:
// A value was received normally.
//
// If ok is false:
// The channel is closed and empty.
//
// This is useful when you need to distinguish an
// actual zero value from a closed channel.
//
// Example:
//
// An int channel may legitimately send 0.
// Therefore, checking only the value is insufficient.
//
// ============================================
// 14. USING RANGE TO RECEIVE CHANNEL VALUES
// ============================================
//
// You can use range to receive values from a channel.
//
// Syntax:
//
// for value := range ch {
//     fmt.Println(value)
// }
//
// The loop receives values until the channel is closed
// and all buffered values have been consumed.
//
// Important:
//
// A range loop on a channel does not finish just
// because the sender has temporarily stopped sending.
//
// The channel must be closed for the loop to finish,
// unless the loop exits some other way.
//
// ============================================

func rangeChannelExample() {
	fmt.Println("\n--- 14. Range Over a Channel ---")

	ch := make(chan string)

	go func() {
		ch <- "Apple"
		ch <- "Banana"
		ch <- "Mango"

		// Tell the receiver there are no more values.
		close(ch)
	}()

	// Receive each value until the channel is closed
	// and all sent values have been consumed.
	for fruit := range ch {
		fmt.Println(fruit)
	}
}

// Expected output:
//
// --- 14. Range Over a Channel ---
// Apple
// Banana
// Mango
//
// ============================================
// 15. DIRECTIONAL CHANNELS
// ============================================
//
// Go allows function parameters to restrict the
// direction of channel communication.
//
// chan<- int
//
// Send-only channel.
//
// <-chan int
//
// Receive-only channel.
//
// chan int
//
// Channel that can send and receive values.
//
// Why use directional channels?
//
// They make the intended responsibility of a
// function clear and help the compiler catch mistakes.
//
// ============================================

// This function can only SEND integers.
func sendNumber(ch chan<- int) {
	ch <- 500
}

// This function can only RECEIVE integers.
func receiveNumber(ch <-chan int) {
	value := <-ch
	fmt.Println("Received number:", value)
}

func directionalChannelExample() {
	fmt.Println("\n--- 15. Directional Channels ---")

	ch := make(chan int)

	go sendNumber(ch)

	receiveNumber(ch)
}

// Expected output:
//
// --- 15. Directional Channels ---
// Received number: 500
//
// Notice:
//
// chan<- int means send-only.
// <-chan int means receive-only.
//
// ============================================
// 16. CHANNELS WITH WAITGROUP
// ============================================
//
// A WaitGroup and a channel solve different problems,
// and you can combine them.
//
// In this example:
//
// - Each worker performs a task.
// - Each worker sends its result to a channel.
// - A WaitGroup tracks when all workers finish.
// - Main closes the results channel after all workers
//   finish sending their results.
// - Main receives results using range.
//
// Important:
//
// The WaitGroup ensures all senders finish before
// the channel is closed.
//
// Closing too early could cause a send on a closed
// channel, which panics.
//
// ============================================

func workerResult(id int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate producing a result.
	result := fmt.Sprintf("Worker %d completed", id)

	// Send the result to the channel.
	results <- result
}

func channelWithWaitGroupExample() {
	fmt.Println("\n--- 16. Channels + WaitGroup ---")

	results := make(chan string, 3)

	var wg sync.WaitGroup

	// We are starting three workers.
	wg.Add(3)

	go workerResult(1, results, &wg)
	go workerResult(2, results, &wg)
	go workerResult(3, results, &wg)

	// This goroutine waits for all workers to finish.
	go func() {
		wg.Wait()

		// No workers remain that can send results.
		close(results)
	}()

	// Receive all worker results.
	//
	// Their order is not guaranteed because workers
	// execute concurrently.
	for result := range results {
		fmt.Println(result)
	}
}

// Expected output:
//
// --- 16. Channels + WaitGroup ---
// Worker 1 completed
// Worker 2 completed
// Worker 3 completed
//
// The order of the three worker messages may vary.
//
// ============================================
// 17. WHAT IS A DEADLOCK?
// ============================================
//
// A deadlock happens when goroutines cannot make
// progress because they are waiting on each other
// or on operations that can never complete.
//
// Example of a deadlock:
//
// ch := make(chan int)
// ch <- 10
//
// If this send happens in main and no other goroutine
// receives from ch, main blocks forever.
//
// An unbuffered send needs a receiver.
//
// Another common deadlock:
//
// ch := make(chan int)
// value := <-ch
//
// If no goroutine ever sends to ch, this receive
// blocks forever.
//
// The Go runtime detects many deadlocks when all
// goroutines in the program are stuck and reports
// a fatal error.
//
// Do not run a deliberate deadlock in this main()
// function. It would prevent the remaining examples
// from running.
//
// ============================================
// 18. IMPORTANT CHANNEL RULES
// ============================================
//
// 1. Create a channel using make:
//
//    ch := make(chan int)
//
// 2. Send a value:
//
//    ch <- 10
//
// 3. Receive a value:
//
//    value := <-ch
//
// 4. An unbuffered send and receive synchronize.
//
// 5. A buffered channel can hold a limited number
//    of values without an immediate receiver.
//
// 6. Receive from a closed channel:
//    Buffered values are still available.
//    Once empty, the receive returns the zero value.
//
// 7. Sending to a closed channel causes a panic.
//
// 8. Closing an already closed channel causes a panic.
//
// 9. A nil channel blocks on sends and receives.
//
// 10. range over a channel ends when it is closed
//     and all buffered values have been consumed.
//
// 11. Usually, the sender closes a channel when
//     there will be no more sends.
//
// 12. Use WaitGroup to wait for completion and
//     channels to communicate values. They can
//     work together.
//
// ============================================
// 19. WHEN SHOULD YOU USE CHANNELS?
// ============================================
//
// Example 1: Send the result of a calculation.
//
// A goroutine calculates a total and sends it to main.
//
// Example 2: Worker pools.
//
// Multiple worker goroutines receive jobs from a
// jobs channel and send results to another channel.
//
// Example 3: Concurrent API requests.
//
// Multiple goroutines perform requests and send
// their results back through a channel.
//
// Example 4: Pipeline processing.
//
// One goroutine reads data, another transforms it,
// and another processes the result.
//
// Example 5: Coordination.
//
// A channel can signal that an event has happened.
//
// Important:
//
// Not every shared value needs a channel.
// A mutex can be a better choice when goroutines
// need to protect shared mutable state.
//
// Choose the simplest synchronization mechanism
// that correctly expresses the problem.
//
// ============================================
// 20. FINAL MENTAL MODEL
// ============================================
//
// GOROUTINE:
//
// Runs a function concurrently.
//
// CHANNEL:
//
// Transfers values between goroutines and can
// synchronize their execution.
//
// WAITGROUP:
//
// Waits until a group of goroutines has finished.
//
// BUFFERED CHANNEL:
//
// Holds a limited number of values temporarily.
//
// UNBUFFERED CHANNEL:
//
// Synchronizes the sender and receiver directly.
//
// CLOSE:
//
// Signals that no more values will be sent.
//
// RANGE:
//
// Receives values until a channel is closed and
// all buffered values have been consumed.
//
// ============================================
// MAIN FUNCTION
// ============================================

func main() {
	basicChannelExample()
	receiveWaitExample()
	multipleValuesExample()
	bufferedChannelExample()
	closeChannelExample()
	rangeChannelExample()
	directionalChannelExample()
	channelWithWaitGroupExample()
}

// Note:
//
// Output sections appear in the order main() calls
// the examples.
//
// Within the WaitGroup example, worker result order
// can vary.
//
// The time.Sleep in receiveWaitExample is only used
// to demonstrate a delayed result.
