package main

import "fmt"

/*
재귀의 대명사 팩토리얼 예제
fact 함수는 베이스 케이스인 fact(0)에 도달할 때까지 자기자신을 호출합니다.
*/
func fact(n int) int {
	// 이게 있어야 0을 곱하지 않는다.
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(4))
}
