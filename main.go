package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numA, numB int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for true {
		fmt.Fscanln(reader, &numA, &numB)
		if numA == 0 && numB == 0 {
			break
		}
		fmt.Fprintln(writer, numA+numB)
	}

	fmt.Fprint(writer, "\n")
}
