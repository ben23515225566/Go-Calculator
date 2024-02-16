package calculate

import (
	"math"
)

func Add(num1, num2 float64) float64 {
	return num1 + num2
}

func Sub(num1, num2 float64) float64 {
	return num1 - num2
}

func Mul(num1, num2 float64) float64 {
	return num1 * num2
}

func Div(num1, num2 float64) float64 {
	return num1 / num2
}

func Pow(num1, num2 float64) float64 {
	return math.Pow(num1, num2)
}

func Sqrt(num1 float64) float64 {
	return math.Sqrt(num1)
}

func Mod(num1, num2 float64) float64 {
	return math.Mod(num1, num2)
}
