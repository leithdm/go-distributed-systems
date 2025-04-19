// In this example, the greet function is called concurrently for “Alice” and “Bob”.
// The main function starts two goroutines to greet Alice and Bob concurrently.
// The program then sleeps for one second to allow the goroutines to complete before exiting.

package main

import (
	"fmt"
	"time"
)

// greet outputs a greeting message for the given name.
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// Start two goroutines to greet Alice and Bob concurrently.
	go greet("Alice")
	go greet("Bob")

	// Allow goroutines time to complete before the program exits.
	time.Sleep(time.Second)
}
