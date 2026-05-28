func rotate(nums []int, k int) {

	k = k % len(nums)
	var count int

	for start := 0; count < len(nums); start++ {
		curr := start
		prev := nums[curr]
		for {
			nxtindx := (curr + k) % len(nums)
			temp := nums[nxtindx]
			nums[nxtindx] = prev
			prev = temp
			curr = nxtindx
			count++
			if curr == start {
				break
			}
		}
	}
}
