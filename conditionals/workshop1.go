package conditionals

import "fmt"

func Demo3() {
	/*var sayi1 int = 50
	var sayi2 int = 40
	var sayi3 int = 30

	if sayi1 > sayi2 {
		if sayi1 > sayi3 {
			fmt.Println("Sayı 1 En Büyüğüdür!")
		} else {
			fmt.Println("Sayı 3 En Büyüğüdür!")
		}
	} else if sayi2 > sayi3 {
		fmt.Println(("Sayı 2 En Büyüğüdür!"))
	} else {
		fmt.Println("Sayı 3 En Büyüğüdür!")
	}*/

	var n1, n2, n3 int = 5, 15, 8

	var biggest int = n1

	if n2 > biggest {
		biggest = n2
	}
	if n3 > biggest {
		biggest = n3
	}

	fmt.Printf("En Büyük Sayı: %v", biggest)

}
