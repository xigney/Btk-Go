package range_test

import "fmt"

func Demo1() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for i, city := range cities {
		fmt.Print(i)
		fmt.Println(city)
	}
}
