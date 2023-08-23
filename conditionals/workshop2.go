package conditionals

import "fmt"

func Workshop2() {
	akildaki_sayi := 69
	tahmin_edilen_sayi := 0
	denemeler := 1

	for tahmin_edilen_sayi != akildaki_sayi {
		fmt.Println("Sayı Tahmin Et: ")
		fmt.Scan(&tahmin_edilen_sayi)

		if tahmin_edilen_sayi > akildaki_sayi {
			fmt.Println("Çok çıktın, in biraz!")
			denemeler += 1

		} else if tahmin_edilen_sayi < akildaki_sayi {
			fmt.Println("Güven kendine, yüksel biraz!")
			denemeler += 1
		} else {
			fmt.Println("Doğru, Tebrikler!")
			fmt.Printf("%v Denemede Buldunuz ", denemeler)
			if denemeler >= 1 && denemeler <= 3 {
				fmt.Print("Süper\n")
			} else if denemeler >= 4 && denemeler <= 6 {
				fmt.Print("Güzel\n")
			} else {
				fmt.Print("Fena Değil\n")
			}
		}
	}
}
