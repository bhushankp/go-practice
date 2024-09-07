package main

import (
	"fmt"
	"sync"
)

// var counter = 0
// var mu sync.Mutex

// // we can prevent race contion using channel like ch:=make(chan bool);ch<-true;<-ch
// func incrementCounter(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mu.Lock()
// 	counter++ // Incrementing the counter without any synchronization
// 	mu.Unlock()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(100)

// 	for i := 0; i < 100; i++ {
// 		go incrementCounter(&wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Final Counter Value:", counter) // Unpredictable output due to race condition
// }

var counter = 0

// we can prevent race contion using channel like ch:=make(chan bool);ch<-true;<-ch
func incrementCounter(wg *sync.WaitGroup, ch chan bool) {

	defer wg.Done() // Mark the goroutine as done when it exits
	ch <- true      // Send a signal to the channel
	counter++       // Incrementing the counter
	<-ch            // Receive from the channel to allow other goroutines to proceed
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan bool, 1)
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go incrementCounter(&wg, ch)
	}

	wg.Wait()
	fmt.Println("Final Counter Value:", counter) // Unpredictable output due to race condition
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type BankAccount struct {
// 	balance int
// 	mu      sync.Mutex
// }

// func (b *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	b.mu.Lock()
// 	defer b.mu.Unlock()
// 	b.balance += amount
// 	fmt.Printf("Deposited %d\n", amount)
// }

// func (b *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	b.mu.Lock()
// 	defer b.mu.Unlock()
// 	if b.balance >= amount {
// 		b.balance -= amount
// 		fmt.Printf("Withdrawn %d\n", amount)
// 	} else {
// 		fmt.Println("Insufficient balance")
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	account := BankAccount{balance: 100}

// 	// Launch multiple goroutines to perform deposit and withdrawal operations concurrently
// 	for i := 0; i < 5; i++ {
// 		wg.Add(2) // Each deposit/withdrawal operation involves 2 goroutines

// 		// Concurrently deposit and withdraw
// 		go func() {
// 			account.Deposit(50, &wg)
// 		}()
// 		go func() {
// 			account.Withdraw(70, &wg)
// 		}()
// 	}

// 	wg.Wait() // Wait for all deposit and withdrawal operations to finish

// 	fmt.Printf("Final balance: %d\n", account.balance)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter = 0
// var mu sync.Mutex

// func incrementCounter(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mu.Lock()
// 	counter++ // Incrementing the counter without any synchronization
// 	mu.Unlock()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(100)

// 	for i := 0; i < 100; i++ {
// 		go incrementCounter(&wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Final Counter Value:", counter) // Unpredictable output due to race condition
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {

// 	var wg sync.WaitGroup
// 	var counter int
// 	var mu sync.Mutex

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			mu.Lock()
// 			counter++
// 			mu.Unlock()
// 		}()

// 	}

// 	wg.Wait()
// 	fmt.Println(counter)
// }
