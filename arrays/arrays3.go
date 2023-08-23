package arrays

import "fmt"

func Demo3() {
	numbers := [5]int{20, 30, 50, 10, 2}
	//fmt.Println(numbers)

	Biggest := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > Biggest {
			Biggest = numbers[i]
		}
	}
	fmt.Printf("En Büyük: %v", Biggest)
}
