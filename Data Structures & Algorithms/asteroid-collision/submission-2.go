func asteroidCollision(asteroids []int) []int {
	ans := make([]int, 0)

	for _, ast := range asteroids {
		if ast > 0 {
			// moving right
			ans = append(ans, ast)
		} else {

			if len(ans) == 0 {
				ans = append(ans, ast)
				continue
			}

			for len(ans) > 0 {
				// moving left
				curr := ans[len(ans)-1]
				if curr < 0 {
					// no more asteroid moving right
					ans = append(ans, ast)
					break
				} else {
					// curr > 0 and ast < 0
					if curr+ast > 0 {
						// right movign ast wins
						break
					} else if curr+ast < 0 {
						// left moving wins
						ans = ans[0 : len(ans)-1]
					} else {
						// both get destroyed
						ans = ans[0 : len(ans)-1]
						break
					}
				}

				if len(ans) == 0 {
					ans = append(ans, ast)
					break
				}
			}

		}
	}

	return ans
}