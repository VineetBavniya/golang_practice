package DataTypeAndConditions

import "fmt"

func Comparison(){
	var a , b int32 
	var c byte 

	fmt.Scanf("%d %c %d", &a, &c , &b)


	if(c == '>'){
		if(a > b ) {
			fmt.Println("Right")
		}else{
			fmt.Println("Wrong")
		}
	}else if(c == '<'){
		if(a < b ) {
			fmt.Println("Right")
		}else{
			fmt.Println("Wrong")
		}
	}else if(c == '='){
		if(a == b ) {
			fmt.Println("Right")
		}else{
			fmt.Println("Wrong")
		}
	}
}