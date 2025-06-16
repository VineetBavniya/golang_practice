package DataTypeAndConditions

import "fmt"

func Calculator(){
	var a,b int32
	var c byte

	fmt.Scanf("%d%c%d",&a, &c, &b)

	if(c == '+'){
		fmt.Printf("%d", a+b)
	}else if(c == '-'){
		fmt.Printf("%d", a-b)
	}else if(c == '*'){
		fmt.Printf("%d", a*b)
	}else if(c == '/'){
		fmt.Printf("%d", a/b)
	}
}