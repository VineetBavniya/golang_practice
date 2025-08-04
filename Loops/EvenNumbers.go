package Loops

import "fmt"

func EvenNumbers(){
	var a uint16 

	fmt.Scan(&a)
	var i uint16
	for i = 2; i<=a; i+=2 {
		fmt.Println(i)
	}
}