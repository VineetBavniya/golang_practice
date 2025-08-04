package Loops

import "fmt"

func CheckLuckyNumber(n int ) bool {
	
	for n > 0 {
		rem := n%10
		n/=10

		if (rem == 4){
			continue
		}else if (rem == 7){
			continue
		}else {
			return false
		}
		
	}

	return true 
}


func LuckyNumbers(){
	var a, b int 
	fmt.Scan(&a, &b)
	for i := a; i<=b; i++ {
		if(CheckLuckyNumber(i)){
			fmt.Printf("%d ", i)
		}

	}
}