// Example: Simple Data Race Condition

// This example demonstrates a simple data race condition.
// Two goroutines read and write to the same variable concurrently,
// causing unpredictable results.

// The race condition occurs because counter++ is not an atomic operation. It actually consists of three steps:
// Read the current value of counter
// Modify by adding 1
// Write the new value back to counter

// Here's what might happen:
// Both goroutines read counter when it's 0
// Both goroutines add 1, getting 1
// Both goroutines write 1 back to counter
// Final result is 1, even though we incremented twice, and are expected to be 2

package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	// First goroutine
	go func() {
		counter++
	}()

	// Second goroutine
	go func() {
		counter++
	}()

	// Give goroutines time to complete
	time.Sleep(time.Millisecond)
	fmt.Println("Final counter:", counter)
}
