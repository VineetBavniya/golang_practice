package DataTypeAndConditions

import "fmt"


func CoordinatesofaPoint(){
	var y,x float64 
	fmt.Scan(&x, &y)

	if(y == 0 && x == 0){
		fmt.Printf("Origem")
	}else if(y > 0 && x > 0){
		fmt.Printf("Q1")
	}else if(y > 0 && x < 0){
		fmt.Printf("Q2")
	}else if(y < 0 && x < 0){
		fmt.Printf("Q3")
	}else if(y < 0 && x > 0){
		fmt.Printf("Q4")
	}else if ((y > 0 || y < 0) && x == 0){
		fmt.Print("Eixo Y")
	}else if(y == 0 && (x > 0 || x < 0 )){
		fmt.Print("Eixo X")
	}
}