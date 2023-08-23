package functions

func DortIslem(n1 int, n2 int) (int, int, int, float64) {
	toplam := n1 + n2
	fark := n1 - n2
	carpim := n1 * n2
	bolum := n1 / n2

	return toplam, fark, carpim, float64(bolum)

}
