package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func Data_Manipulation(data string) {
	var numbers []float64
	tab_data := strings.Split(data, "\n")
	for _, r := range tab_data {
		if r == "" {
			continue
		}

		// Convertir string en flaot64
		temp, err := strconv.ParseFloat(r, 64)
		if err != nil {
			fmt.Println(err)
		}
		numbers = append(numbers, float64(temp))
	}
	a, b, pearson_correlation := Linear_Regression(numbers)

	fmt.Printf("\x1b[33mLinear Regression Line: y = %fx + %f\n\033[0m", a, b)
	fmt.Printf("\x1b[33mPearson Correlation Coefficient: %.10f\n\033[0m", pearson_correlation)
}
