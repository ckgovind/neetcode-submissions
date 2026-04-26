
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, val := range nums {
		freqMap[val] += 1
	}

	a, b := make([]int, 0), make([]int, 0)
	for key, freq := range freqMap {
		a = append(a, key)
		b = append(b, freq)
	}

	_, a = BubbleSortAWithB(b, a)
	return a[:k]

}

func BubbleSortAWithB(a []int, b []int) ([]int, []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i; j < len(a); j++ {
			if a[i] < a[j] {
				a[j], a[i] = a[i], a[j]
				b[j], b[i] = b[i], b[j]
			}
		}
	}
	return a, b
}
