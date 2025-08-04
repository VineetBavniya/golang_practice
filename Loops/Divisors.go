package Loops

import "fmt"

func Divisors() {
	var n uint16
	fmt.Scan(&n)

	var i uint16 

	for i = 1; i<=n; i++ {
		if(n%i == 0){
			fmt.Printf("%d\n", i)
		}
	}
}