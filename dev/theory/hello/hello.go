package hello

import "fmt"

// import 할 func은 대문자로 시작 (public)
func SayHello(name string) {
	// defer를 사용한다면, func 종료시점에 동작할 action 지정 가능
	defer sayBye(name)
	fmt.Println("Hello " + name)
}

func SayBye(name string) {
	fmt.Println("Public Bye " + name)
}

// private
func sayBye(name string) {
	fmt.Println("private Bye " + name)
}
