package generator

import (
	"fmt"
	"math"
)

type Biome int

const (
	Rift Biome = iota
	Plain
	Mountain
)

type DSField struct {
	fieldSize int
	seed      int
	salt      float64
	field     [][]float64
	shifts    []float64
	biome     Biome
}

func newDSField(seed string, fieldSize int, salt float64) (result *DSField) {
	_seed := 0
	for _, ch := range seed {
		_seed += int(ch)
	}

	if fieldSize%2 == 0 {
		fieldSize++
	}

	result.field = make([][]float64, fieldSize)
	for i := range result.field {
		result.field[i] = make([]float64, fieldSize)
	}

	result = &DSField{seed: _seed, fieldSize: fieldSize, salt: salt}
	result.fillShiftMas()

	result.field[0][0] = result.shifts[0]
	result.field[0][fieldSize-1] = result.shifts[1]
	result.field[fieldSize-1][0] = result.shifts[2]
	result.field[fieldSize-1][fieldSize-1] = result.shifts[3]

	dsLength := result.fieldSize - 1
	result.diamondSqure(dsLength/2, dsLength/2, dsLength, 1)

	return
}

func (f *DSField) fillShiftMas() {
	f.shifts = make([]float64, int(math.Pow(float64(f.fieldSize), 2.0)))
	f.shifts[0] = float64(f.seed) / f.salt
	f.shifts[1] = f.shifts[0] + f.salt
	for i := 2; i < len(f.shifts); i++ {
		if i%2 == 0 {
			f.shifts[i] = 0.4*f.shifts[i-1] - 0.6*f.shifts[i-2] + 1.4*f.salt
		} else {
			f.shifts[i] = 0.32*f.shifts[i-1] - 0.87*f.shifts[i-2] - 0.95*f.salt
		}
	}
}

func (f *DSField) diamondSqure(x, y, length, iter int) {
	if length > 1 {
		localShiftIter := (iter - 1) * 5
		lls := length / 2 //aka local length shift
		// diamond step
		f.field[x][y] = (f.field[x-lls][y-lls]+f.field[x-lls][y+lls]+f.field[x+lls][y-lls]+f.field[x+lls][y+lls])/4.0 + f.shifts[localShiftIter]

		// square step
		// x, y + lls; x, y - lls; x + lls, y; x - lls, y

		// edge case handling
		// top edge
		if x-lls == 0 {
			f.field[x-lls][y] = (f.field[x][y]+
				f.field[x-lls][y-lls]+
				f.field[x-lls][y+lls])/3.0 + f.shifts[localShiftIter+1]
		} else {
			f.field[x-lls][y] = (f.field[x][y]+
				f.field[x-lls][y-lls]+
				f.field[x-lls][y+lls]+
				f.field[x-2*lls][y])/4.0 + f.shifts[localShiftIter+1]
		}
		// bot edge
		if x+lls == f.fieldSize-1 {
			f.field[x+lls][y] = (f.field[x][y]+
				f.field[x+lls][y-lls]+
				f.field[x+lls][y+lls])/3.0 + f.shifts[localShiftIter+2]
		} else {
			f.field[x+lls][y] = (f.field[x][y]+
				f.field[x+lls][y-lls]+
				f.field[x+lls][y+lls]+
				f.field[x+2*lls][y])/4.0 + f.shifts[localShiftIter+2]
		}
		// left edge
		if y-lls == 0 {
			f.field[x][y-lls] = (f.field[x][y]+
				f.field[x-lls][y-lls]+
				f.field[x+lls][y-lls])/3.0 + f.shifts[localShiftIter+3]
		} else {
			f.field[x][y-lls] = (f.field[x][y]+
				f.field[x-lls][y-lls]+
				f.field[x+lls][y-lls]+
				f.field[x][y-2*lls])/4.0 + f.shifts[localShiftIter+3]
		}
		// right edge
		if y+lls == f.fieldSize-1 {
			f.field[x][y+lls] = (f.field[x][y]+
				f.field[x-lls][y+lls]+
				f.field[x+lls][y+lls])/3.0 + f.shifts[localShiftIter+4]
		} else {
			f.field[x][y+lls] = (f.field[x][y]+
				f.field[x-lls][y-lls]+
				f.field[x+lls][y-lls]+
				f.field[x][y+2*lls])/4.0 + f.shifts[localShiftIter+4]
		}

		f.diamondSqure(x-lls/2, y-lls/2, lls, iter+1)
		f.diamondSqure(x+lls/2, y-lls/2, lls, iter+2)
		f.diamondSqure(x-lls/2, y+lls/2, lls, iter+3)
		f.diamondSqure(x+lls/2, y+lls/2, lls, iter+4)
	}
}

func (f *DSField) concreteBiome() {
	switch f.biome {
	case Plain:
		max := MaxElem(f.field)
		LambdaMas(f.field, func(x float64) float64 {
			return x / max
		})
	}
}

func (f *DSField) makeRift() {

}

func (f *DSField) makePlain()

func (f *DSField) PrintField() {
	for _, mv := range f.field {
		for _, v := range mv {
			fmt.Printf("|  %2f  ", v)
		}
		fmt.Print("|\n")
	}
}
