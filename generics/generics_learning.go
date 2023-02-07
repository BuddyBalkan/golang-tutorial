package main

import (
	"fmt"
)

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"1st" : 34,
		"2nd" : 12,
		"3rd" : 58,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"1st" : 35.98,
		"2nd" : 26.99,
		"3rd" : 47.03,
	}

	fmt.Printf("Non-Generic Sums: %v and %v \n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v \n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}
// SumInts adds together the values of m.
// 整数累加 map的集合
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m{
		s += v
	}
	return s
}

// Sumloats adds together the values of m.
// 小数累加 map的集合
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m{
		s += v
	}
	return s
}


// SumIntsOfFloats sums the values of map m.
// It supports both int64 and float64 as types for map values.
// comparable -- map的key值必须为comparable的；任何可以用运算符“==”和“！=”操作的类型
// 由于已经定义了K和V的泛型 故 m变量是可以使用的
func SumIntsOrFloats[K comparable, V int64 | float64] (m map[K]V) V {
	var s V
	for _, v := range m{
		s += v
	}
	return s
}
