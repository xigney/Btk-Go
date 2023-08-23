package errorhandling

import "fmt"

type borderException struct {
	message   string
	parameter int
}

func (b borderException) Error() string {
	return fmt.Sprintf("%d--->%s", b.parameter, b.message)

}

func Tahminet2(tahmins int) (string, error) {
	if tahmins < 1 || tahmins > 100 {
		return "", &borderException{"Sınırların Dışında Kaldı!", tahmins}
	}
	return "Bildiniz!", nil
}
