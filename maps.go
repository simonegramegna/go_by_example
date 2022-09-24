package main

import "fmt"

func main(){

	var mymap = make(map[string]int)

	mymap["leclerc"] = 16
	mymap["gasly"] = 10
	mymap["russel"] = 63

	russel, russel_exists := mymap["russel"]
	fmt.Println(russel)
	fmt.Println(russel_exists)	

}