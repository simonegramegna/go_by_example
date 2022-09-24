package main

import "fmt"

func main(){
	var arr_0[2]int

	arr_0[0] = 1
	arr_0[1] = 50

	fmt.Println(arr_0)

	var arr_1 = [3]int{3,4,5}
	fmt.Println(arr_1)

	arr_2 := [5]float32{2.2, 4, 5.5, 6, 0}
	fmt.Println(arr_2)

	var matrix [3][3]int

	for i:=0; i < 3; i++{
		for j:=0; j < 3; j++{
			matrix[i][j] = 0
		}
	}

	fmt.Println(matrix)

}