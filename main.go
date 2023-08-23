package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	/*fmt.Println("selamlar") //Yazdırır ve alt satıra geçer.
	  fmt.Print("metisn")     //Sadece yazdırır

	  sayi3 := 46 //go ya özel tanımlama şekli var kullanmıyoruz
	  fmt.Println(sayi3)
	  fmt.Printf("Veri Tipi: %T", sayi3) //printf böyle formatlı şekilde yazdırmak için kullanılır
	  //                     %T de büyük harf olmalı, küçük t olmuyor*/

	/*//Değişken tanımlamalar

	  var metin string = "Merhabalar! :)" //String
	  fmt.Println(metin)
	  var sayi int = 15 //Int
	  fmt.Println(sayi)
	  var float_sayi float32 = 0.5 //float
	  fmt.Println(float_sayi)
	  var float_sayi2 float64 = 0.5 //float
	  fmt.Println(float_sayi2)*/
	/*var aktarım string = "selamlar"
	  aktarım = 15 //--> Burada da görüldüğü üzere bir değişken tanımlandıktan sonra türü değiştirilemez!
	  fmt.Println(aktarım)*/

	//Example

	/*var durum1 bool
	  var durum2 bool

	  var ad1 string = "Onur"
	  var ad2 string = "Saffet"

	  durum1 = ad1 == ad2
	  fmt.Println(durum1) //false
	  durum2 = ad1 != ad2
	  fmt.Println(durum2) //true*/

	//conditionals.Demo2()
	//conditionals.Demo3()
	//loops.Demo2()
	//loops.Demo3()
	//forremember.Demo1()
	//conditionals.Workshop4()
	//var result = functions.Topla(6, 7)
	//fmt.Println(result)
	//_, _, _, sonuc4 := functions.DortIslem(5, 2)
	/*fmt.Println("Toplam: ", sonuc1)
	  fmt.Println("Çıkarım: ", sonuc2)
	  fmt.Println("Çarpım: ", sonuc3)*/
	//fmt.Println("Bölüm: ", float64(sonuc4))

	/*var a float32 = 5
	  var b float32 = 2

	  fmt.Println(a / b)*/

	/*sonuc := functions.ToplaVariadic(1, 4, 5, 6, 3, 10)
	  fmt.Println(sonuc)*/

	//maps.Demo1()
	//range_test.Demo3()

	/*sayi := 20
	  pointers.Demo1(&sayi)
	  fmt.Println("Maindeki Sayı: ", sayi)*/

	/*sayilar := []int{1, 2, 3}
	  pointers.Demo2(sayilar)
	  fmt.Println("Masadaki Sayı:", sayilar[0])*/

	//structs.Demo2()
	/*ciftsayiToplamCn := make(chan int)
	  teksayiToplamCn := make(chan int)

	  go channels.CiftSayilar(ciftsayiToplamCn)
	  go channels.TekSayilar(teksayiToplamCn)

	  ciftsayiToplamsCn, teksayiToplamsCn := <-ciftsayiToplamCn, <-teksayiToplamCn

	  carpim := ciftsayiToplamsCn * teksayiToplamsCn
	  fmt.Println("Çarpım:", carpim)*/

	// interfaces.Demo2()
	// deferstatement.Demo3()
	// errorhandling.Demo1()
	// interfaces.Demo3()

	//fmt.Println(errorhandling.Tahminet2(990))

	//stringfunctions.Demo2()
	//estful.Demo2()

	project.AddProduct()
	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
}
