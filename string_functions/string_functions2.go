package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Engin"

	fmt.Println(s.HasPrefix(isim, "Eng"))
	fmt.Println(s.HasSuffix(isim, "in"))
	fmt.Println(s.Index(isim, "gi"))
	harfler := []string{"e", "n", "g", "i", "n"}
	fmt.Println(s.Join(harfler, ""))
}
