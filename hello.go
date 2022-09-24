package main

import "fmt"

func main(){
	
	const year = 2000

	var age int
	var name string
	var tall float32
	gender := "m"

	fmt.Scan(&age, &name, &tall)
	

	fmt.Printf("Nome: %s, eta: %d, sesso: %s, altezza: %f, anno: %d", name, age, gender, tall, year)
}