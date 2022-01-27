package main

import (
	"fmt"
)

func main() {
	var total int
	fmt.Scanf("%d", &total)
	for i := 1; total+1 > i; i++ {
		var numA, numB int
		fmt.Scanln(&numA, &numB)
		fmt.Println(numA + numB)
	}

}
