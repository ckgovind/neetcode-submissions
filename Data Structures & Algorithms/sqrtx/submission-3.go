func mySqrt(x int) int {
	for i := 0; i <= x/2;i++{
		curr := i*i

		if curr > x{
			return i-1
		}

		if curr == x{
			return i
		}
	}
	return 1
}
