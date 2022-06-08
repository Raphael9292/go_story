package main

//다중 반환값
import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	//다음은 함수 호출로부터 반환되는 두 반환값을 다중 할당(multiple assignment)으로 받습니다.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//반환값들중 일부만 사용하고 싶을땐, 공백 식별자 _을 사용합니다.
	_, c := vals()
	fmt.Println(c)
}
