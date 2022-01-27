package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var total int
	var numA, numB int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &total)
	for i := 1; total+1 > i; i++ {
		fmt.Fscanln(reader, &numA, &numB)
		fmt.Fprintf(writer, "Case #%d: %d\n", i, numA+numB)
	}
	writer.Flush()
}
