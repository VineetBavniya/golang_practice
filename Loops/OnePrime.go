package Loops

import (
	"fmt"
	"math"
)

func OnePrime(){
	var n int
	fmt.Scan(&n)

	for i :=2; i<=int(math.Sqrt(float64(n))); i++ {
		if (n%i==0){
			fmt.Print("NO")
			return
		}
	}

	fmt.Print("YES")
}