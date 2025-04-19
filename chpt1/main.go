package main

import (
	"fmt"
	"sync"
)

// worker simulates a concurrent task.
func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    // Simulate work with a placeholder.
    fmt.Printf("Worker %d completed\n", id)
}

func main() {
    var wg sync.WaitGroup

    // Start several worker goroutines.
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait() // Wait for all goroutines to finish.
    fmt.Println("All workers have finished")
}
