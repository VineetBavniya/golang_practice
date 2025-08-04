package Loops

import (
	"fmt"
)

func Ans(n int) int {
	if(n <= 1){
		return 1;
	}
	return n * Ans(n-1)
}


func Factorial(){
	var n int8 
	fmt.Scan(&n)

	for range n {
		var a int 
		fmt.Scan(&a)
		ans := Ans(a)
		fmt.Println(ans)
	}
}