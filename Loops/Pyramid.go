package Loops

import "fmt"

func Pyramid() {
	var n, i, j uint8

	fmt.Scan(&n)

	for i = 0; i<n; i++ {
		for j = 0; j<=i; j++ {
			fmt.Print("*")
		}

		fmt.Print("\n")
	}
}