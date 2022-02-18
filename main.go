package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var h, m int
	var hour, min, add int
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n%d", &hour, &min, &add)
	// defer writer.Flush()
	h = add / 60
	m = add % 60
	hour = hour + h
	min = min + m
	if min > 59 {
		min = min - 60
		hour++
	}
	if hour > 23 {
		hour = hour - 24
	}
	// fmt.Fprintln(writer, hour, min)
	fmt.Println(hour, min)
}
