// problemName -> B. Basic Data Types

package DataTypeAndConditions

import "fmt"

func BasicDataType(){
	var n32 int32
	var n64 int64 

	var c byte

	var f32 float32
	var f64 float64


	fmt.Scanf("%d %d %c %f %f", &n32, &n64, &c, &f32, &f64)

	fmt.Printf("%d\n", n32)
	fmt.Printf("%d\n", n64)
	fmt.Printf("%c\n", c)
	fmt.Printf("%.2f\n", f32)
	fmt.Printf("%.1f\n", f64)
}