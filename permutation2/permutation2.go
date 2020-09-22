package permutation2

import "fmt"

type Ints struct {
	S []int
}

func NewInts(length int) *Ints {
	ints := &Ints{S: make([]int, length)}
	return ints
}

func Permutation(length int) {
	if length <= 1 {
		fmt.Println(length)
	} else {
		ints := NewInts(length)
		ints.Permute(length)
	}
}

func (ints *Ints) Permute(index int) {
	if index == 0 {
		fmt.Println(ints.S)
	} else {
		for i := 0; i < len(ints.S); i++ {
			if ints.S[i] == 0 {
				ints.S[i] = index
				ints.Permute(index - 1)
				ints.S[i] = 0
			}
		}
	}
}
