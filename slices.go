package main

import "fmt"

func main(){

	var slice_0 = make([]int, 3)
	slice_0[0] = 1
	slice_0 = append(slice_0, 4)
	slice_0 = append(slice_0, 24)

	fmt.Println(slice_0)


	slice_1 := make([]string, 3)
	slice_1[0] = "simone"
	slice_1[1] = "leo"
	slice_1[2] = "ste"

	slice_1 = append(slice_1, "fede")

	fmt.Println(slice_1)

	slice_2 := []float32{3.2, 4, 5}
	fmt.Println(slice_2)

	// copia di uno slice in un altro
	slice_src := []int{22, 11, 29}
	var slice_d = make([]int, 2)
	copy(slice_d, slice_src)

	fmt.Println(slice_d)

}