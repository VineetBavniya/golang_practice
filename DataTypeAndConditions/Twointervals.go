package DataTypeAndConditions

import (
	"fmt"
)

func TwoIntervals(){
	var l1,l2,r1,r2 int64 

	fmt.Scanf("%d %d %d %d",&l1, &r1, &l2, &r2)

	if(max(l1,l2) > min(r1, r2)){
		fmt.Print("-1")
	}else{
		fmt.Printf("%d %d", max(l1, l2), min(r1, r2))
	}
}