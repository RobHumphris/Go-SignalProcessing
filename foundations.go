package gosignalprocessing

import (
	"fmt"
	"math"
)

/*func meanInt(values []int) (int, error) {
	sum := 0
	l := len(values)
	if l == 0 {
		return sum, fmt.Errorf("divide by zero")
	}

	for _, v := range values {
		sum = sum + v
	}
	return sum / l, nil
}*/

func Mean(values []float64) (float64, error) {
	sum := 0.0
	l := len(values)
	if l == 0 {
		return sum, fmt.Errorf("divide by zero")
	}

	for _, v := range values {
		sum = sum + v
	}
	return (sum / float64(l)), nil
}

func Mean2(values []float64) (float64, error) {
	sum := 0.0
	l := len(values)
	if l == 0 {
		return sum, fmt.Errorf("divide by zero")
	}

	for i := 0; i < l; i++ {
		sum = sum + values[i]
	}
	return (sum / float64(l)), nil
}

func StandardDeviation(values []float64) (float64, float64, error) {
	mean, err := Mean(values)
	variance := 0.0
	if err != nil {
		return mean, variance, err
	}

	for _, v := range values {
		variance = variance + math.Pow((v-mean), 2)
	}
	variance = variance / float64(len(values))
	sd := math.Sqrt(variance)
	return mean, sd, nil
}

func StandardDeviation2(values []float64) (float64, float64, error) {
	mean, err := Mean2(values)
	variance := 0.0
	if err != nil {
		return mean, variance, err
	}

	l := len(values)

	for i := 0; i < l; i++ {
		variance = variance + math.Pow((values[i]-mean), 2)
	}
	variance = variance / float64(len(values))
	sd := math.Sqrt(variance)
	return mean, sd, nil
}

func StandardDeviationRunning(values []float64) (float64, float64, error) {
	n := 0.0
	sum := 0.0
	sumSquares := 0.0

	l := len(values)

	for i := 0; i < l; i++ {
		n = n + 1.0
		sum = sum + values[i]
		sumSquares = sumSquares + math.Pow(values[i], 2)
	}
	mean := sum / n
	variance := (sumSquares - math.Pow(sum, 2)/2) / (n - 1.0)
	sd := math.Sqrt(variance)
	return mean, sd, nil
}
