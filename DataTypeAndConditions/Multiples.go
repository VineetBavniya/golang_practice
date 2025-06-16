package DataTypeAndConditions

import "fmt"

func Multiples(){
	var a, b uint32
	
	fmt.Scanf("%d %d", &a, &b)

	if((a%b) == 0 || (b%a) == 0){
		fmt.Print("Multiples")
	}else{
		fmt.Print("No Multiples")
	}
}