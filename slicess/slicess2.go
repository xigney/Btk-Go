package slicess

import "fmt"

func Demo2() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}
	fmt.Println(cities)
	copy_cities := make([]string, len(cities))
	fmt.Println(copy_cities)
	copy(copy_cities, cities)
	fmt.Println(copy_cities)
	cities = append(cities, "Adana")
	fmt.Println(cities)
	fmt.Println(copy_cities)

}
