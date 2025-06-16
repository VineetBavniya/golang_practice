package DataTypeAndConditions

import "fmt"


func WelcomeforyouwithConditions(){
	var a, b int8 

	fmt.Scanf("%d %d", &a, &b)

	if(a >= b){
		fmt.Print("Yes")
	}else{
		fmt.Print("No")
	}
}