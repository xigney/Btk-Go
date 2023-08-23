package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var istenen float64 = 900

	if istenen > hesap {
		fmt.Println("Hesaptaki Para Yetersiz!")
	} else if istenen == hesap {
		fmt.Println("Paranız hazırlanıyor")
	} else {
		fmt.Println("Paranız hazırlanıyor")
	}

	
}
