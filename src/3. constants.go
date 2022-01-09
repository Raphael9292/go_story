package main

/*
const로 상수값을 선언합니다.
const문은 var문과 동일하게 사용할 수 있습니다.

상수 표현식은 임의의 정밀도로 산술 연산을 수행합니다.
숫자 상수는 명시적 캐스팅등으로 타입이 주어지기 전까진 타입을 가지지 않습니다.

숫자는 변수 할당이나 함수 호출과 같은 컨텍스트에서 사용하여 타입을 부여할 수 있습니다.
예를 들면, math.Sin은 float64를 기대합니다.
*/

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 10000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
