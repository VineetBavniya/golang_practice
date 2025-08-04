package Loops

import (
	"fmt"
	"math"
)

func Check(n int) bool {

	for i :=2; i<=int(math.Sqrt(float64(n))); i++ {
		if (n%i==0){
			return false 
		}
	}

	return true 
}


func PrimesFrom1ToN(){
	var n int
	fmt.Scan(&n)

	for i:=2; i<=n; i++ {
		primeNum := Check(i)
		if(primeNum){
			fmt.Printf("%d ", i)
		}
	}
}