package quick_sort

func QuickSort(s []int) {
	if len(s) < 2 {
		return
	}
	mid := s[0]
	tail, head := len(s)-1, 0
	for i := 1; i <= tail; {
		if s[i] > mid {
			s[tail], s[i] = s[i], s[tail]
			tail--
		} else {
			s[head], s[i] = s[i], s[head]
			head++
			i++
		}
	}
	QuickSort(s[:head])
	QuickSort(s[head+1:])
}


