package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d", &a, &b) // a,b를 입력 받는다
	if a < b {
		a, b = b, a // 나머지 연산을 하기 위해 b가 a보다 클 경우 자리를 바꿔준다
	}
	gcd := 0 // 최대공약수 선언
	for {
		r := a % b
		if r == 0 {
			//나머지 연산 값이 0일 경우 최대 공약수 값은 0이 되며 반복문은 종료
			gcd = b
			break
		}
		// 나머지 연산 값이 0이 아닐 경우  a에는 b를, b에는 나머지 연산 값을 선언해서 반복
		a = b
		b = r
	}
	fmt.Println(gcd)
}
