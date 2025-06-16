package DataTypeAndConditions

import "fmt"

func FloatOrInt(){
	var n float64
	fmt.Scan(&n)

	var intValue int32 = int32(n)
	
	var flaotValue float64 = n - float64(intValue)

	if (flaotValue > 0) {
		fmt.Printf("float %d %.3f", intValue, flaotValue)
	}else{
		fmt.Printf("int %d", intValue)
	}
	
}