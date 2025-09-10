package main

func main() {

	iv := 10

	ArgumentPointer(&iv)
	SlicePointer(&[]int{1, 2, 3, 4, 5})

	StartGoroutine()
	GoTasks()

	TestInterfaceImplementation()
	TestObjectComposition()

	MultipleChannelCommunication()

	StartLockOperation(10)

}
