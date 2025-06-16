package DataTypeAndConditions

import "fmt"

func TheBrothers(){
	var f1,f2, s1, s2 string

	fmt.Scanln("%s %s", &f1, &s1)
	fmt.Scanln("%s %s", &f2, &s2)

	if(s1 == s2 ){
		fmt.Print("ARE Brothers")
	}else{
		fmt.Print("NOT")
	}
}