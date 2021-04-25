package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("negative value not support, %v", x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	// 1.414213562373095 10次 2*z修正结果
	// 1.414213562373095 100次 2*z修正结果
	// 1.4142040671509484 10次 1.5*z修正
	// 1.414213562373095 100次 1.5*z修正
	// 1.4142098061696458 10次 3*z修正结果
	// 1.414213562373095 100次 3*z修正结果
	// 1.414213562373095048 计算器结果
}
