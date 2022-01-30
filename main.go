package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var number, temp int
	cycle := 0
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &number)
	//newN에 입력받은 number 값 넣기
	var newN = number

	defer writer.Flush()
	for true {
		//반복될 때 마다 cycle +1
		cycle++
		if newN < 10 {
			//newN이 10보다 작을 경우
			temp = newN
		} else {
			//10보다 클 경우
			temp = newN/10 + newN%10
		}
		//다시 newN에 입력
		newN = newN%10*10 + temp%10
		if newN == number {
			break
		}
	}
	fmt.Fprintln(writer, cycle)
}
