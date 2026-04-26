func merge(a, b []int) []int {
	c := make([]int, len(a)+len(b))
	j, k := 0, 0
	for i := 0; i < len(a)+len(b); i++ {
		if j == len(a) {
			c[i] = b[k]
			k += 1
		} else if k == len(b) {
			c[i] = a[j]
			j += 1
		} else {
			if a[j] <= b[k] {
				c[i] = a[j]
				j += 1
			} else {
				c[i] = b[k]
				k += 1
			}
		}
	}
	return c
}

func Mergesort(inp []int) []int {
	if len(inp) == 1 || len(inp) == 0 {
		return inp
	}
	x := int(len(inp) / 2)
	a := Mergesort(inp[0:x])
	b := Mergesort(inp[x:])
	c := merge(a, b)
	return c
}


func sortArray(nums []int) []int {
    return Mergesort(nums)
}
