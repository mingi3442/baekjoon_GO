package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var newNum int
	idx := 1
	temp := 1
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &newNum)
	maxNum := newNum
	defer writer.Flush()
	for i := 0; 8 > i; i++ {
		fmt.Fscanln(reader, &newNum)
		temp++
		if maxNum < newNum {
			maxNum = newNum
			idx = temp
		}
	}
	if maxNum == newNum {
		idx = 9
	}
	// fmt.Println(maxNum)
	// fmt.Println(idx)
	fmt.Fprintf(writer, "%d\n%d\n", maxNum, idx)
}
