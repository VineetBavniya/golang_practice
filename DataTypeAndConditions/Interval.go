package DataTypeAndConditions

import "fmt"


func Interval(){
	var n float64 

	fmt.Scan(&n)

	if(n >= 0 && n <= 25){
		fmt.Print("Interval [0,25]")
	}else if(n > 25 && n <= 50){
		fmt.Print("Interval (25,50]")
	}else if(n > 50 && n <= 75){
		fmt.Print("Interval (50,75]")
	}else if(n > 75 && n <= 100){
		fmt.Print("Interval (75,100]")
	}else{
		fmt.Print("Out of Intervals")
	}
}