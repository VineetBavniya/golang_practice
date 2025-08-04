package Loops

import "fmt"

func Palindrome(){
	var n int
	fmt.Scan(&n)
	var org int = n 
	var rev int = 0
	for 0 < n {
		rem := n%10
		n /=10
		rev = rev*10 + rem
	}

	if (org == rev){
		fmt.Printf("%d\nYES", org)
	}else{
		fmt.Printf("%d\nNO", org)
	}
}
