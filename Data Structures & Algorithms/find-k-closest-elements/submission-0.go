
func binarySearch(arr []int, x int) int {
	left, right := 0, len(arr)-1
	var mid int
	var best int
	for left <= right {
		mid = left + (right-left)/2

		if abs(arr[mid]-x) < abs(arr[best]-x) {
			best = mid
		}

		if arr[mid] == x {
			return mid
		}

		if x > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	curr := binarySearch(arr, x)

	fmt.Println("The curr obtained is ", curr)
	ans := make([]int, 0)
	l, r := curr, curr
	k -= 1
	for k > 0 {
		if l == 0 {
			r += 1
		} else if r == len(arr)-1 {
			l -= 1
		} else {
			v1 := arr[l-1]
			v2 := arr[r+1]

			if abs(x-v1) > abs(x-v2) {
				// v1 is farther
				r += 1
			} else {
				l -= 1
			}
		}

		k--

		fmt.Printf("At the end of the current loop the values are l: %v r: %v k: %v\n", l, r, k)
	}

	for i := l; i <= r; i++ {
		ans = append(ans, arr[i])
	}

	return ans
}
