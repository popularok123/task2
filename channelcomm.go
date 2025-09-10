package main

func produceNumbers(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func printNumbers(ch chan int, done chan bool) {
	for num := range ch {
		println(num)
	}
	done <- true
}

func ChannelCommunication() {
	ch := make(chan int)
	done := make(chan bool)

	go produceNumbers(ch)
	go printNumbers(ch, done)

	<-done
	close(ch)
	close(done)
}

func Producer(ch []chan int, count int) {
	for i := 0; i < count; i++ {
		ch[i%len(ch)] <- i
	}
	for _, c := range ch {
		close(c)
	}
}

func Consumer(ch chan int, done chan bool) {
	for num := range ch {
		println("received", num)
	}
	done <- true
}

func MultipleChannelCommunication() {
	numChannels := 10
	ch := make([]chan int, numChannels)
	done := make(chan bool)

	for i := 0; i < numChannels; i++ {
		ch[i] = make(chan int)
		go Consumer(ch[i], done)
	}

	Producer(ch, 100)

}
