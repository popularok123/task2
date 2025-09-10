package main

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	println("Name:", e.Name)
	println("Age:", e.Age)
	println("EmployeeID:", e.EmployeeID)
}

func TestObjectComposition() {
	employee := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		EmployeeID: "E001",
	}

	employee.PrintInfo()
}
