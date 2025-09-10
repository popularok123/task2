package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func StartLockOperation(rc int) {
	var counter int32 = 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < rc; i++ {
		fmt.Println("goroutine_lock:", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Println("Counter:", counter)
}

func StartAtomicOperation(rc int) {
	var counter int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < rc; i++ {
		fmt.Println("goroutine_atom:", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1)
			}
		}()
	}
	wg.Wait()

	fmt.Println("Counter:", counter)
}
