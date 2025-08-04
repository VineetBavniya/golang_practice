package Loops

import (
	"fmt"
	"math"
)

func SequenceOfNumbersAndSum(){
	
	for{
		var n,m uint8 
		fmt.Scanf("%d %d",&n, &m)

		if(n <= 0 || m <= 0) {
			return
		}else {
			var nMin int = int(math.Min(float64(n), float64(m)))- 1
			var nSum int = (nMin * (nMin + 1))/2

			var mMax int = int(math.Max(float64(n), float64(m)))
			var mSum int = (mMax * (mMax + 1))/2

			var i uint8 

			for i = uint8(nMin)+1; i<=uint8(mMax); i++ {
				fmt.Printf("%d ", i)
			}
			fmt.Printf("sum =%d\n", mSum - nSum)
		}

	}
	
}