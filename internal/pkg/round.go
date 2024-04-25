package pkg

import "math"

// Функция для округления числа до указанного количества десятичных знаков
func Round(num float64, decimals int) float64 {
	// Вычисление множителя для заданного количества десятичных знаков
	multiplier := math.Pow(10, float64(decimals))

	// Округление числа до ближайшего целого
	rounded := math.Round(num * multiplier)

	// Деление на множитель для возвращения округленного значения с указанным количеством десятичных знаков
	result := rounded / multiplier

	return result
}
