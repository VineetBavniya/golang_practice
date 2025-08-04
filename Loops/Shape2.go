package Loops

import "fmt"

func Shape2(){
	var n uint8 
	fmt.Scan(&n)

	// for i:=n; i>=1; i-- {

	// }

	for i:= uint8(0); i<n; i++ {
		
		for k:=uint8(0); k<n-i-1; k++{
			fmt.Print(" ")
			for j:=uint8(0); j<=i; j++ {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}