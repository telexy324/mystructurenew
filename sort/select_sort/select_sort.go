package select_sort

func SelectSort(s []int) {
	for i := 0; i < len(s); i++ {
		minIndex := i
		minValue := s[minIndex]
		for j := i + 1; j < len(s); j++ {
			if s[j] < minValue {
				minIndex = j
				minValue = s[j]
			}
		}
		s[i], s[minIndex] = s[minIndex], s[i]
	}
}
