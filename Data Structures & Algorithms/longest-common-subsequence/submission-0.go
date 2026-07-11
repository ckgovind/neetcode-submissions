func mx(arr ...int) int {
	ans := -1
	for _, val := range arr {
		ans = max(ans, val)
	}

	return ans
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	dp := make([][]int, m)
	// m rows - text1
	// n cols - text2

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i, val := range text1 {
		for j, val2 := range text2 {
			t := 0
			if val == val2 {
				t = 1
			}

			var c1, c2, c3 int
			if i > 0 {
				c1 = dp[i-1][j]
			}

			if j > 0 {
				c2 = dp[i][j-1]
			}

			if i > 0 && j > 0 {
				c3 = dp[i-1][j-1]
			}

			dp[i][j] = mx(c1, c2, c3+t)
		}
	}

	return dp[m-1][n-1]

}
