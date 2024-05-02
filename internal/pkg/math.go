package pkg

import (
	"math"
	"sort"
)

// Функция для округления числа до указанного количества десятичных знаков
func Round(num float64) float64 {
	// Вычисление множителя для заданного количества десятичных знаков
	multiplier := math.Pow(10, float64(2))

	// Округление числа до ближайшего целого
	rounded := math.Round(num * multiplier)

	// Деление на множитель для возвращения округленного значения с указанным количеством десятичных знаков
	result := rounded / multiplier

	return result
}

func Median(numbers []int) float64 {
	// Сортируем срез чисел
	sort.Ints(numbers)

	length := len(numbers)
	if length == 0 {
		panic("Empty slice")
	}

	// Если количество элементов нечетное, возвращаем средний элемент
	if length%2 != 0 {
		return float64(numbers[length/2])
	}

	// Если количество элементов четное, возвращаем среднее значение двух средних элементов
	mid := length / 2
	return float64(numbers[mid-1]+numbers[mid]) / 2.0
}
