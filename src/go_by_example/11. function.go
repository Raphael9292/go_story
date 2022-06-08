package main

import "fmt"

//Go에선 명시적 반환을 해야합니다. 즉, 마지막 식의 값을 자동으로 반환하지 않습니다.
func plus(a int, b int) int {
	return a + b
}

//만약 같은 타입을 갖는 파라미터들이 연속적으로 주어지면, 마지막 파라미터만 타입을 선언하고 나머지 파라미터들에 대해서는 타입명을 생략할 수 있습니다.
func plusPlus(a, b, c int) int {
	return a + b + c
}

//name(args) return_type
func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
