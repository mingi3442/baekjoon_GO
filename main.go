package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)
	for i := 1; 10 > i; i++ {
		fmt.Println(num, "*", i, "=", num*i)
	}

}
