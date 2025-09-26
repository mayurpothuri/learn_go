package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%gC = %gF\n", FreezingC, CToF(FreezingC))
	fmt.Printf("%gC = %gF\n", BoilingC, CToF(BoilingC))
	fmt.Printf("%gF = %gC\n", CToF(FreezingC), FToC(CToF(FreezingC)))
	fmt.Printf("%gF = %gC\n", CToF(BoilingC), FToC(CToF(BoilingC)))
}
