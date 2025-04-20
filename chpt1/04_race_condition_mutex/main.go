package main

// Example: Data Race Condition without Mutex
// Counter value is not predictable, and is not 1000

import (
	"fmt"
	"sync"
)

// Example: Data Race Condition without Mutex
func noMutex() {
	var counter int
	var wg sync.WaitGroup

	// Increment counter concurrently.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++ // Data race occurs here.
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter value:", counter)

}

// Example: Data Race Condition with Mutex
// Counter value is predictable, and always is 1000
func muTex() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Safely increment counter using a mutex.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter value:", counter)
}

func main() {
	noMutex()
	muTex()
}
