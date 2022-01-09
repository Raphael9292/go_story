package main

/*
배열과 다르게, 슬라이스는 원소의 갯수가 아닌 포함된 원소들로만 작성됩니다.
*/
import "fmt"

func main() {
	//길이가 0인 빈 배열을 만들기 위해선 내장 함수 make를 사용하면 됩니다. 제로값으로 초기화된 길이가 3인 문자열 배열을 만듭니다.
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	//이러한 기본 연산에 이어 슬라이스는 배열보다 더 풍부한 기능들을 지원합니다. 그 중 하나는 새로운 값이 추가된 슬라이스를 반환하는 append 함수입니다.
	//주의할건 새로운 슬라이스를 사용하기위해선 append로부터 반환되는 값을 사용해야합니다.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	//다음처럼 쓰면 s[5]까지 자릅니다. (단, 마지막 원소는 포함하지 않습니다.)
	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	//한 줄로 슬라이스 선언 및 초기화를 할 수도 있습니다.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//슬라이스는 다차원 데이터를 구성할 수도 있습니다 다차원 배열과 다르게 다중 슬라이스의 내부 슬라이스들은 가변적 길이를 가질 수 있습니다.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	//참고로 슬라이스는 배열과는 다른 타입이지만 fmt.Println로 출력을 하면 유사한 형태로 출력됩니다.
	fmt.Println("2d: ", twoD)
}
