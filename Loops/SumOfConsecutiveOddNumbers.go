package Loops

import "fmt"

func SumOfConsecutiveOddNumbers(){
	var testCase uint8 

	fmt.Scan(&testCase)

	for range testCase {
		var n , m int 
		fmt.Scan(&n, &m)
		sum := 0
		if(n >= m ){
			for i:=m+1; i<n; i++ {
				if(i%2 == 1){
					sum+=i
				}
			}
			fmt.Printf("%d\n",sum)
		}else {
			for i:=n+1; i<m; i++ {
				if(i%2 == 1){
					sum+=i
				}
			}
			fmt.Printf("%d\n",sum)
		}
		
	}
}

