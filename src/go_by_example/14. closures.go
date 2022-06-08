package main

/*
Go는 closures를 만들 수 있는 anonymous functions를 지원합니다.
익명 함수(Anonymous functions)는 이름 없이 한줄로 함수를 정의하고 싶을때 유용합니다.
*/
import "fmt"

/*
intSeq 함수는 내부에 익명으로 정의한 또 다른 함수를 반환합니다.
반환된 함수는 클로저를 만들기 위해 i 변수를 가둬둡니다(closes over).
*/
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	//intSeq를 호출하여, 결괏값(함수)을 nextInt에 할당합니다. 이 함숫값은 nextInt를 호출할 때마다 업데이트 되는 i 값을 캡쳐합니다.
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//특정 함수에서 상태값이 유일한지 확인하기 위해, 하나를 새롭게 생성하고 테스트 해봅니다.
	newInts := intSeq()
	fmt.Println(newInts())
}
