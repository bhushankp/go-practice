package main

import (
	"fmt"
	"sync"
	"time"
)

func WorkerTask(workerID int, task string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second * 2)
	fmt.Printf("Worker %d Completed  %s\n", workerID, task)

}

func main() {

	numbOfWorkers := 3
	//tasks := []string{"cleanning", "planting", "painting"}

	var wg sync.WaitGroup

	fmt.Println("Workers start their work")
	for i := 0; i < numbOfWorkers; i++ {
		//for _, task := range tasks {
		wg.Add(1)
		go WorkerTask(i+1, fmt.Sprintf("task %d", i+1), &wg)

		//}
	}
	wg.Wait()
	fmt.Println("All task Completed ")
}

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Function to represent a task that a worker performs
// func WorkerTask(workerID int, task string, wg *sync.WaitGroup) {
// 	defer wg.Done() // Notify the WaitGroup when the goroutine completes

// 	// Simulating the time it takes for the task to complete
// 	time.Sleep(time.Second * 2)

// 	fmt.Printf("Worker %d completed task: %s\n", workerID, task)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Number of workers and tasks
// 	numWorkers := 3
// 	tasks := []string{"Cleanning", "Planting", "Watering"}

// 	fmt.Println("Workers Started their Works")

// 	// Launch goroutines for each worker and task
// 	for i := 0; i < numWorkers; i++ {
// 		for _, task := range tasks {
// 			wg.Add(1) // Increment the WaitGroup counter for each goroutine

// 			// Goroutine for parallel execution
// 			go WorkerTask(i+1, task, &wg)
// 		}
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	fmt.Println("All tasks completed.")
// }
