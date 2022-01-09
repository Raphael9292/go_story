package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//동일한 case문에서 여러개의 표현식을 구분하기 위해 콤마(,)를 사용할 수 있습니다.
	//이 예시에서 우리는 또한 default 케이스도 사용했습니다. default는 선택사항입니다.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It`s weekends")
	default:
		fmt.Println("It`s a weekdays")
	}

	//표현식이 없는 switch는 if/else 로직을 표현하기 위한 또 다른 방법입니다.
	//여기서 우리는 또한 상수가 아닌 case문을 사용하는 방법을 볼 수 있습니다.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It`s before noon")
	default:
		fmt.Println("It`s afternoon")
	}

	fmt.Println("----------")

	//타입 switch는 값 대신 타입을 비교합니다.
	//여러분은 인터페이스 값의 타입을 알아내기 위해 사용할 수 있습니다.
	//이 예시에서, 변수 t는 해당 절에 해당하는 타입을 가질 것입니다.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I`m a bool")
		case int:
			fmt.Println("I`m an int")
		default:
			fmt.Println("Don`t know type %T|n", t)
		}
	}

	whatAmI(true)
	whatAmI(123)
	whatAmI("test")
}
