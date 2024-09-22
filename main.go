package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\x1b[33mUsage: go run . data.txt\033[0m")
		return
	}
	file, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	Data_Manipulation(string(file))
}

func Data_Manipulation(data string) {
	var numbers []float64
	tab_data := strings.Split(data, "\n")
	for _, r := range tab_data {
		if r == "" {
			continue
		}
		temp, err := strconv.ParseFloat(r, 64)
		if err != nil {
			fmt.Println(err)
		}
		numbers = append(numbers, float64(temp))
	}
	m, b := Linear_Regression(numbers)
	pearson := Pearson_Correlation(numbers)

	fmt.Printf("Linear Regression Line: y = %.fx + %.f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.f\n", pearson)
}

// function linear 
func Linear_Regression(numbers []float64) (float64, float64) {
	lenght := float64(len(numbers))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i, y := range numbers {
		x := float64(i)
		sumX = sumX + x
		sumY = sumY + y
		sumXY = sumXY + (x * y)
		sumX2 = sumX2 + (x * x)
		sumY2 = sumY2 + (y * y)
	}

	m := (lenght*sumXY - sumX*sumY) / (lenght*sumX2 - sumX*sumX)
	b := (sumY - m*sumX) / lenght

	return m, b
}

func Pearson_Correlation(data []float64) float64 {
	lenght := float64(len(data))
	var sumX, sumY, sumXY, sumX2, sumY2 float64

	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}

	numerator := (lenght*sumXY - sumX*sumY)
	denominator := math.Sqrt((lenght*sumX2 - sumX*sumX) * (lenght*sumY2 - sumY*sumY))

	if denominator == 0 {  //Eliminer la division par zero
		return 0   
	}
	return numerator / denominator
}
