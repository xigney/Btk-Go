package range_test

import (
	"fmt"
)

func Demo2() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	for _, sayi := range numbers {
		if sayi%2 != 0 {
			toplam = toplam + sayi
		}
	}
	fmt.Println("Toplam: ", toplam)
}
