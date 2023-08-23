package pointers

import "fmt"

func Demo1(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("Demodaki Sayı: ", sayi)
}
