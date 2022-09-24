package main

import "fmt"

func sum(v1 int, v2 int) int{
	return v1 + v2
}

// variadic function
func sum_seq(seq...int) int{

	var s = 0

	for _, val := range seq{
		s = s + val
	}
	return s
}

func sum_k(k int, sequence...int) []int{

	var sum_slice = []int{}

	for _, s := range sequence{
		sum_slice = append(sum_slice, s+k)
	}
	return sum_slice
}


func get_drivers()map[string]int{

	drivers := make(map[string]int)

	drivers["leclerc"] = 16
	drivers["sainz"] = 55

	return drivers

}

func main(){

	sumab := sum(11,55)
	fmt.Println(sumab)

	sumseq1 := sum_seq(12,22)
	sumseq2 := sum_seq(1,1,1,1)

	fmt.Printf("%d %d\n\n",sumseq1, sumseq2)

	slice_sum := sum_k(100, 10, 20, 30)
	fmt.Println(slice_sum)
	
	drivers := get_drivers()
	fmt.Println(drivers)


}