package conditionals

import "fmt"

func Workshop3() {

	var alinan_sayi int

	fmt.Println("Lütfen Bir Sayı Giriniz: ")
	fmt.Scan(&alinan_sayi)

	if alinan_sayi <= 0 {
		fmt.Println("Asallık Aranmaz!")
	} else if alinan_sayi == 1 {
		fmt.Println("Asal Değildir!")
	} else {
		var i int = 1
		for alinan_sayi != i {
			if i == 1 {
				i = i + 1
				continue
			} else {
				if alinan_sayi%i == 0 {
					fmt.Println("Asal Değildir")
					goto end
				} else {
					i = i + 1
					continue
				}
			}

		}
		fmt.Println("Asaldır!")
	end:
	}
}
