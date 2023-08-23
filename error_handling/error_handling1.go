package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya Bulunamadı!", pErr.Path)
			return
		} else {
			fmt.Println("Dosya Bulunamadı!")
		}
	} else {
		fmt.Println("Dosya Adı: ", f.Name())
	}
}
