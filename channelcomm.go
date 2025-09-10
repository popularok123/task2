package main

import (
	"fmt"
	"sync"
)

func produceNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func printNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		println(num)
	}
}

func ChannelCommunication() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go produceNumbers(ch, &wg)
	go printNumbers(ch, &wg)

	wg.Wait()

}

// Producer 生产者函数：向通道发送 100 个整数
func Producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch) // 生产完成后关闭通道
}

// Consumer 消费者函数：从通道接收整数并打印
func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("received:", num)
	}
}

func MultipleChannelCommunication() {
	ch := make(chan int, 10) // 带缓冲的通道
	var wg sync.WaitGroup

	// 启动生产者
	wg.Add(1)
	go Producer(ch, &wg)

	// 启动消费者
	wg.Add(1)
	go Consumer(ch, &wg)

	// 等待完成
	wg.Wait()

}
