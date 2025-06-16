package DataTypeAndConditions

import (
	"fmt"
	"math"
)

func HardCompare(){
	var a, b, c, d float64
    fmt.Scanf("%f %f %f %f", &a, &b, &c, &d)

    left := b * math.Log(a)
    right := d * math.Log(c)

    if left > right {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}


// Notes:
// math.Log is the natural logarithm (ln)
// This method avoids overflow and is accurate for comparing large exponentials
// No need for math/big or slow big exponentiation