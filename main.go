package main

import "fmt"

func main() {

	iv := 10
	ArgumentPointer(&iv)
	fmt.Println("Value after pointer:", iv)

	ia := []int{1, 2, 3, 4, 5}
	SlicePointer(&ia)
	fmt.Println("Value after slice pointer:", ia)

	StartGoroutine()
	GoTasks()

	TestInterfaceImplementation()
	TestObjectComposition()

	ChannelCommunication()
	MultipleChannelCommunication()

	StartLockOperation(10)
	StartAtomicOperation(10)

}
