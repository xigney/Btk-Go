package arrays

import (
	"fmt"
)

func Demo4() {
	var numbers [2][3]int

	//Take a number
	for i := 0; i < 2; i++ {
		{
			for j := 0; j < 3; j++ {
				fmt.Println("Sayı Giriniz: ")
				fmt.Scan(&numbers[i][j])
			}
		}
	}
	fmt.Print("\n")
	//Write numbers
	for i := 0; i < 2; i++ {
		{
			for j := 0; j < 3; j++ {
				fmt.Println(numbers[i][j])
			}
		}
	}
}
