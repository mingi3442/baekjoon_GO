package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var total int
	var a, b int
	//reader를 통해 읽기 인스턴스 생성
	reader := bufio.NewReader(os.Stdin)
	//write를 통해 쓰기 인스턴스 생성
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &total)

	for i := 0; total > i; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintln(writer, a+b)
	}
	//데이터 출력
	writer.Flush()
}
