package generator

type Less func(a, b interface{}) bool

func MaxElem(mas [][]float64) (result float64) {
	result = mas[0][0]
	for _, i := range mas {
		for _, j := range i {
			if j > result {
				result = j
			}
		}
	}
	return
}

func LambdaMas(mas [][]float64, function func(float64) float64) {
	for _, i := range mas {
		for _, j := range i {
			j = function(j)
		}
	}
}
