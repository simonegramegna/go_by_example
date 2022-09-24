package main

import "fmt"

func main(){

	// array
	nums := [...]int{5,77,62}

	for i, num := range nums{
		fmt.Printf("i: %d, num: %d\n",i,num)
	}

	// slice
	var slice1 = make([]int, 3)
	slice1 = append(slice1, 22)

	for j, ns := range slice1{
		fmt.Printf("j: %d, ns: %d\n",j,ns)
	}

	// maps
	mymap := make(map[string]int)
	mymap["vettel"] = 5
	mymap["leclerc"] = 16
	mymap["sainz"] = 55

	for k, val := range mymap{

		fmt.Printf("chiave: %s, valore: %d\n",k,val)
	}

}