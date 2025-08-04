package Loops

import (
	"fmt"
)

func rev(n uint ) {
	for n > 0 {
		rem := n%10
		n/=10
		fmt.Printf("%d ",rem)
	}
	fmt.Print("\n")
}


func Digits()  {
	var testCase uint8 
	fmt.Scan(&testCase)

	for range(testCase) {
		var n uint 
		fmt.Scan(&n) 
		rev(n)
	}
}