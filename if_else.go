package main

import "fmt"

func main(){

	cond := true
	var x = 0

	if cond && x > 0 {
		fmt.Print("condizione vera")

	}else if cond || x == 0{

		fmt.Printf("Condizione vera valore: %d",x)
	}else{
		fmt.Printf("%d", x)
	}

}