package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type Kubus struct {
	sisi float64
}

func (k *Kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *Kubus) luas() float64{
	return math.Pow(k.sisi, 4) * 6
}

func (k *Kubus) keliling() float64{
	return k.sisi * 12
}

func main(){
	var bangunRuang hitung = &Kubus{4}

	fmt.Println("============== Kubus")
	fmt.Println("Luas \t:",bangunRuang.luas())
	fmt.Println("Keliling \t:",bangunRuang.keliling())
	fmt.Println("Volume \t:", bangunRuang.volume())
}