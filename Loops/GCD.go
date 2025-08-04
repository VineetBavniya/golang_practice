package Loops

import (
	"fmt"
	"math"
)

// func EcuidAlgo(a int, b int ) int {
// 	if (b <= 0){
// 		return a
// 	}
// 	return EcuidAlgo(b, b%a)
// }


func GCD(){
	var n, m, gcd int 
	fmt.Scan(&n, &m)

	for i:=1; i<=int(math.Min(float64(n), float64(m))); i++ {
		if(n%i == 0 && m % i == 0){
			gcd = i
		}
	}
	fmt.Print(gcd)
}
