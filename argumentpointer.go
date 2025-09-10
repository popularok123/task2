package main

func ArgumentPointer(arg *int) {
	*arg += 10
}

func SlicePointer(args *[]int) {
	for i := range *args {
		(*args)[i] *= 2
	}
}
