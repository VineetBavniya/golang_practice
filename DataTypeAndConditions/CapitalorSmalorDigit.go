package DataTypeAndConditions

import "fmt"

func CapitalorSmalorDigit(){
	var a byte 
	fmt.Scanf("%c",&a)

	var b int8 = int8(a)

	if(b >= 48 && b <= 57){
		fmt.Print("IS DIGIT")
	}else if (b >= 65 && b <= 90){
		fmt.Print("ALPHA\nIS CAPITAL")
	}else{
		fmt.Print("ALPHA\nIS SMALL")
	}

}