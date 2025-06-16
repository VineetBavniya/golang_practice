package DataTypeAndConditions

import "fmt"

func Char(){
	var a byte
	fmt.Scanf("%c", a)
	
	var c int8 = int8(a)

	if (c >= 97 && c <= 122){
		fmt.Printf("%c", c-32)
	}else{
		fmt.Printf("%c", c+32)
	}
}