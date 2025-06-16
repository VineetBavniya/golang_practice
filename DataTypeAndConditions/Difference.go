package DataTypeAndConditions

import "fmt"

func Difference(){
	var a, b, c, d int64 
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	var ans int64 = (a*b) - (c*d)

	fmt.Printf("Difference = %d", ans)
}