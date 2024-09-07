package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Print(&wg)

	}
	wg.Wait()

}

func Print(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("go routine call")
}
