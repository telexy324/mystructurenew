package permutation

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPermutation(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	i:=random.Intn(10)
	fmt.Println(i)
	Permutation(i)
}

