package main

/*
기본적인 초기화/조건/후연산 for 루프
조건이 없는 for는 루프 밖으로 break 하거나 함수 내부에서 return을 하기 전까지 계속 반복합니다.
루프의 다음 반복을 위해 continue를 사용할 수도 있습니다.
*/

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
