package Loops

import "fmt"

func Shape1(){
	var n,i,j uint8 

	fmt.Scan(&n)

	for i = n; i>=1; i-- {
		for j=1; j<=i; j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}