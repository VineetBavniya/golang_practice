package DataTypeAndConditions

import (
	"fmt"
	"math"
)

func Twonumbers(){
	var a,b float64

	fmt.Scanf("%f %f",&a, &b )

	fmt.Printf("floor %.0f / %.0f = %.0f\n", a, b, math.Floor(a/b))
	fmt.Printf("ceil %.0f / %.0f = %.0f\n", a, b, math.Ceil(a/b))
	fmt.Printf("round %.0f / %.0f = %.0f", a, b, math.Round(a/b))
}