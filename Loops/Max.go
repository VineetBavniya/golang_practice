package Loops

import (
	"fmt"
	"math"
)

func Max()  {
    var n int32 
    var mx int = math.MinInt 
    var a int 
    fmt.Scan(&n)
    for range n {
        fmt.Scan(&a)

        if(a > mx){
           mx = a
        }
    }

    fmt.Print(mx)
	
}