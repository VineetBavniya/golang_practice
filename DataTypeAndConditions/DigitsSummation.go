package DataTypeAndConditions

import "fmt"

func DigitsSummation(){
	var a, b int64

	fmt.Scanf("%d %d", &a, &b)

	fmt.Printf("%d", (a%10) + (b%10))
}