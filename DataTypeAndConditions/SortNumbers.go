package DataTypeAndConditions

import "fmt"

func SortNumber(){
	var a [3]int32
	var aa [3]int32 
	for i := 0; i<3; i++ {
		var b int32 
		fmt.Scan(&b)
		a[i] = b
	}
	aa = a 
	for i := 0; i<3; i++ {
		for j :=0; j<3-i-1; j++ {
			if( a[j] > a[j+1] ){
				temp := a[j+1]
				a[j+1] = a[j]
				a[j] = temp 
			}
		}
	}

	for i := 0; i<3; i++ {
		fmt.Println(a[i])
	}
	fmt.Println()
	for i := 0; i<3; i++ {
		fmt.Println(aa[i])
	}
}