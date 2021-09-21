package main

// 문자열은 +로 합칠 수 있습니다.
// https://pkg.go.dev/fmt@go1.17.1

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
