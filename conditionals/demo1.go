package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var istenen float64 = 900

	if istenen > hesap {
		fmt.Println("Hesaptaki Para Yetersiz!")
	}
	if istenen <= hesap {
		fmt.Println("Paranız hazırlanıyor")
	}
}
