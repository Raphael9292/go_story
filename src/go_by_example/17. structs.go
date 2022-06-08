package main

import "fmt"

/*
Go의 structs는 필드들로 이루어진 타입을 갖는 컬렉션입니다.
레코드를 구성하기 위해 데이터들을 그룹핑 하는데 유용합니다.
*/

//이 person 구조체 타입은 name과 age 필드를 갖습니다.
type person struct {
	name string
	age  int
}

func main() {
	//새로운 구조체를 생성한다
	fmt.Println(person{"Bob", 10})

	//필드명을 사용해 구조체를 초기화 할 수 있다
	fmt.Println(person{name: "Alice", age: 20})

	// 생략된 필드는 0값을 갖는다
	fmt.Println(person{name: "Fred"})

	//앞에 &를 붙이면 구조체 포인터를 반환한다
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "raphael", age: 30}
	// .을 이용해서 구조체 필드에 접근한다
	fmt.Println(s.name)

	sp := &s
	//구조체 포인터에서도 .을 사용할 수 있다. 포인터는 자동으로 역참조된다
	fmt.Println(sp.age)

	sp.age = 40
	//구조체는 mutable하다
	fmt.Println(sp.age)
}
