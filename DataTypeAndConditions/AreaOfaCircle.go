// AreaofaCircle

package DataTypeAndConditions

import "fmt"

func AreaofaCircle(){
	const pi float64 = 3.141592653
	var r float64
	fmt.Scanf("%f", &r)
	var area float64 = pi * (r*r)

	fmt.Printf("%.9f", area)
}