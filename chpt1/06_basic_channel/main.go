package main

import (
	"fmt"
)

// generateNumbers sends a sequence of numbers to the provided channel.
func generateNumbers(count int, ch chan int) {
	for i := 0; i < count; i++ {
		ch <- i*10
	}
	close(ch) // Close the channel to signal completion.
}

func main() {
	numbers := make(chan int)

	// Start a goroutine to generate numbers.
	go generateNumbers(5, numbers)

	// Receive and print numbers from the channel.

	for num := range numbers {
		fmt.Println(num)
	}
}
