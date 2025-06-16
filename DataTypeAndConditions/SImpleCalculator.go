// C. Simple Calculator

package DataTypeAndConditions

import "fmt"


func SimpleCalculator(){
	var n int64 
	var m int64 

	fmt.Scanf("%d %d", &n, &m)

	fmt.Printf("%d + %d = %d\n", n, m, (n + m))
	fmt.Printf("%d * %d = %d\n", n, m, (n * m))
	fmt.Printf("%d - %d = %d\n", n, m, (n - m))
}
