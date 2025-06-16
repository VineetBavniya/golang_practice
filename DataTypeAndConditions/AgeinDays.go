package DataTypeAndConditions

import "fmt"


func AgeinDays(){
	var Days int32
	fmt.Scanf("%d", &Days)
	
	var years int32 = Days/365
	var Months int32 = (Days - (365*years))/30
	var Day int32 = (Days - (365*years) - Months*30)

	fmt.Printf("%d years\n", years)
	fmt.Printf("%d months\n", Months)
	fmt.Printf("%d days", Day)
}