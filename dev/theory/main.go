package main

import (
	"fmt"

	"github.com/minkj1992/go_nomad/theory/collection"
	"github.com/minkj1992/go_nomad/theory/hello"
)

// 리터럴(축약형)은 func 밖에서는 사용 불가능하다.
// name := "minwook"
func main() {
	// 패키지의 func은 대문자로 시작
	name := "minwook"
	hello.SayHello(name)
	hello.SayBye(name)

	total := collection.Add(1, 2, 3, 4, 5)
	fmt.Println(total)

	oldSlice := []string{"leoo.j", "minkj1992", "jmu2001", "jejulover"}
	newSlice := collection.NewSlice(oldSlice, "kakao", "naver", "google")
	fmt.Println(&oldSlice[0])
	fmt.Println(&newSlice[0])
}
