package errorhandling

import (
	"errors"
	"fmt"
)

func tahminet(tahmin int) (string, error) {

	akildakisayi := 56

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 Arasında Bir Sayı Giriniz!")
	}

	if tahmin > akildakisayi {
		return "Daha Küçük!", nil
	}

	if tahmin < akildakisayi {
		return "Daha Büyük!", nil
	}

	return "Bildiniz", nil
}

func Demo2() {
	mesaj1, hata1 := tahminet(150)
	fmt.Println(mesaj1)
	fmt.Println(hata1)

	/*_, hata2 := tahminet(-2)

	fmt.Println(hata2)
	_, hata3 := tahminet(101)

	fmt.Println(hata3)*/

}
