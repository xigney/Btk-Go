package slicess

import "fmt"

func Demo1() {
	names := make([]string, 3)

	fmt.Println(names)
	for i := 0; i < len(names); i++ {
		fmt.Println("Ad Gir: ")
		fmt.Scan(&names[i])
	}

	//names[3] = "Onur"

	fmt.Println(len(names))

	names = append(names, "Onur")

	fmt.Println(names)
	fmt.Println(len(names))

}
