package Loops

import "fmt"

func NumbersHistogram() {
	var ch byte
	var n uint8


	fmt.Scanf("%c ",&ch)
	fmt.Scan(&n)


	for range n {
		var nums uint8 
		fmt.Scan(&nums)

		var j uint8 

		for j = 0; j<nums; j++ {
			fmt.Printf("%c", ch)
		}
		fmt.Print("\n")
	}
}