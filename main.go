package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(a int) {
			fmt.Printf("without lock %d \n", a)
			wg.Done()
		}(i)
	}

	wg.Wait()

	mutex := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		mutex.Lock()
		go func(a int) {

			defer mutex.Unlock()
			fmt.Printf("with lock %d \n", a)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
