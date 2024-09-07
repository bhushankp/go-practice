package main

import (
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf(" Worker %d processing taskt %d \n", id, task)
		results <- task * 2
	}

}

func main() {

	numOfWorker := 3
	numOfTasks := 5

	var wg sync.WaitGroup

	tasks := make(chan int, numOfTasks)
	results := make(chan int, numOfTasks)

	for i := 1; i <= numOfWorker; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	for i := 1; i <= numOfTasks; i++ {
		tasks <- i
	}

	close(tasks)

	wg.Wait()
	close(results)
	for result := range results {
		fmt.Printf("Received Result : %d\n", result)
	}
}
