package channels

func CiftSayilar(ciftSayicn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam += i
	}

	ciftSayicn <- toplam
}

func TekSayilar(TekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam += i
	}

	TekSayiCn <- toplam

}
