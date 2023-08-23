package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)
	sozluk["Table"] = "Masa"
	sozluk["Book"] = "Kitap"
	sozluk["Pencil"] = "Kalem"

	fmt.Println(sozluk["Book"])
	fmt.Println("Eleman Sayısı :", len(sozluk))
	fmt.Println(sozluk)
	delete(sozluk, "Book")
	fmt.Println("Eleman Sayısı :", len(sozluk))
	fmt.Println(sozluk)

	surec, varmi := sozluk["Table"]
	fmt.Println(surec)
	fmt.Println(varmi)
}
