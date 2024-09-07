package main

import (
	"fmt"
	"sync"
	"time"
)

func chefTask(chefID int, task string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate the chef working on a task
	fmt.Printf("Chef %d is preparing %s\n", chefID, task)
	time.Sleep(time.Second * time.Duration(chefID+1))
	fmt.Printf("Chef %d finished preparing %s\n", chefID, task)
}

func main() {
	// Number of chefs (tasks)
	numChefs := 3

	//dishes := []string{"dish 1", "dish 2", "dish 3"}
	// Use a WaitGroup to wait for all chefs to finish
	var wg sync.WaitGroup

	// Start goroutines for each chef (task)
	for i := 1; i <= numChefs; i++ {

		wg.Add(1)

		go chefTask(i, fmt.Sprintf("dish %d", i), &wg)
	}

	// Wait for all chefs to finish
	wg.Wait()

	fmt.Println("All tasks in the kitchen completed.")
}
