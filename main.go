package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Multiply(num1, num2 float32) float32 {
	n1 := int(num1)
	n2 := int(num2)
	str := strconv.Itoa(n2)
	arr := strings.Split(str, "")

	for i := 1; len(arr) >= i; i++ {
		e, _ := strconv.Atoi(arr[len(arr)-i])
		fmt.Println(float32(e * n1))
	}
	return num1 * num2
}

func main() {
	// var num1, num2 float32
	// fmt.Scanln(&num1, &num2)
	fmt.Println(Multiply(472, 385))
}
