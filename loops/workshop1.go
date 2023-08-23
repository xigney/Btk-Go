package loops

import "fmt"

func Demo3() {
	tutulansayi := 76
	tahminedilensayi := 0

	i := 0

	for i == 0 {
		fmt.Print("Bir Sayı Tahmin Et: ")
		fmt.Scan(&tahminedilensayi)
		if tahminedilensayi < tutulansayi {
			fmt.Println("Daha Yüksek!")
		} else if tahminedilensayi > tutulansayi {
			fmt.Println("Daha Düşük!")
		} else if tahminedilensayi == tutulansayi {
			fmt.Println("Doğru Sayı, Oyun Bitti!")
			i = 1
		}
	}

}
