package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeDataPointer(arg *Data) { // 매개변수로 Data pointer를 받는다
	arg.value = 999 // arg 데이터를 변경한다
	//(*arg).value = 999 와 같은말
	arg.data[100] = 999
}

func ChangeData(arg Data) { // 매개변수로 Data를 받는다
	arg.value = 777 // arg 데이터를 변경한다
	arg.data[100] = 777
}

func main() {
	var data Data

	ChangeDataPointer(&data) // 인수로 data 주소를 넘긴다.
	fmt.Printf("value = &d\n", data.value)
	fmt.Printf("data[100] = &d\n", data.data[100])

	ChangeData(data) // 인수로 data를 넘긴다.
	fmt.Printf("value = &d\n", data.value)
	fmt.Printf("data[100] = &d\n", data.data[100])

}

/*
ChangeData는 data 변수를 인수로 넣는다.
data 변수가 모두 복사되었기에 매개변수 arg, data는 서로 다른 메모리 공간을 갖는 변수이다.
그렇기에 777로 선언해도 적용이 되지 않는거고..
*/
