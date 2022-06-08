package main

//variadic-functions: 가변 함수
import "fmt"

// int형 인자를 임의의 개수만큼 받는 가변 함수
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
}

func main() {
	//가변 함수는 일반적인 방법인 개별적인 인자들로 호출할 수 있습니다.
	sum(1, 2)
	sum(1, 2, 3)

	//만약 값들이 이미 슬라이스에 들어 있다면, 다음과 같이 func(slice...)를 사용하여 가변 함수의 인자에 적용하세요.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
