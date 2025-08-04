package Loops

import "fmt"

func OneToN() {
	var a int16
	fmt.Scan(&a)
	var i int16 
	for  i = 1; i<=a; i++{
		fmt.Println(i)
	}
}