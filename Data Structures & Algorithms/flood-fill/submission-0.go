func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	var dfs func(int, int)
	og := image[sr][sc]

	seen := make([][]int, len(image))
	for i := range seen {
		seen[i] = make([]int, len(image[0]))
	}

	dfs = func(i, j int) {
		COL, ROWS := len(image[0]), len(image)

		if i == ROWS || j == COL {
			return
		}

		if i < 0 || j < 0 {
			return
		}

		if seen[i][j] == 1{
			return
		}

		if image[i][j] != og {
			return
		}

		image[i][j] = color

		seen[i][j] = 1

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)

		seen[i][j] = 0
	}

	dfs(sr,sc)

	return image

}
