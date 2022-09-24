package main

import "fmt"

func main(){

	age := 18

	switch age {

	case 17:
		fmt.Println("minorenne")
	
	case 18:
		fmt.Println("appena maggiorenne")
	
	default:
		fmt.Println("molto")

	}

	i := 4

	for i > 0{
		i--
	}

	fmt.Printf("%d",i)
}