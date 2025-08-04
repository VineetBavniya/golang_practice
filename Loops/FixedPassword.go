package Loops

import "fmt"

func FixedPassword(){
	for {
		var a uint16 
		fmt.Scan(&a)

		if(a == 1999){
			fmt.Print("Correct")
			return
		}else{
			fmt.Println("Wrong")
		}
	}

}