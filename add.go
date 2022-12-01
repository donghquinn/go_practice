package main

import (
	"fmt"
)

// 총합 함수
func sum(first_number int, second_number int) int {

	return first_number + second_number

}

// defer는 먼저 함수의 최종 리턴을 한 후 그 과정을 역순으로 보여준다.
func deferTest() {
	total := 0

	for i := 0; i < 5; i += 1 {
		defer fmt.Println("Counts: %d", i)

		total += i

		defer fmt.Println("Total Count: ", total)
	}

	fmt.Println("Final Total Count Ended: ", total)

	return
}

func main() {
	const a = 3
	const b = 3

	// 짧은 대입문. var 변수를 할당한다.
	global_function := 1324

	// Println 에 %d 를 입력하면 그대로 출력된다.
	fmt.Println("The Sum Value: %d", sum(a, b))
	// The Sum Value: %d 6 이 출력된다.

	// 정상출력
	fmt.Println("The Sum Value:", sum(a, b))
	fmt.Printf("%d", global_function)

	fmt.Println("Defer Test Start")
	deferTest()
}
