package main

import (
	"fmt"
	"sync"
	"time"
)

func printOddNumbers() {
	for i := 1; i < 10; i += 2 {
		println(i)
	}
}

func printEvenNumbers() {
	for i := 2; i < 10; i += 2 {
		println(i)
	}
}

func StartGoroutine() {
	go printOddNumbers()
	go printEvenNumbers()
	select {}
}

type Task func()

type Result struct {
	ID       int
	Duration time.Duration
}

type TaskSchedule struct {
	tasks []Task
}

func NewTaskSchedule(tasks []Task) *TaskSchedule {
	return &TaskSchedule{tasks: tasks}
}

func (ts *TaskSchedule) start() []Result {
	var wg sync.WaitGroup
	results := make([]Result, len(ts.tasks))
	for i, task := range ts.tasks {
		wg.Add(1)
		go func(i int, task Task) {
			defer wg.Done()
			start := time.Now()
			task()
			results[i] = Result{ID: i, Duration: time.Since(start)}
		}(i, task)
	}
	wg.Wait()
	return results
}

func GoTasks() {
	ts := NewTaskSchedule([]Task{printEvenNumbers, printOddNumbers})
	results := ts.start()
	for _, result := range results {
		fmt.Printf("Task %d took %s\n", result.ID, result.Duration)
	}
}
