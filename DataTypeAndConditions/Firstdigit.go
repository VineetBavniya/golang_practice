package DataTypeAndConditions

import "fmt"


func Firstdigit()  {
	var a uint16 

	fmt.Scanf("%d", &a)

	if((a/1000)%2 == 0){
		fmt.Print("EVEN")
	}else{
		fmt.Print("ODD")
	}

}