package library

import "math"

type Hitung interface {
	Luas() float64
	Keliling() float64
}

type Lingkaran struct {
	Diameter float64
}

func (l Lingkaran) JariJari() float64 {
	return l.Diameter / 2
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.JariJari(), 2)
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}

type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return math.Pow(p.Sisi, 2)
}

func (p Persegi) Keliling() float64 {
	return p.Sisi * 4
}
