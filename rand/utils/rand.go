package utils

import "github.com/Pallinder/go-randomdata"

func GenerateRandomIntData(min, max, count int) []int {
	var arr []int
	for i := 0; i < count; i++ {
		arr = append(arr, randomdata.Number(min, max))
	}
	return arr
}
func GenerateRandomDecimalData(min, max, count, precision int) []float64 {
	var arr []float64
	for i := 0; i < count; i++ {
		arr = append(arr, randomdata.Decimal(min, max, precision))
	}
	return arr
}
