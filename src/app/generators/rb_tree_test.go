package generator

import "math/rand"

type UniqueRand struct {
	generated map[int]bool
}

func (u UniqueRand) Intn(n int) int {
	for {
		i := rand.Intn(n)
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}
