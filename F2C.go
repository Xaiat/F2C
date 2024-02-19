package main

import (
	"fmt"
)

/*
// Method 1
// CelsiusToFahrenheit Transforms Celsius to Fahrenheit
// celsiusToFahrenheit 将摄氏度转换为华氏度
func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	// Define the start and end Celsius
	// 定义起始和结束的摄氏度
	startCelsius, endCelsius := 40, -10
	// Calculate the start and end Fahrenheit
	// 计算起始和结束的华氏度
	startFahrenheit := celsiusToFahrenheit(float64(startCelsius))
	endFahrenheit := celsiusToFahrenheit(float64(endCelsius))

	fmt.Println("华氏度\t摄氏度")
	// Traverse from the Fahrenheit corresponding to 40°C to the Fahrenheit corresponding to -10°C in reverse order
	// 从40°C对应的华氏度到-10°C对应的华氏度，逆序遍历
	for F := startFahrenheit; F >= endFahrenheit; F-- {
		// Convert Fahrenheit back to Celsius for display
		// 将华氏度转换回摄氏度进行显示
		C := (F - 32) * 5 / 9
		fmt.Printf("%.0f°F\t%.2f°C\n", F, C)
	}
}
*/

// Method 2
// fahrenheitToCelsius Transforms Fahrenheit to Celsius
// fahrenheitToCelsius 将华氏度转换为摄氏度
func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	// Define the start and end Celsius
	// 定义起始和结束的摄氏度
	startCelsius, endCelsius := 40.0, -10.0
	// Calculate the start and end Fahrenheit
	// 计算起始和结束的华氏度
	startFahrenheit := startCelsius*9/5 + 32
	endFahrenheit := endCelsius*9/5 + 32

	fmt.Println("华氏度\t\t摄氏度")
	fmt.Println("Fahrenheit\tCelsius")

	// Traverse from the Fahrenheit corresponding to 40°C to the Fahrenheit corresponding to -10°C in reverse order
	// 从40°C对应的华氏度开始逆序遍历到-10°C对应的华氏度
	for F := startFahrenheit; F >= endFahrenheit; F-- {
		// Convert Fahrenheit back to Celsius for display
		// 使用提供的函数将华氏度转换回摄氏度进行显示
		C := fahrenheitToCelsius(F)
		fmt.Printf("%.0f°F\t%13.2f°C\n", F, C)
	}
}
