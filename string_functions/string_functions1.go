package stringfunctions

import (
	"fmt"
	s "strings" //alias
)

// case sensitive
// ascii
func Demo1() {
	isim := "Engin"
	fmt.Println(s.Count(isim, "E"))
	fmt.Println(s.Contains(isim, "a"))

}
