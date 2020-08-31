package insert_sort

func InsertSort(s []int)  {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}
