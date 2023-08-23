package deferstatement

import "fmt"

func Demo2(sayi int) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		return "Çift Sayıdır"
	}
	if sayi != 2 {
		return "Tek Sayıdır"
	}
	return "Belli Değil"
}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)

}

func DeferFunc() {
	fmt.Println("Bitti")
}
