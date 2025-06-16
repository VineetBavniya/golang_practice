package DataTypeAndConditions

import "fmt"


func Summationfrom1toN(){
	var n int64
	fmt.Scanf("%d", &n)

	n = (n * (n + 1))/2

	fmt.Printf("%d", n)
}