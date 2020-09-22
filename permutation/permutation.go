package permutation

import "fmt"

type Ints struct {
	S []int
}

func NewInts(length int) *Ints {
	ints := &Ints{S: make([]int, 0, length)}
	for i := 1; i <= length; i++ {
		ints.S = append(ints.S, i)
	}
	return ints
}

func Permutation(length int) {
	if length <= 1 {
		fmt.Println(length)
	} else {
		ints := NewInts(length)
		ints.Permute(0)
	}
}

func (ints *Ints) Permute(index int) {
	n := len(ints.S)
	if index >= n {
		fmt.Println(ints.S)
	} else {
		for i := index; i < n; i++ {
			ints.S[i], ints.S[index] = ints.S[index], ints.S[i]
			fmt.Println("before",index,i)
			fmt.Println(ints.S)
			ints.Permute(i + 1)
			fmt.Println("after",index,i)
			fmt.Println(ints.S)
			ints.S[i], ints.S[index] = ints.S[index], ints.S[i]
		}
	}
}
