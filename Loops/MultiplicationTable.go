package Loops

import "fmt"

func MultiplicationTable(){
	var n int

	fmt.Scan(&n)

	for i :=1; i<= 12; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, (n*i))
	}
}