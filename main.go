package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var result int
	result = 0
	var a, b, c int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a, &b, &c)
	if b >= c {
		result = -1
	} else {
		result = a/(c-b) + 1
	}
	fmt.Println(result)
}
