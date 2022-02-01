package main

import (
	"fmt"
	. "learn_go/library"
)

func main() {

	var bd Hitung

	bd = Persegi{Sisi: 5}

	fmt.Println(bd.Keliling())
	fmt.Println(bd.Luas())

	bd = Lingkaran{Diameter: 82}
	fmt.Println(bd.Keliling())
	fmt.Println(bd.Luas())
	fmt.Println(bd.(Lingkaran).JariJari())

}
