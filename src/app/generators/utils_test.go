package generator_test

import "testing"

func makeEinMatrix() [][]float64 {
	mas := make([][]float64, 5)
	for i := range mas {
		mas[i] = make([]float64, 5)
		for j := range mas[i] {
			mas[i][j] = 1
		}
	}
	return mas
}
func TestLambdaMas(t *testing.T) {
	mas := makeEinMatrix()
	LambdaMas
}
