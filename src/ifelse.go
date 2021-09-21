package main

/*
else 없이 if문을 사용할 수 있습니다.
조건절 전에 명령문이 위치할 수 있습니다. 이 명령문에서 선언된 모든 변수들은 모든 분기문에서 사용 가능합니다.
Go에서 주의할 건 조건절엔 괄호가 필요없지만, 중괄호는 필수입니다.
Go에는 삼항 조건문이 없습니다. 따라서 간단한 조건문이라도 완전한 if문을 사용해야 합니다.
*/

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
