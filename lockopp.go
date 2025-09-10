package main

import (
	"fmt"
	"sync"
)

var i = 0

var mu sync.Mutex

func LockOperation() {
	mu.Lock()
	for j := 0; j < 1000; j++ {
		i++
	}
	fmt.Println("Current value:", i)
	mu.Unlock()
}

func StartLockOperation(rc int) {
	for i := 0; i < rc; i++ {
		go LockOperation()
	}
}
