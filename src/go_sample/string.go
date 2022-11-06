package main

/*
문자 하나를 표현하는 데 rune 타입을 사용합니다. UTF-8은 한 글자가 1 ~3바이트 크기이기 때문
에 UTF-8 문자값을 가지려면 3바이트가 필요합니다. 하지만 Go 언어 기본 타입에서 3바이트 정
수타입은 제공되지 않기 때문에 rune 타입은 4바이트 정수 타입인 int32 타입의 별칭 타입입니다.
*/

import "fmt"

func main() {
	/*
		문자 한 개는 작은따옴표로 묶어서 표시합니다
		문자 한개를 표현하는데 rune type을 사용합니다.
	*/
	var char rune = '한'

	fmt.Printf("%T\n", char) // char type
	fmt.Println(char)        // char value
	fmt.Printf("%c\n", char) // print string

	/*
		영문 문자열 크기는 5입니다. UTF-8에서 한글은 글자당 3바이트를 차지하기 때문에 총 3 x
		5를 하면 15바이트가 나오게 됩니다. 또 UTF-8에서 영문자는 글자당 1바이트라서 총 5바이트
		가 됨니다.
	*/
	str1 := "가나다라마바사"
	str2 := "abcdef"

	fmt.Printf("len(str1) = %d\n\n", len(str1))
	fmt.Printf("len(str2) = %d\n\n", len(str2))

	str := "Hello World"
	// ‘H’, ‘e’, ‘l'，‘l'，‘o'， '', ‘W', ‘o', ‘r', ‘l', ‘d',   문자코드 배열
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	fmt.Println(str)
	fmt.Println(string(runes))

	/*
		string 타입과 []byte 타입은 상호 타입 변환이 가능합니다. []byte는 byte 즉 1바이트 부호 없
		는 정수 타입의 가변 길이 배열입니다. 문자열이란 것도 결국 메모리에 있는 데이터이고, 메모리
		는 1바이트 단위로 저장되기 때문에 모든 문자열은 1바이트 배열로 변환 가능합니다. 파일을 쓰
		거나 네트워크로 데이터를 전송하는 경우 io.Writer 인터페이스를 사용하고 io.Writer 인터페
		이스는 []byte 타입을 인수로 받기 때문에 []byte 타입으로 변환해야 합니다. 그래서 문자열을 쉼
		게 전송하고자 string에서 []byte 타입으로 변환을 지원합니다. 이 부분에 대해서는 20장 ‘인터
		페이스’와 A.3절 ‘입출력 처리’에서 다룹니다. 지금은 string과 []byte 타입 간 상호 변환이 가능
		하다는 것만 아시고 넘어가도 됩니다.
	*/

	// 인덱스를 사용한 바이트 단위 순회하기 (한글은 깨졌다)
	// str[i]처럼 인덱스로 접근하면 요소의 타입은 uint8 즉 byte입니다. 그래서 1바이트 크기인 영문자는 잘 표시되는데 3바이트 크기인 한글은 깨져 표시된 겁니다.
	str = "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}
	fmt.Printf("\n")

	// []rune type으로 변환 후 한글자씩 순회
	/*
		이렇게 []rune으로 변환한 다음에 순회하면 한 글자씩 순회할 수 있지만 []rune으로 변환되는 과정에서 별도의 배열을 할당하므로 불필요한 메모리를 사용하게 됩니다.
		range 키워드를 사용해 순회하면 이를 방지할 수 있습니다.
	*/
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}
	fmt.Printf("\n")

	//	range keyword를 이용한 한글자씩 순회
	/*
		모든 문자 타입이 int32, 즉 rune입니다. rune은 기본적으로 숫자값이기 때문에 어떤 수인지 줄력하고, %c를 이용해 해당 문자를 줄력했습니다.
		이처럼 range를 이용하면 추가 메모리 할당 없이 문자열을 한 글자씩 순회할 수 있어서 불필요한 메모리 낭비를 없앨 수 있습니다.
	*/
	for _, v := range str {
		fmt.Printf("타입:%T 값:%d 문자:%c\n", v, v, v)
	}

}
