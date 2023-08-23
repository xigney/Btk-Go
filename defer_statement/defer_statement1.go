package deferstatement

import "fmt"

func A() {
	fmt.Println("A Fonksiyonu Çalıştı")
}

func C() {
	fmt.Println("C Fonksiyonu Çalıştı")
}

func D() {
	fmt.Println("D Fonksiyonu Çalıştı")
}

func B() {
	defer A()
	defer C()
	defer D()

	fmt.Println("B Fonksiyonu Çalıştı")

}
