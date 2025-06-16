package DataTypeAndConditions

import "fmt"

func Thelast2digits(){
	var a,b,c,d uint64 

	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	var ans uint64 =( a%100 * b%100 * c%100 * d%100) % 100

	if(ans < 10){
		fmt.Printf("0%d",ans)
	}else {
		fmt.Print(ans)
	}
}
