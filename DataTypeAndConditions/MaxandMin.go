package DataTypeAndConditions

import "fmt"

func MaxandMin(){
	var a, b , c int32

	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Printf("%d %d", min(min(a, b), c), max(max(a, b), c))
}