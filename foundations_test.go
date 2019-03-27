package gosignalprocessing

import (
	"testing"
)

var SDTestData = createSDTestData()

func TestMean(t *testing.T) {
	values := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}
	m, err := Mean(values)
	if err != nil {
		t.Errorf("mean returned an unexpected error: %s\n", err)
	}
	if m != 5.5 {
		t.Errorf("mean failed, returned %f instead of 5.5", m)
	}

	values = []float64{}
	_, err = Mean(values)
	if err == nil {
		t.Error("mean should return an error for zero length arrays")
	}
}

func createSDTestData() []float64 {
	var tst [512]float64
	for i := 0; i < 512; i = i + 2 {
		tst[i] = 2.0
		tst[i+1] = -2.0
	}
	return tst[:]
}

func TestStandardDeviation(t *testing.T) {
	mn, sd, err := StandardDeviation(SDTestData)
	if err != nil {
		t.Errorf("standardDeviation returned an unexpected error: %s\n", err)
	}

	if mn != 0 {
		t.Errorf("standardDeviation returned a mean value of %f instead of 0", mn)
	}

	if sd != 2 {
		t.Errorf("standardDeviation returned a standard deviation value of %f instead of 2", sd)
	}
}

func TestStandardDeviationRunning(t *testing.T) {
	mn, sd, err := StandardDeviationRunning(SDTestData)
	if err != nil {
		t.Errorf("StandardDeviationRunning returned an unexpected error: %s\n", err)
	}

	if mn != 0 {
		t.Errorf("StandardDeviationRunning returned a mean value of %f instead of 0", mn)
	}

	if sd != 2 {
		t.Errorf("StandardDeviationRunning returned a standard deviation value of %f instead of 2", sd)
	}
}

func BenchmarkStandardDeviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardDeviation(SDTestData)
	}
}

func BenchmarkStandardDeviation2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardDeviation2(SDTestData)
	}
}

func BenchmarkStandardDeviationRunning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardDeviationRunning(SDTestData)
	}
}
