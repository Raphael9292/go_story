package main

import "fmt"

func main() {
	//빈 맵을 생성하기 위한 내장 함수 make 사용
	//make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	//len을 맵에 사용하면 key/value 쌍의 갯수를 반환
	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	/*
		맵에서 값을 꺼내오면서 반환되는 선택적인 두번째 반환값은 해당 키가 맵에 존재하는지에 대한 여부를 나타냅니다.
		이는 키값이 존재하지 않는건지 또는 해당 값이 0이나 ""과 같은 제로값 인지를 명확하게 구분하는데 사용할 수 있습니다.
		값 자체가 필요없는 경우엔, 공백 식별자(blank identifier) _를 사용해 무시할 수 있습니다.
	*/
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
