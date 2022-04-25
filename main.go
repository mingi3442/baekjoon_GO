package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdin)
	fmt.Fscanf(reader, "%d %d", &a, &b)
	// defer writer.Flush()
	r1 := a
	r2 := b
	t1 := 0
	t2 := 1
	for r2 > 0 {
		q := r1 / r2
		r := r1 - q*r2
		r1, r2 = r2, r
		t := t1 - q*t2
		t1, t2 = t2, t
		r = a % b
		if r1 != 1 {
			fmt.Println(0)
		}
		if t1 < 0 {
			t1 = a + t1
		}
	}
	// fmt.Fprintln(writer, hour, min)
	fmt.Println(t1)
}
