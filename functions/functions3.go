package functions

func ToplaVariadic(numbers ...int) int {
	toplam := 0
	for i := 0; i < len(numbers); i++ {
		toplam = toplam + numbers[i]
	}
	return toplam
}
