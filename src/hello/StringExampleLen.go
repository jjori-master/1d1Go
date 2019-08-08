package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 = "한글"
	var s2 = "Hello"

	fmt.Println(len(s1)) // 6자리로 나옴
	fmt.Println(len(s2)) // 5자리로 나옴

	// 한글, 한자, 일보어 등 UTF-8로 저장했을 때 2바이트가 넘는 문자열의 길이를 구하려면
	// 다음과 같이 unicode/utf8 패키지의 RuneCountInString 함수를 사용
	fmt.Println(utf8.RuneCountInString(s1))
}
