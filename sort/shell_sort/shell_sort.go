package shell_sort

func ShellSort(data []int) {
	n := len(data)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && data[j] < data[j-h]; j -= h {
				data[j], data[j-h] = data[j-h], data[j]
			}
		}
		h /= 3
	}
}
