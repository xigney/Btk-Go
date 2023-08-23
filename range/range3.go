package range_test

import "fmt"

func Demo3() {
	sozluk := map[string]string{"book": "kitap", "table": "masa"}
	for k, v := range sozluk {
		fmt.Println(k)
		fmt.Println(v)
	}
}
