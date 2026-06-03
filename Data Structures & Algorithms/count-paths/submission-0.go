func uniquePaths(m int, n int) int {
	
	currRow := make([]int, n+1)
	currRow[n-1] = 1
	prevRow := make([]int,n)
	for i := m-1 ; i >=0 ;i--{
		for j := n-1 ; j >=0 ; j--{
			if currRow[j] != 0{
				continue
			}
			currRow[j] = currRow[j+1] + prevRow[j]
		}
		prevRow = currRow
		currRow = make([]int,n+1)
	}
	return prevRow[0]
}
