package DataTypeAndConditions

import "fmt"


func MathematicalExpression(){
	var a,b,c int32 
	var e,f byte 

	fmt.Scanf("%d %c %d %c %d", &a, &e, &b, &f, &c)


	if (e == '+'){
		if(a + b == c){
			fmt.Print("Yes")
		}else{
			fmt.Print(a+b)
		}
	}else if (e == '-'){
		if(a - b == c){
			fmt.Print("Yes")
		}else{
			fmt.Print(a-b)
		}

	}else if (e == '*'){
		if(a * b == c){
			fmt.Print("Yes")
		}else{
			fmt.Print(a*b)
		}
	}
}
