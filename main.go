package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var total, flag int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	//첫 번째 입력 받기
	fmt.Fscanln(reader, &total, &flag)
	//함수 종료 직전 실행
	defer writer.Flush()
	// 두 번째 입력을 받기 위해 slice 선언
	var numbers = make([]int, total)
	//slice의 총 길이, total만큼 반복
	for i := range numbers {
		// 두 번째 입력을 slice에 담아서 입력 받기
		fmt.Fscanf(reader, "%d ", &numbers[i])
		if numbers[i] < flag {
			fmt.Fprintf(writer, "%d ", numbers[i])
		}
	}
	fmt.Fprint(writer, "\n")
}
